package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang-jwt/jwt/v5"
	"github.com/metal-toolbox/iam-runtime-contrib/iamruntime"
	"github.com/metal-toolbox/iam-runtime-contrib/mockruntime"
	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/authorization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/tenant-api/internal/ent/generated"
	"go.infratographer.com/tenant-api/internal/testclient"
)

func TestTenantQueryByID(t *testing.T) {
	ctx := context.Background()

	runtime := new(mockruntime.MockRuntime)
	runtime.On("CheckAccess", mock.Anything).Return(authorization.CheckAccessResponse_RESULT_ALLOWED, nil)
	runtime.On("CreateRelationships", mock.Anything, mock.Anything).Return(nil)

	ctx = iamruntime.SetContextRuntime(ctx, runtime)
	ctx = iamruntime.SetContextToken(ctx, &jwt.Token{})

	tenant := TenantBuilder{}.MustNew(ctx)
	tenantChild := TenantBuilder{Parent: tenant}.MustNew(ctx)

	testCases := []struct {
		TestName       string
		ID             gidx.PrefixedID
		ResponseTenant *ent.Tenant
		errorMsg       string
	}{
		{
			TestName:       "root tenant",
			ID:             tenant.ID,
			ResponseTenant: tenant,
		},
		{
			TestName:       "child tenant of a tenant",
			ID:             tenantChild.ID,
			ResponseTenant: tenantChild,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient(testTools.entClient).GetTenant(ctx, tt.ID)

			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)

				return
			}

			require.NoError(t, err)
			assert.NotNil(t, resp.Tenant)
			assert.EqualValues(t, tt.ResponseTenant.ID, resp.Tenant.ID)
			assert.EqualValues(t, tt.ResponseTenant.Name, resp.Tenant.Name)
			assert.EqualValues(t, &tt.ResponseTenant.Description, resp.Tenant.Description)

			if tt.ResponseTenant.ParentTenantID == gidx.NullPrefixedID {
				assert.Nil(t, resp.Tenant.Parent)
			} else {
				require.NotNil(t, resp.Tenant.Parent)

				parent, err := tt.ResponseTenant.Parent(ctx)
				require.NoError(t, err)
				require.NotNil(t, parent)

				assert.Equal(t, parent.ID, resp.Tenant.Parent.ID)
				assert.Equal(t, parent.Name, resp.Tenant.Parent.Name)
			}
		})
	}
}

