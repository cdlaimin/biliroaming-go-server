// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTHSubtitleCaches(t *testing.T) {
	t.Parallel()

	query := THSubtitleCaches()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTHSubtitleCachesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSubtitleCachesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := THSubtitleCaches().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSubtitleCachesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := THSubtitleCachSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTHSubtitleCachesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := THSubtitleCachExists(ctx, tx, o.EpisodeID)
	if err != nil {
		t.Errorf("Unable to check if THSubtitleCach exists: %s", err)
	}
	if !e {
		t.Errorf("Expected THSubtitleCachExists to return true, but got false.")
	}
}

func testTHSubtitleCachesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	thSubtitleCachFound, err := FindTHSubtitleCach(ctx, tx, o.EpisodeID)
	if err != nil {
		t.Error(err)
	}

	if thSubtitleCachFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTHSubtitleCachesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = THSubtitleCaches().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTHSubtitleCachesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := THSubtitleCaches().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTHSubtitleCachesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	thSubtitleCachOne := &THSubtitleCach{}
	thSubtitleCachTwo := &THSubtitleCach{}
	if err = randomize.Struct(seed, thSubtitleCachOne, thSubtitleCachDBTypes, false, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}
	if err = randomize.Struct(seed, thSubtitleCachTwo, thSubtitleCachDBTypes, false, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = thSubtitleCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = thSubtitleCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := THSubtitleCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTHSubtitleCachesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	thSubtitleCachOne := &THSubtitleCach{}
	thSubtitleCachTwo := &THSubtitleCach{}
	if err = randomize.Struct(seed, thSubtitleCachOne, thSubtitleCachDBTypes, false, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}
	if err = randomize.Struct(seed, thSubtitleCachTwo, thSubtitleCachDBTypes, false, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = thSubtitleCachOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = thSubtitleCachTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func thSubtitleCachBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func thSubtitleCachAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *THSubtitleCach) error {
	*o = THSubtitleCach{}
	return nil
}

func testTHSubtitleCachesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &THSubtitleCach{}
	o := &THSubtitleCach{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, false); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach object: %s", err)
	}

	AddTHSubtitleCachHook(boil.BeforeInsertHook, thSubtitleCachBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachBeforeInsertHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.AfterInsertHook, thSubtitleCachAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachAfterInsertHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.AfterSelectHook, thSubtitleCachAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachAfterSelectHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.BeforeUpdateHook, thSubtitleCachBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachBeforeUpdateHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.AfterUpdateHook, thSubtitleCachAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachAfterUpdateHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.BeforeDeleteHook, thSubtitleCachBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachBeforeDeleteHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.AfterDeleteHook, thSubtitleCachAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachAfterDeleteHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.BeforeUpsertHook, thSubtitleCachBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachBeforeUpsertHooks = []THSubtitleCachHook{}

	AddTHSubtitleCachHook(boil.AfterUpsertHook, thSubtitleCachAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	thSubtitleCachAfterUpsertHooks = []THSubtitleCachHook{}
}

func testTHSubtitleCachesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTHSubtitleCachesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(thSubtitleCachColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTHSubtitleCachesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
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

func testTHSubtitleCachesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := THSubtitleCachSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTHSubtitleCachesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := THSubtitleCaches().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	thSubtitleCachDBTypes = map[string]string{`EpisodeID`: `bigint`, `Data`: `json`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                     = bytes.MinRead
)

func testTHSubtitleCachesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(thSubtitleCachPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(thSubtitleCachAllColumns) == len(thSubtitleCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTHSubtitleCachesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(thSubtitleCachAllColumns) == len(thSubtitleCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &THSubtitleCach{}
	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, thSubtitleCachDBTypes, true, thSubtitleCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(thSubtitleCachAllColumns, thSubtitleCachPrimaryKeyColumns) {
		fields = thSubtitleCachAllColumns
	} else {
		fields = strmangle.SetComplement(
			thSubtitleCachAllColumns,
			thSubtitleCachPrimaryKeyColumns,
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

	slice := THSubtitleCachSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTHSubtitleCachesUpsert(t *testing.T) {
	t.Parallel()

	if len(thSubtitleCachAllColumns) == len(thSubtitleCachPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := THSubtitleCach{}
	if err = randomize.Struct(seed, &o, thSubtitleCachDBTypes, true); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert THSubtitleCach: %s", err)
	}

	count, err := THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, thSubtitleCachDBTypes, false, thSubtitleCachPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize THSubtitleCach struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert THSubtitleCach: %s", err)
	}

	count, err = THSubtitleCaches().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
