// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"go.infratographer.com/x/gidx"

	"go.infratographer.com/tenant-api/internal/ent/generated/predicate"
	"go.infratographer.com/tenant-api/internal/ent/generated/tenant"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTenant = "Tenant"
)

// TenantMutation represents an operation that mutates the Tenant nodes in the graph.
type TenantMutation struct {
	config
	op              Op
	typ             string
	id              *gidx.PrefixedID
	created_at      *time.Time
	updated_at      *time.Time
	name            *string
	description     *string
	clearedFields   map[string]struct{}
	parent          *gidx.PrefixedID
	clearedparent   bool
	children        map[gidx.PrefixedID]struct{}
	removedchildren map[gidx.PrefixedID]struct{}
	clearedchildren bool
	done            bool
	oldValue        func(context.Context) (*Tenant, error)
	predicates      []predicate.Tenant
}

var _ ent.Mutation = (*TenantMutation)(nil)

// tenantOption allows management of the mutation configuration using functional options.
type tenantOption func(*TenantMutation)

// newTenantMutation creates new mutation for the Tenant entity.
func newTenantMutation(c config, op Op, opts ...tenantOption) *TenantMutation {
	m := &TenantMutation{
		config:        c,
		op:            op,
		typ:           TypeTenant,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTenantID sets the ID field of the mutation.
func withTenantID(id gidx.PrefixedID) tenantOption {
	return func(m *TenantMutation) {
		var (
			err   error
			once  sync.Once
			value *Tenant
		)
		m.oldValue = func(ctx context.Context) (*Tenant, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tenant.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTenant sets the old Tenant of the mutation.
func withTenant(node *Tenant) tenantOption {
	return func(m *TenantMutation) {
		m.oldValue = func(context.Context) (*Tenant, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TenantMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TenantMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("generated: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Tenant entities.
func (m *TenantMutation) SetID(id gidx.PrefixedID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TenantMutation) ID() (id gidx.PrefixedID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TenantMutation) IDs(ctx context.Context) ([]gidx.PrefixedID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []gidx.PrefixedID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tenant.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *TenantMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TenantMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TenantMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TenantMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TenantMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TenantMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *TenantMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TenantMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TenantMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *TenantMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *TenantMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ClearDescription clears the value of the "description" field.
func (m *TenantMutation) ClearDescription() {
	m.description = nil
	m.clearedFields[tenant.FieldDescription] = struct{}{}
}

// DescriptionCleared returns if the "description" field was cleared in this mutation.
func (m *TenantMutation) DescriptionCleared() bool {
	_, ok := m.clearedFields[tenant.FieldDescription]
	return ok
}

// ResetDescription resets all changes to the "description" field.
func (m *TenantMutation) ResetDescription() {
	m.description = nil
	delete(m.clearedFields, tenant.FieldDescription)
}

// SetParentTenantID sets the "parent_tenant_id" field.
func (m *TenantMutation) SetParentTenantID(gi gidx.PrefixedID) {
	m.parent = &gi
}

// ParentTenantID returns the value of the "parent_tenant_id" field in the mutation.
func (m *TenantMutation) ParentTenantID() (r gidx.PrefixedID, exists bool) {
	v := m.parent
	if v == nil {
		return
	}
	return *v, true
}

// OldParentTenantID returns the old "parent_tenant_id" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldParentTenantID(ctx context.Context) (v gidx.PrefixedID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldParentTenantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldParentTenantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldParentTenantID: %w", err)
	}
	return oldValue.ParentTenantID, nil
}

// ClearParentTenantID clears the value of the "parent_tenant_id" field.
func (m *TenantMutation) ClearParentTenantID() {
	m.parent = nil
	m.clearedFields[tenant.FieldParentTenantID] = struct{}{}
}

// ParentTenantIDCleared returns if the "parent_tenant_id" field was cleared in this mutation.
func (m *TenantMutation) ParentTenantIDCleared() bool {
	_, ok := m.clearedFields[tenant.FieldParentTenantID]
	return ok
}

// ResetParentTenantID resets all changes to the "parent_tenant_id" field.
func (m *TenantMutation) ResetParentTenantID() {
	m.parent = nil
	delete(m.clearedFields, tenant.FieldParentTenantID)
}

// SetParentID sets the "parent" edge to the Tenant entity by id.
func (m *TenantMutation) SetParentID(id gidx.PrefixedID) {
	m.parent = &id
}

// ClearParent clears the "parent" edge to the Tenant entity.
func (m *TenantMutation) ClearParent() {
	m.clearedparent = true
	m.clearedFields[tenant.FieldParentTenantID] = struct{}{}
}

// ParentCleared reports if the "parent" edge to the Tenant entity was cleared.
func (m *TenantMutation) ParentCleared() bool {
	return m.ParentTenantIDCleared() || m.clearedparent
}

// ParentID returns the "parent" edge ID in the mutation.
func (m *TenantMutation) ParentID() (id gidx.PrefixedID, exists bool) {
	if m.parent != nil {
		return *m.parent, true
	}
	return
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *TenantMutation) ParentIDs() (ids []gidx.PrefixedID) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *TenantMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// AddChildIDs adds the "children" edge to the Tenant entity by ids.
func (m *TenantMutation) AddChildIDs(ids ...gidx.PrefixedID) {
	if m.children == nil {
		m.children = make(map[gidx.PrefixedID]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the Tenant entity.
func (m *TenantMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the Tenant entity was cleared.
func (m *TenantMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the Tenant entity by IDs.
func (m *TenantMutation) RemoveChildIDs(ids ...gidx.PrefixedID) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[gidx.PrefixedID]struct{})
	}
	for i := range ids {
		delete(m.children, ids[i])
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the Tenant entity.
func (m *TenantMutation) RemovedChildrenIDs() (ids []gidx.PrefixedID) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *TenantMutation) ChildrenIDs() (ids []gidx.PrefixedID) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *TenantMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// Where appends a list predicates to the TenantMutation builder.
func (m *TenantMutation) Where(ps ...predicate.Tenant) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TenantMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TenantMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Tenant, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TenantMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TenantMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Tenant).
func (m *TenantMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TenantMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, tenant.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, tenant.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, tenant.FieldName)
	}
	if m.description != nil {
		fields = append(fields, tenant.FieldDescription)
	}
	if m.parent != nil {
		fields = append(fields, tenant.FieldParentTenantID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TenantMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tenant.FieldCreatedAt:
		return m.CreatedAt()
	case tenant.FieldUpdatedAt:
		return m.UpdatedAt()
	case tenant.FieldName:
		return m.Name()
	case tenant.FieldDescription:
		return m.Description()
	case tenant.FieldParentTenantID:
		return m.ParentTenantID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TenantMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tenant.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case tenant.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case tenant.FieldName:
		return m.OldName(ctx)
	case tenant.FieldDescription:
		return m.OldDescription(ctx)
	case tenant.FieldParentTenantID:
		return m.OldParentTenantID(ctx)
	}
	return nil, fmt.Errorf("unknown Tenant field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TenantMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tenant.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case tenant.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case tenant.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case tenant.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case tenant.FieldParentTenantID:
		v, ok := value.(gidx.PrefixedID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetParentTenantID(v)
		return nil
	}
	return fmt.Errorf("unknown Tenant field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TenantMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TenantMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TenantMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tenant numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TenantMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(tenant.FieldDescription) {
		fields = append(fields, tenant.FieldDescription)
	}
	if m.FieldCleared(tenant.FieldParentTenantID) {
		fields = append(fields, tenant.FieldParentTenantID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TenantMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TenantMutation) ClearField(name string) error {
	switch name {
	case tenant.FieldDescription:
		m.ClearDescription()
		return nil
	case tenant.FieldParentTenantID:
		m.ClearParentTenantID()
		return nil
	}
	return fmt.Errorf("unknown Tenant nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TenantMutation) ResetField(name string) error {
	switch name {
	case tenant.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case tenant.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case tenant.FieldName:
		m.ResetName()
		return nil
	case tenant.FieldDescription:
		m.ResetDescription()
		return nil
	case tenant.FieldParentTenantID:
		m.ResetParentTenantID()
		return nil
	}
	return fmt.Errorf("unknown Tenant field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TenantMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.parent != nil {
		edges = append(edges, tenant.EdgeParent)
	}
	if m.children != nil {
		edges = append(edges, tenant.EdgeChildren)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TenantMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tenant.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	case tenant.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TenantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchildren != nil {
		edges = append(edges, tenant.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TenantMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case tenant.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TenantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedparent {
		edges = append(edges, tenant.EdgeParent)
	}
	if m.clearedchildren {
		edges = append(edges, tenant.EdgeChildren)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TenantMutation) EdgeCleared(name string) bool {
	switch name {
	case tenant.EdgeParent:
		return m.clearedparent
	case tenant.EdgeChildren:
		return m.clearedchildren
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TenantMutation) ClearEdge(name string) error {
	switch name {
	case tenant.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown Tenant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TenantMutation) ResetEdge(name string) error {
	switch name {
	case tenant.EdgeParent:
		m.ResetParent()
		return nil
	case tenant.EdgeChildren:
		m.ResetChildren()
		return nil
	}
	return fmt.Errorf("unknown Tenant edge %s", name)
}
