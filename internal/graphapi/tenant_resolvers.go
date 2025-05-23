package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"errors"

	"github.com/metal-toolbox/iam-runtime-contrib/iamruntime"

	"go.infratographer.com/x/gidx"

	"go.infratographer.com/tenant-api/internal/ent/generated"
	"go.infratographer.com/tenant-api/internal/ent/generated/tenant"
)

// ErrTenantHasChildren is returned when attempting to delete a tenant which has child tenants.
var ErrTenantHasChildren = errors.New("tenant has children and can't be deleted")

// TenantCreate is the resolver for the tenantCreate field.
func (r *mutationResolver) TenantCreate(ctx context.Context, input generated.CreateTenantInput) (*TenantCreatePayload, error) {
	resource := gidx.NullPrefixedID

	if input.ParentID != nil {
		resource = *input.ParentID
	}

	if err := iamruntime.ContextCheckAccessTo(ctx, resource.String(), actionTenantCreate); err != nil {
		return nil, err
	}

	tnt, err := r.client.Tenant.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &TenantCreatePayload{Tenant: tnt}, nil
}

// TenantUpdate is the resolver for the tenantUpdate field.
func (r *mutationResolver) TenantUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateTenantInput) (*TenantUpdatePayload, error) {
	if err := iamruntime.ContextCheckAccessTo(ctx, id.String(), actionTenantUpdate); err != nil {
		return nil, err
	}

	tnt, err := r.client.Tenant.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &TenantUpdatePayload{Tenant: tnt}, nil
}

// TenantDelete is the resolver for the tenantDelete field.
func (r *mutationResolver) TenantDelete(ctx context.Context, id gidx.PrefixedID) (*TenantDeletePayload, error) {
	if err := iamruntime.ContextCheckAccessTo(ctx, id.String(), actionTenantDelete); err != nil {
		return nil, err
	}

	childrenCount, err := r.client.Tenant.Query().Where(tenant.ParentTenantID(id)).Count(ctx)
	if err != nil {
		return nil, err
	}

	if childrenCount != 0 {
		return nil, ErrTenantHasChildren
	}

	if err := r.client.Tenant.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &TenantDeletePayload{DeletedID: id}, nil
}

// Tenant is the resolver for the tenant field.
func (r *queryResolver) Tenant(ctx context.Context, id gidx.PrefixedID) (*generated.Tenant, error) {
	if err := iamruntime.ContextCheckAccessTo(ctx, id.String(), actionTenantGet); err != nil {
		return nil, err
	}

	return r.client.Tenant.Get(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