func TestTenantChildrenSorting(t *testing.T) {
	ctx := context.Background()

	runtime := new(mockruntime.MockRuntime)
	runtime.On("CheckAccess", mock.Anything).Return(authorization.CheckAccessResponse_RESULT_ALLOWED, nil)
	runtime.On("CreateRelationships", mock.Anything, mock.Anything).Return(nil)

	ctx = iamruntime.SetContextRuntime(ctx, runtime)
	ctx = iamruntime.SetContextToken(ctx, &jwt.Token{})

	tenant := TenantBuilder{}.MustNew(ctx)
	nicole := TenantBuilder{Parent: tenant, Name: "Nicole"}.MustNew(ctx)
	sarah := TenantBuilder{Parent: tenant, Name: "Sarah"}.MustNew(ctx)
	andy := TenantBuilder{Parent: tenant, Name: "Andy"}.MustNew(ctx)
	// Update sarah so it's updated at is most recent to verify sorting timestamps
	sarah.Update().SaveX(ctx)

	testCases := []struct {
		TestName      string
		OrderBy       *testclient.TenantOrder
		TenantID      gidx.PrefixedID
		ResponseOrder []*ent.Tenant
		errorMsg      string
	}{
		{
			TestName:      "Ordered By NAME ASC",
			OrderBy:       &testclient.TenantOrder{Field: "NAME", Direction: "ASC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{andy, nicole, sarah},
		},
		{
			TestName:      "Ordered By NAME DESC",
			OrderBy:       &testclient.TenantOrder{Field: "NAME", Direction: "DESC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{sarah, nicole, andy},
		},
		{
			TestName:      "Ordered By CREATED_AT ASC",
			OrderBy:       &testclient.TenantOrder{Field: "CREATED_AT", Direction: "ASC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{nicole, sarah, andy},
		},
		{
			TestName:      "Ordered By CREATED_AT DESC",
			OrderBy:       &testclient.TenantOrder{Field: "CREATED_AT", Direction: "DESC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{andy, sarah, nicole},
		},
		{
			TestName:      "Ordered By UPDATED_AT ASC",
			OrderBy:       &testclient.TenantOrder{Field: "UPDATED_AT", Direction: "ASC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{nicole, andy, sarah},
		},
		{
			TestName:      "Ordered By UPDATED_AT DESC",
			OrderBy:       &testclient.TenantOrder{Field: "UPDATED_AT", Direction: "DESC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{sarah, andy, nicole},
		},
		{
			TestName:      "No Children",
			TenantID:      andy.ID,
			ResponseOrder: []*ent.Tenant{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient(testTools.entClient).GetTenantChildren(ctx, tt.TenantID, tt.OrderBy)

			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp.Tenant)
			require.Len(t, resp.Tenant.Children.Edges, len(tt.ResponseOrder))

			for i, tnt := range tt.ResponseOrder {
				respTnt := resp.Tenant.Children.Edges[i].Node
				require.Equal(t, tnt.ID, respTnt.ID)
				require.Equal(t, tnt.Name, respTnt.Name)
			}

			runtime.AssertExpectations(t)
		})
	}
}

func TestTenantChildrenWhereFiltering(t *testing.T) {
	ctx := context.Background()

	runtime := new(mockruntime.MockRuntime)
	runtime.On("CheckAccess", mock.Anything).Return(authorization.CheckAccessResponse_RESULT_ALLOWED, nil)
	runtime.On("CreateRelationships", mock.Anything, mock.Anything).Return(nil)

	ctx = iamruntime.SetContextRuntime(ctx, runtime)
	ctx = iamruntime.SetContextToken(ctx, &jwt.Token{})

	parent1 := TenantBuilder{}.MustNew(ctx)
	parent1Child := TenantBuilder{Parent: parent1}.MustNew(ctx)
	parent1ChildChild := TenantBuilder{Parent: parent1Child}.MustNew(ctx)
	parent2 := TenantBuilder{}.MustNew(ctx)
	parent2Child := TenantBuilder{Parent: parent2}.MustNew(ctx)

	testCases := []struct {
		TestName      string
		TenantID      gidx.PrefixedID
		ChildID       gidx.PrefixedID
		ResponseChild *ent.Tenant
		errorMsg      string
	}{
		{
			TestName:      "Parent can access a child tenant",
			TenantID:      parent1.ID,
			ChildID:       parent1Child.ID,
			ResponseChild: parent1Child,
		},
		{
			TestName: "Parent can't go directly to a child tenant's child",
			TenantID: parent1.ID,
			ChildID:  parent1ChildChild.ID,
		},
		{
			TestName: "Parent can't go access any ID",
			TenantID: parent1.ID,
			ChildID:  parent2Child.ID,
		},
		{
			TestName:      "Child tenant can access it's child tenant",
			TenantID:      parent1Child.ID,
			ChildID:       parent1ChildChild.ID,
			ResponseChild: parent1ChildChild,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.TestName, func(t *testing.T) {
			resp, err := graphTestClient(testTools.entClient).GetTenantChildByID(ctx, tt.TenantID, tt.ChildID)

			if tt.errorMsg != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.errorMsg)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp.Tenant)

			if tt.ResponseChild == nil {
				require.Empty(t, resp.Tenant.Children.Edges)
				return
			}

			require.Len(t, resp.Tenant.Children.Edges, 1)
			assert.Equal(t, tt.ResponseChild.ID, resp.Tenant.Children.Edges[0].Node.ID)
			assert.Equal(t, tt.ResponseChild.Name, resp.Tenant.Children.Edges[0].Node.Name)

			runtime.AssertExpectations(t)
		})
	}
}

func TestFullTenantLifecycle(t *testing.T) {
	ctx := context.Background()

	name := gofakeit.DomainName()
	description := gofakeit.Phrase()

	graphC := graphTestClient(testTools.entClient)

	permitRuntime := new(mockruntime.MockRuntime)
	permitRuntime.On("CheckAccess", mock.Anything).Return(authorization.CheckAccessResponse_RESULT_ALLOWED, nil)
	permitRuntime.On("CreateRelationships", mock.Anything, mock.Anything).Return(nil)
	permitRuntime.On("DeleteRelationships", mock.Anything, mock.Anything).Return(nil)

	deniedRuntime := new(mockruntime.MockRuntime)
	deniedRuntime.On("CheckAccess", mock.Anything).Return(authorization.CheckAccessResponse_RESULT_DENIED, nil)

	permitCtx := iamruntime.SetContextRuntime(ctx, permitRuntime)
	permitCtx = iamruntime.SetContextToken(permitCtx, &jwt.Token{})

	deniedCtx := iamruntime.SetContextRuntime(ctx, deniedRuntime)
	deniedCtx = iamruntime.SetContextToken(deniedCtx, &jwt.Token{})

	// deny create the Root tenant
	rootResp, err := graphC.TenantCreate(deniedCtx, testclient.CreateTenantInput{
		Name:        name,
		Description: &description,
	})

	require.Error(t, err)
	require.ErrorContains(t, err, iamruntime.ErrAccessDenied.Error())
	require.Nil(t, rootResp)

	// create the Root tenant
	rootResp, err = graphC.TenantCreate(permitCtx, testclient.CreateTenantInput{
		Name:        name,
		Description: &description,
	})

	require.NoError(t, err)
	require.NotNil(t, rootResp)
	require.NotNil(t, rootResp.TenantCreate.Tenant)

	rootTenant := rootResp.TenantCreate.Tenant
	assert.NotNil(t, rootTenant.ID)
	assert.Equal(t, name, rootTenant.Name)
	assert.Equal(t, &description, rootTenant.Description)
	assert.Equal(t, "tnntten", rootTenant.ID.Prefix())
	assert.Nil(t, rootTenant.Parent)

	// Deny Update the tenant
	newName := gofakeit.DomainName()
	updatedTenantResp, err := graphC.TenantUpdate(deniedCtx, rootTenant.ID, testclient.UpdateTenantInput{Name: &newName})

	require.Error(t, err)
	require.ErrorContains(t, err, iamruntime.ErrAccessDenied.Error())
	require.Nil(t, updatedTenantResp)

	// Update the tenant
	updatedTenantResp, err = graphC.TenantUpdate(permitCtx, rootTenant.ID, testclient.UpdateTenantInput{Name: &newName})

	require.NoError(t, err)
	require.NotNil(t, updatedTenantResp)
	require.NotNil(t, updatedTenantResp.TenantUpdate.Tenant)

	updatedRootTenant := updatedTenantResp.TenantUpdate.Tenant
	assert.EqualValues(t, rootTenant.ID, updatedRootTenant.ID)
	assert.Equal(t, newName, updatedRootTenant.Name)

	// Deny query the tenant
	queryRootResp, err := graphC.GetTenant(deniedCtx, rootTenant.ID)

	require.Error(t, err)
	require.ErrorContains(t, err, iamruntime.ErrAccessDenied.Error())
	require.Nil(t, queryRootResp)

	// Query the tenant
	queryRootResp, err = graphC.GetTenant(permitCtx, rootTenant.ID)
	require.NoError(t, err)
	require.NotNil(t, queryRootResp)
	require.NotNil(t, queryRootResp.Tenant)
	require.Equal(t, newName, queryRootResp.Tenant.Name)

	// Add a child tenant with no description
	childResp, err := graphC.TenantCreate(permitCtx, testclient.CreateTenantInput{
		Name:     "child",
		ParentID: &rootTenant.ID,
	})

	require.NoError(t, err)
	require.NotNil(t, childResp)
	require.NotNil(t, childResp.TenantCreate.Tenant)

	childTenant := childResp.TenantCreate.Tenant
	assert.NotNil(t, childTenant.ID)
	assert.Equal(t, "child", childTenant.Name)
	assert.Empty(t, childTenant.Description)
	assert.Equal(t, "tnntten", childTenant.ID.Prefix())
	assert.NotNil(t, childTenant.Parent)
	assert.Equal(t, rootTenant.ID, childTenant.Parent.ID)

	// Try to delete the root tenant, it should fail since there are children
	deletedResp, err := graphC.TenantDelete(permitCtx, rootTenant.ID)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "tenant has children")
	assert.Nil(t, deletedResp)

	// delete the child tenant
	deletedResp, err = graphC.TenantDelete(permitCtx, childTenant.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.TenantDelete)
	assert.EqualValues(t, childTenant.ID, deletedResp.TenantDelete.DeletedID.String())

	// Deny delete the root tenant
	deletedResp, err = graphC.TenantDelete(deniedCtx, rootTenant.ID)

	require.Error(t, err)
	require.ErrorContains(t, err, iamruntime.ErrAccessDenied.Error())
	require.Nil(t, deletedResp)

	// delete the root tenant
	deletedResp, err = graphC.TenantDelete(permitCtx, rootTenant.ID)
	require.NoError(t, err)
	require.NotNil(t, deletedResp)
	require.NotNil(t, deletedResp.TenantDelete)
	assert.EqualValues(t, rootTenant.ID, deletedResp.TenantDelete.DeletedID.String())

	// Query the deleted root tenant to ensure it's no longer available
	queryResp, err := graphC.GetTenant(permitCtx, rootTenant.ID)
	assert.Error(t, err)
	assert.Nil(t, queryResp)
	assert.ErrorContains(t, err, "tenant not found")

	// Query the deleted tenant's child to ensure it's no longer available as well
	queryResp, err = graphC.GetTenant(permitCtx, childTenant.ID)
	assert.Error(t, err)
	assert.Nil(t, queryResp)
	assert.ErrorContains(t, err, "tenant not found")

	permitRuntime.AssertNumberOfCalls(t, "CheckAccess", 9)
	deniedRuntime.AssertNumberOfCalls(t, "CheckAccess", 4)

	// only one of each call should be registered, as the root create/delete have no relationships.
	permitRuntime.AssertNumberOfCalls(t, "CreateRelationships", 1)
	permitRuntime.AssertNumberOfCalls(t, "DeleteRelationships", 1)
}
