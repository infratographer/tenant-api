package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"

	"go.infratographer.com/x/gidx"

	"go.infratographer.com/tenant-api/internal/ent/generated"
)

// FindTenantByID is the resolver for the findTenantByID field.
func (r *entityResolver) FindTenantByID(ctx context.Context, id gidx.PrefixedID) (*generated.Tenant, error) {
	return r.client.Tenant.Get(ctx, id)
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
