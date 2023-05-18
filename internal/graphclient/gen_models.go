// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphclient

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"go.infratographer.com/x/gidx"
)

type MetadataNode interface {
	IsMetadataNode()
	GetID() gidx.PrefixedID
}

// An object with an ID.
// Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
type Node interface {
	IsNode()
	// The id of the object.
	GetID() gidx.PrefixedID
}

type ResourceOwner interface {
	IsResourceOwner()
	GetID() gidx.PrefixedID
}

type Entity interface {
	IsEntity()
}

// Input information to create a tenant.
type CreateTenantInput struct {
	// The name of a tenant.
	Name string `json:"name"`
	// An optional description of the tenant.
	Description *string          `json:"description,omitempty"`
	ParentID    *gidx.PrefixedID `json:"parentID,omitempty"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor,omitempty"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor,omitempty"`
}

type Tenant struct {
	// ID for the tenant.
	ID        gidx.PrefixedID `json:"id"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	// The name of a tenant.
	Name string `json:"name"`
	// An optional description of the tenant.
	Description *string          `json:"description,omitempty"`
	Parent      *Tenant          `json:"parent,omitempty"`
	Children    TenantConnection `json:"children"`
}

func (Tenant) IsMetadataNode()             {}
func (this Tenant) GetID() gidx.PrefixedID { return this.ID }

func (Tenant) IsNode() {}

// The id of the object.

func (Tenant) IsResourceOwner() {}

func (Tenant) IsEntity() {}

// A connection to a list of items.
type TenantConnection struct {
	// A list of edges.
	Edges []*TenantEdge `json:"edges,omitempty"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// Return response from tenantCreate.
type TenantCreatePayload struct {
	// The created tenant.
	Tenant Tenant `json:"tenant"`
}

// Return response from tenantDelete.
type TenantDeletePayload struct {
	// The ID of the deleted tenant.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// An edge in a connection.
type TenantEdge struct {
	// The item at the end of the edge.
	Node *Tenant `json:"node,omitempty"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// Ordering options for Tenant connections
type TenantOrder struct {
	// The ordering direction.
	Direction OrderDirection `json:"direction"`
	// The field by which to order Tenants.
	Field TenantOrderField `json:"field"`
}

// Return response from tenantUpdate.
type TenantUpdatePayload struct {
	// The updated tenant.
	Tenant Tenant `json:"tenant"`
}

// TenantWhereInput is used for filtering Tenant objects.
// Input was generated by ent.
type TenantWhereInput struct {
	Not *TenantWhereInput   `json:"not,omitempty"`
	And []*TenantWhereInput `json:"and,omitempty"`
	Or  []*TenantWhereInput `json:"or,omitempty"`
	// id field predicates
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNeq   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGt    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGte   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLt    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLte   *gidx.PrefixedID  `json:"idLTE,omitempty"`
	// created_at field predicates
	CreatedAt      *time.Time   `json:"createdAt,omitempty"`
	CreatedAtNeq   *time.Time   `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []*time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []*time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGt    *time.Time   `json:"createdAtGT,omitempty"`
	CreatedAtGte   *time.Time   `json:"createdAtGTE,omitempty"`
	CreatedAtLt    *time.Time   `json:"createdAtLT,omitempty"`
	CreatedAtLte   *time.Time   `json:"createdAtLTE,omitempty"`
	// updated_at field predicates
	UpdatedAt      *time.Time   `json:"updatedAt,omitempty"`
	UpdatedAtNeq   *time.Time   `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []*time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []*time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGt    *time.Time   `json:"updatedAtGT,omitempty"`
	UpdatedAtGte   *time.Time   `json:"updatedAtGTE,omitempty"`
	UpdatedAtLt    *time.Time   `json:"updatedAtLT,omitempty"`
	UpdatedAtLte   *time.Time   `json:"updatedAtLTE,omitempty"`
	// parent edge predicates
	HasParent     *bool               `json:"hasParent,omitempty"`
	HasParentWith []*TenantWhereInput `json:"hasParentWith,omitempty"`
	// children edge predicates
	HasChildren     *bool               `json:"hasChildren,omitempty"`
	HasChildrenWith []*TenantWhereInput `json:"hasChildrenWith,omitempty"`
}

// Input information to update a tenant.
type UpdateTenantInput struct {
	// The name of a tenant.
	Name *string `json:"name,omitempty"`
	// An optional description of the tenant.
	Description      *string `json:"description,omitempty"`
	ClearDescription *bool   `json:"clearDescription,omitempty"`
}

type Service struct {
	Sdl *string `json:"sdl,omitempty"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Properties by which Tenant connections can be ordered.
type TenantOrderField string

const (
	TenantOrderFieldCreatedAt TenantOrderField = "CREATED_AT"
	TenantOrderFieldUpdatedAt TenantOrderField = "UPDATED_AT"
	TenantOrderFieldName      TenantOrderField = "NAME"
)

var AllTenantOrderField = []TenantOrderField{
	TenantOrderFieldCreatedAt,
	TenantOrderFieldUpdatedAt,
	TenantOrderFieldName,
}

func (e TenantOrderField) IsValid() bool {
	switch e {
	case TenantOrderFieldCreatedAt, TenantOrderFieldUpdatedAt, TenantOrderFieldName:
		return true
	}
	return false
}

func (e TenantOrderField) String() string {
	return string(e)
}

func (e *TenantOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TenantOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TenantOrderField", str)
	}
	return nil
}

func (e TenantOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
