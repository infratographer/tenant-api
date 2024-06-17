package graphapi_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang-jwt/jwt/v5"
	"github.com/metal-toolbox/iam-runtime-contrib/iamruntime"
	"github.com/metal-toolbox/iam-runtime-contrib/mockruntime"
	"github.com/metal-toolbox/iam-runtime/pkg/iam/runtime/authorization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/tenant-api/internal/testclient"
)

func TestTenantPubsub(t *testing.T) {
	ctx := context.Background()

	runtime := new(mockruntime.MockRuntime)
	runtime.On("CheckAccess", mock.Anything).Return(authorization.CheckAccessResponse_RESULT_ALLOWED, nil)
	runtime.On("CreateRelationships", mock.Anything, mock.Anything).Return(nil)
	runtime.On("DeleteRelationships", mock.Anything, mock.Anything).Return(nil)

	ctx = iamruntime.SetContextRuntime(ctx, runtime)
	ctx = iamruntime.SetContextToken(ctx, &jwt.Token{})

	name := gofakeit.DomainName()
	description := gofakeit.Phrase()

	graphC := graphTestClient(testTools.pubsubEntClient)

	sub, err := events.NewConnection(testTools.eventsConfig)
	require.NoError(t, err)

	defer sub.Shutdown(ctx) //nolint:errcheck // skip check in test

	messages, err := sub.SubscribeChanges(ctx, ">")
	require.NoError(t, err)

	// create a root tenant and ensure fields are set
	rootResp, err := graphC.TenantCreate(ctx, testclient.CreateTenantInput{
		Name:        name,
		Description: &description,
	})
	require.NoError(t, err)

	rootTenant := rootResp.TenantCreate.Tenant
	msg := getSingleMessage(t, messages)
	assert.Equal(t, "testing-roundtrip-actor", msg.ActorID.String())
	assert.Equal(t, "create", msg.EventType)
	assert.Equal(t, "tenant-api-test", msg.Source)
	assert.Equal(t, rootTenant.ID, msg.SubjectID)
	assert.Empty(t, msg.AdditionalSubjectIDs)
	// expect created_at, updated_at, name, and description changeset
	assert.Len(t, msg.FieldChanges, 4)

	var createdAtVisited, updatedAtVisited, nameVisited, descriptionVisited bool

	for _, change := range msg.FieldChanges {
		assert.Empty(t, change.PreviousValue)

		switch change.Field {
		case "created_at":
			createdAtVisited = true
			ts, err := time.Parse(time.RFC3339, change.CurrentValue)
			assert.NoError(t, err)
			assert.WithinDuration(t, time.Now(), ts, 2*time.Second)
		case "updated_at":
			updatedAtVisited = true
			ts, err := time.Parse(time.RFC3339, change.CurrentValue)
			assert.NoError(t, err)
			assert.WithinDuration(t, time.Now(), ts, 2*time.Second)
		case "name":
			nameVisited = true

			assert.EqualValues(t, name, change.CurrentValue)
		case "description":
			descriptionVisited = true

			assert.EqualValues(t, description, change.CurrentValue)
		default:
			assert.Fail(t, "unexpected field in changeset %s")
			t.Fail()
		}
	}

	assert.True(t, createdAtVisited)
	assert.True(t, updatedAtVisited)
	assert.True(t, nameVisited)
	assert.True(t, descriptionVisited)

	// Add a child tenant with no description
	childResp, err := graphC.TenantCreate(ctx, testclient.CreateTenantInput{
		Name:     "child",
		ParentID: &rootTenant.ID,
	})

	require.NoError(t, err)
	require.NotNil(t, childResp)

	childTnt := childResp.TenantCreate.Tenant

	msg = getSingleMessage(t, messages)
	assert.Equal(t, "testing-roundtrip-actor", msg.ActorID.String())
	assert.Equal(t, "create", msg.EventType)
	assert.Equal(t, "tenant-api-test", msg.Source)
	assert.Equal(t, childTnt.ID, msg.SubjectID)
	assert.EqualValues(t, []gidx.PrefixedID{rootTenant.ID}, msg.AdditionalSubjectIDs)
	// expect created_at, updated_at, name, and description changeset
	assert.Len(t, msg.FieldChanges, 4)

	createdAtVisited = false
	updatedAtVisited = false
	nameVisited = false

	var parentIDVisited bool

	for _, change := range msg.FieldChanges {
		assert.Empty(t, change.PreviousValue)

		switch change.Field {
		case "created_at":
			createdAtVisited = true
			ts, err := time.Parse(time.RFC3339, change.CurrentValue)
			assert.NoError(t, err)
			assert.WithinDuration(t, time.Now(), ts, 2*time.Second)
		case "updated_at":
			updatedAtVisited = true
			ts, err := time.Parse(time.RFC3339, change.CurrentValue)
			assert.NoError(t, err)
			assert.WithinDuration(t, time.Now(), ts, 2*time.Second)
		case "name":
			nameVisited = true

			assert.EqualValues(t, "child", change.CurrentValue)
		case "parent_tenant_id":
			parentIDVisited = true

			assert.EqualValues(t, rootTenant.ID.String(), change.CurrentValue)
		default:
			assert.Fail(t, fmt.Sprintf("unexpected field in changeset %s", change.Field))
			t.Fail()
		}
	}

	assert.True(t, createdAtVisited)
	assert.True(t, updatedAtVisited)
	assert.True(t, nameVisited)
	assert.True(t, parentIDVisited)

	// Update the tenant
	newName := gofakeit.DomainName()
	updatedTenantResp, err := graphC.TenantUpdate(ctx, childTnt.ID, testclient.UpdateTenantInput{Name: &newName})

	require.NoError(t, err)
	require.NotNil(t, updatedTenantResp)

	msg = getSingleMessage(t, messages)
	assert.Equal(t, "testing-roundtrip-actor", msg.ActorID.String())
	assert.Equal(t, "update", msg.EventType)
	assert.Equal(t, "tenant-api-test", msg.Source)
	assert.Equal(t, childTnt.ID, msg.SubjectID)
	assert.EqualValues(t, []gidx.PrefixedID{rootTenant.ID}, msg.AdditionalSubjectIDs)
	// expect updated_at, and name changeset
	assert.Len(t, msg.FieldChanges, 2)

	updatedAtVisited = false
	nameVisited = false

	for _, change := range msg.FieldChanges {
		assert.NotEmpty(t, change.PreviousValue)

		switch change.Field {
		case "updated_at":
			updatedAtVisited = true
			ts, err := time.Parse(time.RFC3339, change.CurrentValue)
			assert.NoError(t, err)
			assert.WithinDuration(t, time.Now(), ts, 2*time.Second)
		case "name":
			nameVisited = true

			assert.EqualValues(t, "child", change.PreviousValue)
			assert.EqualValues(t, newName, change.CurrentValue)
		default:
			assert.Fail(t, "unexpected field in changeset %s")
			t.Fail()
		}
	}

	assert.True(t, updatedAtVisited)
	assert.True(t, nameVisited)

	// delete the child tenant
	_, err = graphC.TenantDelete(ctx, childTnt.ID)
	require.NoError(t, err)

	msg = getSingleMessage(t, messages)
	assert.Equal(t, "testing-roundtrip-actor", msg.ActorID.String())
	assert.Equal(t, "delete", msg.EventType)
	assert.Equal(t, "tenant-api-test", msg.Source)
	assert.Equal(t, childTnt.ID, msg.SubjectID)
	assert.EqualValues(t, []gidx.PrefixedID{rootTenant.ID}, msg.AdditionalSubjectIDs)
	// expect created_at, updated_at, name, and description changeset
	assert.Len(t, msg.FieldChanges, 0)

	// delete the root tenant
	_, err = graphC.TenantDelete(ctx, rootTenant.ID)
	require.NoError(t, err)

	msg = getSingleMessage(t, messages)
	assert.Equal(t, "testing-roundtrip-actor", msg.ActorID.String())
	assert.Equal(t, "delete", msg.EventType)
	assert.Equal(t, "tenant-api-test", msg.Source)
	assert.Equal(t, rootTenant.ID, msg.SubjectID)
	assert.Empty(t, msg.AdditionalSubjectIDs)
	// expect created_at, updated_at, name, and description changeset
	assert.Len(t, msg.FieldChanges, 0)

	runtime.AssertNumberOfCalls(t, "CreateRelationships", 1)
	runtime.AssertNumberOfCalls(t, "DeleteRelationships", 1)
}

func getSingleMessage[T any](t *testing.T, messages <-chan events.Message[T]) T {
	select {
	case message := <-messages:
		require.NoError(t, message.Error())
		assert.NoError(t, message.Ack())

		return message.Message()
	case <-time.After(time.Second * 2):
		require.Fail(t, "timeout waiting for change message")
	}

	var empty T

	return empty
}
