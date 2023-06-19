// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TenantsColumns holds the columns for the "tenants" table.
	TenantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "parent_tenant_id", Type: field.TypeString, Nullable: true},
	}
	// TenantsTable holds the schema information for the "tenants" table.
	TenantsTable = &schema.Table{
		Name:       "tenants",
		Columns:    TenantsColumns,
		PrimaryKey: []*schema.Column{TenantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tenants_tenants_children",
				Columns:    []*schema.Column{TenantsColumns[5]},
				RefColumns: []*schema.Column{TenantsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tenant_created_at",
				Unique:  false,
				Columns: []*schema.Column{TenantsColumns[1]},
			},
			{
				Name:    "tenant_updated_at",
				Unique:  false,
				Columns: []*schema.Column{TenantsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TenantsTable,
	}
)

func init() {
	TenantsTable.ForeignKeys[0].RefTable = TenantsTable
}
