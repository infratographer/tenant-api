// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

func testOidcIssuersUpsert(t *testing.T) {
	t.Parallel()

	if len(oidcIssuerAllColumns) == len(oidcIssuerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OidcIssuer{}
	if err = randomize.Struct(seed, &o, oidcIssuerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OidcIssuer: %s", err)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, oidcIssuerDBTypes, false, oidcIssuerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OidcIssuer: %s", err)
	}

	count, err = OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOidcIssuers(t *testing.T) {
	t.Parallel()

	query := OidcIssuers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOidcIssuersSoftDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOidcIssuersQuerySoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OidcIssuers().DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOidcIssuersSliceSoftDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OidcIssuerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, false); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOidcIssuersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOidcIssuersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OidcIssuers().DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOidcIssuersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OidcIssuerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx, true); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOidcIssuersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OidcIssuerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OidcIssuer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OidcIssuerExists to return true, but got false.")
	}
}

func testOidcIssuersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	oidcIssuerFound, err := FindOidcIssuer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if oidcIssuerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOidcIssuersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OidcIssuers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOidcIssuersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OidcIssuers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOidcIssuersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	oidcIssuerOne := &OidcIssuer{}
	oidcIssuerTwo := &OidcIssuer{}
	if err = randomize.Struct(seed, oidcIssuerOne, oidcIssuerDBTypes, false, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}
	if err = randomize.Struct(seed, oidcIssuerTwo, oidcIssuerDBTypes, false, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = oidcIssuerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = oidcIssuerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OidcIssuers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOidcIssuersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	oidcIssuerOne := &OidcIssuer{}
	oidcIssuerTwo := &OidcIssuer{}
	if err = randomize.Struct(seed, oidcIssuerOne, oidcIssuerDBTypes, false, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}
	if err = randomize.Struct(seed, oidcIssuerTwo, oidcIssuerDBTypes, false, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = oidcIssuerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = oidcIssuerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func oidcIssuerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func oidcIssuerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OidcIssuer) error {
	*o = OidcIssuer{}
	return nil
}

func testOidcIssuersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OidcIssuer{}
	o := &OidcIssuer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OidcIssuer object: %s", err)
	}

	AddOidcIssuerHook(boil.BeforeInsertHook, oidcIssuerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	oidcIssuerBeforeInsertHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.AfterInsertHook, oidcIssuerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	oidcIssuerAfterInsertHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.AfterSelectHook, oidcIssuerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	oidcIssuerAfterSelectHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.BeforeUpdateHook, oidcIssuerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	oidcIssuerBeforeUpdateHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.AfterUpdateHook, oidcIssuerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	oidcIssuerAfterUpdateHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.BeforeDeleteHook, oidcIssuerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	oidcIssuerBeforeDeleteHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.AfterDeleteHook, oidcIssuerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	oidcIssuerAfterDeleteHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.BeforeUpsertHook, oidcIssuerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	oidcIssuerBeforeUpsertHooks = []OidcIssuerHook{}

	AddOidcIssuerHook(boil.AfterUpsertHook, oidcIssuerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	oidcIssuerAfterUpsertHooks = []OidcIssuerHook{}
}

func testOidcIssuersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOidcIssuersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(oidcIssuerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOidcIssuerToManyUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OidcIssuer
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.OidcIssuerID = a.ID
	c.OidcIssuerID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Users().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.OidcIssuerID == b.OidcIssuerID {
			bFound = true
		}
		if v.OidcIssuerID == c.OidcIssuerID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OidcIssuerSlice{&a}
	if err = a.L.LoadUsers(ctx, tx, false, (*[]*OidcIssuer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Users = nil
	if err = a.L.LoadUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Users); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOidcIssuerToManyAddOpUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OidcIssuer
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, oidcIssuerDBTypes, false, strmangle.SetComplement(oidcIssuerPrimaryKeyColumns, oidcIssuerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.OidcIssuerID {
			t.Error("foreign key was wrong value", a.ID, first.OidcIssuerID)
		}
		if a.ID != second.OidcIssuerID {
			t.Error("foreign key was wrong value", a.ID, second.OidcIssuerID)
		}

		if first.R.OidcIssuer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.OidcIssuer != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Users[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Users[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Users().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOidcIssuerToOneTenantUsingTenant(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OidcIssuer
	var foreign Tenant

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, oidcIssuerDBTypes, false, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TenantID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Tenant().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OidcIssuerSlice{&local}
	if err = local.L.LoadTenant(ctx, tx, false, (*[]*OidcIssuer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Tenant == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Tenant = nil
	if err = local.L.LoadTenant(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Tenant == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOidcIssuerToOneSetOpTenantUsingTenant(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OidcIssuer
	var b, c Tenant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, oidcIssuerDBTypes, false, strmangle.SetComplement(oidcIssuerPrimaryKeyColumns, oidcIssuerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Tenant{&b, &c} {
		err = a.SetTenant(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Tenant != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OidcIssuers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TenantID != x.ID {
			t.Error("foreign key was wrong value", a.TenantID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TenantID))
		reflect.Indirect(reflect.ValueOf(&a.TenantID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TenantID != x.ID {
			t.Error("foreign key was wrong value", a.TenantID, x.ID)
		}
	}
}

func testOidcIssuersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOidcIssuersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OidcIssuerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOidcIssuersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OidcIssuers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	oidcIssuerDBTypes = map[string]string{`ID`: `uuid`, `Name`: `string`, `TenantID`: `uuid`, `URI`: `string`, `Audience`: `string`, `ClientID`: `string`, `SubjectClaim`: `string`, `EmailClaim`: `string`, `NameClaim`: `string`, `CreatedAt`: `timestamptz`, `UpdatedAt`: `timestamptz`, `DeletedAt`: `timestamptz`}
	_                 = bytes.MinRead
)

func testOidcIssuersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(oidcIssuerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(oidcIssuerAllColumns) == len(oidcIssuerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOidcIssuersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(oidcIssuerAllColumns) == len(oidcIssuerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OidcIssuer{}
	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OidcIssuers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, oidcIssuerDBTypes, true, oidcIssuerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OidcIssuer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(oidcIssuerAllColumns, oidcIssuerPrimaryKeyColumns) {
		fields = oidcIssuerAllColumns
	} else {
		fields = strmangle.SetComplement(
			oidcIssuerAllColumns,
			oidcIssuerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OidcIssuerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}