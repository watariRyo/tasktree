// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testScheduledTasks(t *testing.T) {
	t.Parallel()

	query := ScheduledTasks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testScheduledTasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
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

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScheduledTasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ScheduledTasks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScheduledTasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScheduledTaskSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testScheduledTasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ScheduledTaskExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ScheduledTask exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ScheduledTaskExists to return true, but got false.")
	}
}

func testScheduledTasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	scheduledTaskFound, err := FindScheduledTask(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if scheduledTaskFound == nil {
		t.Error("want a record, got nil")
	}
}

func testScheduledTasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ScheduledTasks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testScheduledTasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ScheduledTasks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testScheduledTasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	scheduledTaskOne := &ScheduledTask{}
	scheduledTaskTwo := &ScheduledTask{}
	if err = randomize.Struct(seed, scheduledTaskOne, scheduledTaskDBTypes, false, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}
	if err = randomize.Struct(seed, scheduledTaskTwo, scheduledTaskDBTypes, false, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scheduledTaskOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scheduledTaskTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ScheduledTasks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testScheduledTasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	scheduledTaskOne := &ScheduledTask{}
	scheduledTaskTwo := &ScheduledTask{}
	if err = randomize.Struct(seed, scheduledTaskOne, scheduledTaskDBTypes, false, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}
	if err = randomize.Struct(seed, scheduledTaskTwo, scheduledTaskDBTypes, false, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = scheduledTaskOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = scheduledTaskTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func scheduledTaskBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func scheduledTaskAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ScheduledTask) error {
	*o = ScheduledTask{}
	return nil
}

func testScheduledTasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ScheduledTask{}
	o := &ScheduledTask{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ScheduledTask object: %s", err)
	}

	AddScheduledTaskHook(boil.BeforeInsertHook, scheduledTaskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	scheduledTaskBeforeInsertHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.AfterInsertHook, scheduledTaskAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	scheduledTaskAfterInsertHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.AfterSelectHook, scheduledTaskAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	scheduledTaskAfterSelectHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.BeforeUpdateHook, scheduledTaskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	scheduledTaskBeforeUpdateHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.AfterUpdateHook, scheduledTaskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	scheduledTaskAfterUpdateHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.BeforeDeleteHook, scheduledTaskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	scheduledTaskBeforeDeleteHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.AfterDeleteHook, scheduledTaskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	scheduledTaskAfterDeleteHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.BeforeUpsertHook, scheduledTaskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	scheduledTaskBeforeUpsertHooks = []ScheduledTaskHook{}

	AddScheduledTaskHook(boil.AfterUpsertHook, scheduledTaskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	scheduledTaskAfterUpsertHooks = []ScheduledTaskHook{}
}

func testScheduledTasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testScheduledTasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(scheduledTaskColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testScheduledTaskToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ScheduledTask
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, scheduledTaskDBTypes, false, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ScheduledTaskSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*ScheduledTask)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testScheduledTaskToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ScheduledTask
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, scheduledTaskDBTypes, false, strmangle.SetComplement(scheduledTaskPrimaryKeyColumns, scheduledTaskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ScheduledTasks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testScheduledTasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
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

func testScheduledTasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ScheduledTaskSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testScheduledTasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ScheduledTasks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	scheduledTaskDBTypes = map[string]string{`ID`: `bigint`, `UserID`: `int`, `Title`: `varchar`, `Content`: `text`, `ScheduledDate`: `date`, `StartTime`: `time`, `EndTime`: `time`, `CreatedAt`: `datetime`, `UpdatedAt`: `datetime`}
	_                    = bytes.MinRead
)

func testScheduledTasksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(scheduledTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(scheduledTaskAllColumns) == len(scheduledTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testScheduledTasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(scheduledTaskAllColumns) == len(scheduledTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ScheduledTask{}
	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, scheduledTaskDBTypes, true, scheduledTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(scheduledTaskAllColumns, scheduledTaskPrimaryKeyColumns) {
		fields = scheduledTaskAllColumns
	} else {
		fields = strmangle.SetComplement(
			scheduledTaskAllColumns,
			scheduledTaskPrimaryKeyColumns,
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

	slice := ScheduledTaskSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testScheduledTasksUpsert(t *testing.T) {
	t.Parallel()

	if len(scheduledTaskAllColumns) == len(scheduledTaskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLScheduledTaskUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ScheduledTask{}
	if err = randomize.Struct(seed, &o, scheduledTaskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ScheduledTask: %s", err)
	}

	count, err := ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, scheduledTaskDBTypes, false, scheduledTaskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ScheduledTask struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ScheduledTask: %s", err)
	}

	count, err = ScheduledTasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}