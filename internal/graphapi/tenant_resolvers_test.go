package graphapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/tenant-api/internal/ent/generated"
	"go.infratographer.com/tenant-api/internal/graphclient"
)

func TestTenantQueryByID(t *testing.T) {
	ctx := context.Background()
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
			resp, err := graphTestClient().GetTenant(ctx, tt.ID)

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
	tenant := TenantBuilder{}.MustNew(ctx)
	nicole := TenantBuilder{Parent: tenant, Name: "Nicole"}.MustNew(ctx)
	sarah := TenantBuilder{Parent: tenant, Name: "Sarah"}.MustNew(ctx)
	andy := TenantBuilder{Parent: tenant, Name: "Andy"}.MustNew(ctx)
	// Update sarah so it's updated at is most recent to verify sorting timestamps
	sarah.Update().SaveX(ctx)

	testCases := []struct {
		TestName      string
		OrderBy       *graphclient.TenantOrder
		TenantID      gidx.PrefixedID
		ResponseOrder []*ent.Tenant
		errorMsg      string
	}{
		{
			TestName:      "Ordered By NAME ASC",
			OrderBy:       &graphclient.TenantOrder{Field: "NAME", Direction: "ASC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{andy, nicole, sarah},
		},
		{
			TestName:      "Ordered By NAME DESC",
			OrderBy:       &graphclient.TenantOrder{Field: "NAME", Direction: "DESC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{sarah, nicole, andy},
		},
		{
			TestName:      "Ordered By CREATED_AT ASC",
			OrderBy:       &graphclient.TenantOrder{Field: "CREATED_AT", Direction: "ASC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{nicole, sarah, andy},
		},
		{
			TestName:      "Ordered By CREATED_AT DESC",
			OrderBy:       &graphclient.TenantOrder{Field: "CREATED_AT", Direction: "DESC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{andy, sarah, nicole},
		},
		{
			TestName:      "Ordered By UPDATED_AT ASC",
			OrderBy:       &graphclient.TenantOrder{Field: "UPDATED_AT", Direction: "ASC"},
			TenantID:      tenant.ID,
			ResponseOrder: []*ent.Tenant{nicole, andy, sarah},
		},
		{
			TestName:      "Ordered By UPDATED_AT DESC",
			OrderBy:       &graphclient.TenantOrder{Field: "UPDATED_AT", Direction: "DESC"},
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
			resp, err := graphTestClient().GetTenantChildren(ctx, tt.TenantID, tt.OrderBy)

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
		})
	}
}

func TestTenantChildrenWhereFiltering(t *testing.T) {
	ctx := context.Background()
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
			resp, err := graphTestClient().GetTenantChildByID(ctx, tt.TenantID, tt.ChildID)

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
		})
	}
}
