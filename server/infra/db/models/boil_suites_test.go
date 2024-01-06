// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("BaseTasks", testBaseTasks)
	t.Run("DayOfWeeks", testDayOfWeeks)
	t.Run("ScheduledTasks", testScheduledTasks)
	t.Run("ShareTasks", testShareTasks)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksDelete)
	t.Run("DayOfWeeks", testDayOfWeeksDelete)
	t.Run("ScheduledTasks", testScheduledTasksDelete)
	t.Run("ShareTasks", testShareTasksDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksQueryDeleteAll)
	t.Run("DayOfWeeks", testDayOfWeeksQueryDeleteAll)
	t.Run("ScheduledTasks", testScheduledTasksQueryDeleteAll)
	t.Run("ShareTasks", testShareTasksQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksSliceDeleteAll)
	t.Run("DayOfWeeks", testDayOfWeeksSliceDeleteAll)
	t.Run("ScheduledTasks", testScheduledTasksSliceDeleteAll)
	t.Run("ShareTasks", testShareTasksSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksExists)
	t.Run("DayOfWeeks", testDayOfWeeksExists)
	t.Run("ScheduledTasks", testScheduledTasksExists)
	t.Run("ShareTasks", testShareTasksExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksFind)
	t.Run("DayOfWeeks", testDayOfWeeksFind)
	t.Run("ScheduledTasks", testScheduledTasksFind)
	t.Run("ShareTasks", testShareTasksFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksBind)
	t.Run("DayOfWeeks", testDayOfWeeksBind)
	t.Run("ScheduledTasks", testScheduledTasksBind)
	t.Run("ShareTasks", testShareTasksBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksOne)
	t.Run("DayOfWeeks", testDayOfWeeksOne)
	t.Run("ScheduledTasks", testScheduledTasksOne)
	t.Run("ShareTasks", testShareTasksOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksAll)
	t.Run("DayOfWeeks", testDayOfWeeksAll)
	t.Run("ScheduledTasks", testScheduledTasksAll)
	t.Run("ShareTasks", testShareTasksAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksCount)
	t.Run("DayOfWeeks", testDayOfWeeksCount)
	t.Run("ScheduledTasks", testScheduledTasksCount)
	t.Run("ShareTasks", testShareTasksCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksHooks)
	t.Run("DayOfWeeks", testDayOfWeeksHooks)
	t.Run("ScheduledTasks", testScheduledTasksHooks)
	t.Run("ShareTasks", testShareTasksHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksInsert)
	t.Run("BaseTasks", testBaseTasksInsertWhitelist)
	t.Run("DayOfWeeks", testDayOfWeeksInsert)
	t.Run("DayOfWeeks", testDayOfWeeksInsertWhitelist)
	t.Run("ScheduledTasks", testScheduledTasksInsert)
	t.Run("ScheduledTasks", testScheduledTasksInsertWhitelist)
	t.Run("ShareTasks", testShareTasksInsert)
	t.Run("ShareTasks", testShareTasksInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BaseTaskToDayOfWeekUsingDayOfWeek", testBaseTaskToOneDayOfWeekUsingDayOfWeek)
	t.Run("BaseTaskToUserUsingUser", testBaseTaskToOneUserUsingUser)
	t.Run("ScheduledTaskToUserUsingUser", testScheduledTaskToOneUserUsingUser)
	t.Run("ShareTaskToUserUsingUser", testShareTaskToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("DayOfWeekToBaseTasks", testDayOfWeekToManyBaseTasks)
	t.Run("UserToBaseTasks", testUserToManyBaseTasks)
	t.Run("UserToScheduledTasks", testUserToManyScheduledTasks)
	t.Run("UserToShareTasks", testUserToManyShareTasks)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BaseTaskToDayOfWeekUsingBaseTasks", testBaseTaskToOneSetOpDayOfWeekUsingDayOfWeek)
	t.Run("BaseTaskToUserUsingBaseTasks", testBaseTaskToOneSetOpUserUsingUser)
	t.Run("ScheduledTaskToUserUsingScheduledTasks", testScheduledTaskToOneSetOpUserUsingUser)
	t.Run("ShareTaskToUserUsingShareTasks", testShareTaskToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("DayOfWeekToBaseTasks", testDayOfWeekToManyAddOpBaseTasks)
	t.Run("UserToBaseTasks", testUserToManyAddOpBaseTasks)
	t.Run("UserToScheduledTasks", testUserToManyAddOpScheduledTasks)
	t.Run("UserToShareTasks", testUserToManyAddOpShareTasks)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksReload)
	t.Run("DayOfWeeks", testDayOfWeeksReload)
	t.Run("ScheduledTasks", testScheduledTasksReload)
	t.Run("ShareTasks", testShareTasksReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksReloadAll)
	t.Run("DayOfWeeks", testDayOfWeeksReloadAll)
	t.Run("ScheduledTasks", testScheduledTasksReloadAll)
	t.Run("ShareTasks", testShareTasksReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksSelect)
	t.Run("DayOfWeeks", testDayOfWeeksSelect)
	t.Run("ScheduledTasks", testScheduledTasksSelect)
	t.Run("ShareTasks", testShareTasksSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksUpdate)
	t.Run("DayOfWeeks", testDayOfWeeksUpdate)
	t.Run("ScheduledTasks", testScheduledTasksUpdate)
	t.Run("ShareTasks", testShareTasksUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("BaseTasks", testBaseTasksSliceUpdateAll)
	t.Run("DayOfWeeks", testDayOfWeeksSliceUpdateAll)
	t.Run("ScheduledTasks", testScheduledTasksSliceUpdateAll)
	t.Run("ShareTasks", testShareTasksSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}