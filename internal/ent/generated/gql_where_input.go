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
	"errors"
	"fmt"
	"time"

	"go.infratographer.com/x/gidx"

	"go.infratographer.com/tenant-api/internal/ent/generated/predicate"
	"go.infratographer.com/tenant-api/internal/ent/generated/tenant"
)

// TenantWhereInput represents a where input for filtering Tenant queries.
type TenantWhereInput struct {
	Predicates []predicate.Tenant  `json:"-"`
	Not        *TenantWhereInput   `json:"not,omitempty"`
	Or         []*TenantWhereInput `json:"or,omitempty"`
	And        []*TenantWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNEQ   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGT    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGTE   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLT    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLTE   *gidx.PrefixedID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "parent" edge predicates.
	HasParent     *bool               `json:"hasParent,omitempty"`
	HasParentWith []*TenantWhereInput `json:"hasParentWith,omitempty"`

	// "children" edge predicates.
	HasChildren     *bool               `json:"hasChildren,omitempty"`
	HasChildrenWith []*TenantWhereInput `json:"hasChildrenWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *TenantWhereInput) AddPredicates(predicates ...predicate.Tenant) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the TenantWhereInput filter on the TenantQuery builder.
func (i *TenantWhereInput) Filter(q *TenantQuery) (*TenantQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyTenantWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyTenantWhereInput is returned in case the TenantWhereInput is empty.
var ErrEmptyTenantWhereInput = errors.New("generated: empty predicate TenantWhereInput")

// P returns a predicate for filtering tenants.
// An error is returned if the input is empty or invalid.
func (i *TenantWhereInput) P() (predicate.Tenant, error) {
	var predicates []predicate.Tenant
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, tenant.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Tenant, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, tenant.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Tenant, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, tenant.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, tenant.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, tenant.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, tenant.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, tenant.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, tenant.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, tenant.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, tenant.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, tenant.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, tenant.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, tenant.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, tenant.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, tenant.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, tenant.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, tenant.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, tenant.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, tenant.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, tenant.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, tenant.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, tenant.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, tenant.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, tenant.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, tenant.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, tenant.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, tenant.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	if i.HasParent != nil {
		p := tenant.HasParent()
		if !*i.HasParent {
			p = tenant.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasParentWith) > 0 {
		with := make([]predicate.Tenant, 0, len(i.HasParentWith))
		for _, w := range i.HasParentWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasParentWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, tenant.HasParentWith(with...))
	}
	if i.HasChildren != nil {
		p := tenant.HasChildren()
		if !*i.HasChildren {
			p = tenant.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasChildrenWith) > 0 {
		with := make([]predicate.Tenant, 0, len(i.HasChildrenWith))
		for _, w := range i.HasChildrenWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasChildrenWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, tenant.HasChildrenWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyTenantWhereInput
	case 1:
		return predicates[0], nil
	default:
		return tenant.And(predicates...), nil
	}
}
