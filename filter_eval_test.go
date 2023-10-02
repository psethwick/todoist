package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	todoist "github.com/sachaos/todoist/lib"
)

const DateFormat = "Mon 2 Jan 2006 15:04:05 +0000"

var testTimeZone = time.FixedZone("JST", 9*60*60)

func due(s string) *todoist.Due {
	t, _ := time.Parse(DateFormat, s)
	t = t.In(testTimeZone)
	date := t.Format(todoist.RFC3339DateTime)
	return &todoist.Due{
		Date: date,
		// these tests will break in other timezones
		TimeZone: "Asia/Tokyo",
	}
}

func testFilterEval(t *testing.T, f string, item todoist.Item, expect bool) {
	actual, _ := Eval(Filter(f)[0], &item, &todoist.Store{})
	assert.Equal(t, expect, actual, "they should be equal")
}

func testFilterEvalWithProject(t *testing.T, f string, item todoist.Item, projects todoist.Projects, expect bool) {
	actual, _ := Eval(Filter(f)[0], &item, &todoist.Store{Projects: projects})
	assert.Equal(t, expect, actual, "they should be equal")
}

func testFilterEvalWithCollaborators(t *testing.T, f string, item todoist.Item, user todoist.User, collaborators todoist.Collaborators, expect bool) {
	actual, _ := Eval(Filter(f)[0], &item, &todoist.Store{User: user, Collaborators: collaborators})
	assert.Equal(t, expect, actual, "they should be equal")
}

func testFilterEvalWithLabel(t *testing.T, f string, item todoist.Item, expect bool) {
	actual, _ := Eval(Filter(f)[0], &item, &todoist.Store{})
	assert.Equal(t, expect, actual, "they should be equal")
}

func TestEval(t *testing.T) {
	testFilterEval(t, "", todoist.Item{}, true)
}

func TestPriorityEval(t *testing.T) {
	testFilterEval(t, "p4", todoist.Item{Priority: 1}, true)
	testFilterEval(t, "p3", todoist.Item{Priority: 1}, false)
}

func TestLabelEval(t *testing.T) {
	item1 := todoist.Item{}
	item1.LabelNames = []string{"must", "icebox"}

	testFilterEvalWithLabel(t, "@must", item1, true)
	testFilterEvalWithLabel(t, "@icebox", item1, true)
	testFilterEvalWithLabel(t, "@another", item1, false)
}

func TestAssignedEval(t *testing.T) {
	assigned := todoist.Item{
		ResponsibleUID: "12346",
	}
	notAssigned := todoist.Item{}
	testFilterEval(t, "assigned", assigned, true)
	testFilterEval(t, "assigned", notAssigned, false)
}

func TestNoPriorityEval(t *testing.T) {
	// p4 is 'no priority'
	// and p4 maps to 1
	testFilterEval(t, "no priority", todoist.Item{Priority: 1}, true)
	testFilterEval(t, "no priority", todoist.Item{Priority: 2}, false)
	testFilterEval(t, "no priority", todoist.Item{Priority: 3}, false)
	testFilterEval(t, "no priority", todoist.Item{Priority: 4}, false)
}

func TestProjectEval(t *testing.T) {
	projects := todoist.Projects{
		todoist.Project{
			HaveID: todoist.HaveID{ID: "1"},
			Name:   "private",
		},
		todoist.Project{
			HaveID:       todoist.HaveID{ID: "2"},
			HaveParentID: todoist.HaveParentID{ParentID: &[]string{"1"}[0]},
			Name:         "nested",
		},
		todoist.Project{
			HaveID: todoist.HaveID{ID: "3"},
			Name:   "Work",
		},
	}

	item1 := todoist.Item{}
	item1.ProjectID = "1"

	item2 := todoist.Item{}
	item2.ProjectID = "2"

	item3 := todoist.Item{}
	item3.ProjectID = "3"
	testFilterEvalWithProject(t, "#private", item1, projects, true)
	testFilterEvalWithProject(t, "#hoge", item1, projects, false)
	testFilterEvalWithProject(t, "#private", item2, projects, false)
	testFilterEvalWithProject(t, "##private", item2, projects, true)
	testFilterEvalWithProject(t, "#Work", item3, projects, true)
}

func TestBoolInfixOpExp(t *testing.T) {
	testFilterEval(t, "p3 | p4", todoist.Item{Priority: 1}, true)
	testFilterEval(t, "p3 | p4", todoist.Item{Priority: 2}, true)
	testFilterEval(t, "p3 | p4", todoist.Item{Priority: 3}, false)

	testFilterEval(t, "p3 & p4", todoist.Item{Priority: 1}, false)
	testFilterEval(t, "p3 & p4", todoist.Item{Priority: 2}, false)
	testFilterEval(t, "p3 & p4", todoist.Item{Priority: 3}, false)
}

func TestNotOpEval(t *testing.T) {
	testFilterEval(t, "!p4", todoist.Item{Priority: 1}, false)
	testFilterEval(t, "!(p3 | p4)", todoist.Item{Priority: 2}, false)
	testFilterEval(t, "!(p3 | p4)", todoist.Item{Priority: 3}, true)
}

func TestViewAllExp(t *testing.T) {
	testFilterEval(t, "view all", todoist.Item{Priority: 1}, true)
	testFilterEval(t, "view all", todoist.Item{Priority: 2}, true)
	testFilterEval(t, "view all", todoist.Item{Priority: 3}, true)
}

func TestNoTimeExp(t *testing.T) {
	timeNow := time.Date(2017, time.October, 2, 1, 0, 0, 0, testTimeZone) // JST: Mon 2 Oct 2017 00:00:00
	setNow(timeNow)
	// not midnight
	testFilterEval(t, "no time", todoist.Item{Due: due("Sun 1 Oct 2017 14:00:00 +0000")}, false)
	// midnight is a time
	testFilterEval(t, "no time", todoist.Item{Due: due("Sun 1 Oct 2017 15:00:00 +0000")}, false)

	testFilterEval(t, "no time", todoist.Item{}, true)
	// actually no time
	testFilterEval(t, "no time", todoist.Item{Due: &todoist.Due{Date: "2015-03-15"}}, true)
}

func TestDueOnEval(t *testing.T) {
	timeNow := time.Date(2017, time.October, 2, 1, 0, 0, 0, testTimeZone) // JST: Mon 2 Oct 2017 00:00:00
	setNow(timeNow)

	testFilterEval(t, "today", todoist.Item{Due: due("Sun 1 Oct 2017 15:00:00 +0000")}, true)  // JST: Mon 2 Oct 2017 00:00:00
	testFilterEval(t, "today", todoist.Item{Due: due("Mon 2 Oct 2017 14:59:59 +0000")}, true)  // JST: Mon 2 Oct 2017 23:59:59
	testFilterEval(t, "today", todoist.Item{Due: due("Mon 2 Oct 2017 15:00:00 +0000")}, false) // JST: Tue 3 Oct 2017 00:00:00

	testFilterEval(t, "yesterday", todoist.Item{Due: due("Sun 1 Oct 2017 14:59:59 +0000")}, true)   // JST: Sun 1 Oct 2017 23:59:59
	testFilterEval(t, "yesterday", todoist.Item{Due: due("Sat 30 Sep 2017 15:00:00 +0000")}, true)  // JST: Sun 1 Oct 2017 00:00:00
	testFilterEval(t, "yesterday", todoist.Item{Due: due("Sat 30 Sep 2017 14:59:59 +0000")}, false) // JST: Sat 30 Sept 2017 23:59:59
	testFilterEval(t, "tomorrow", todoist.Item{Due: due("Mon 2 Oct 2017 15:00:00 +0000")}, true)    // JST: Tue 3 Oct 2017 00:00:00
	testFilterEval(t, "tomorrow", todoist.Item{Due: due("Tue 3 Oct 2017 14:59:59 +0000")}, true)    // JST: Tue 3 Oct 2017 23:59:59
	testFilterEval(t, "tomorrow", todoist.Item{Due: due("Tue 3 Oct 2017 15:00:00 +0000")}, false)   // JST: Wed 4 Oct 2017 00:00:00

	testFilterEval(t, "10/2/2017", todoist.Item{Due: due("Mon 2 Oct 2017 01:00:00 +0000")}, true)        // JST: Mon 2 Oct 2017 10:00:00
	testFilterEval(t, "10/2/2017 10:00", todoist.Item{Due: due("Mon 2 Oct 2017 01:00:00 +0000")}, false) // JST: Mon 2 Oct 2017 10:00:00
}

func TestNoDateEval(t *testing.T) {
	testFilterEval(t, "no date", todoist.Item{Due: nil}, true)
	testFilterEval(t, "no due date", todoist.Item{Due: nil}, true)

	testFilterEval(t, "no date", todoist.Item{Due: due("Sun 1 Oct 2017 15:00:00 +0000")}, false) // JST: Mon 2 Oct 2017 00:00:00
}

func TestDueBeforeEval(t *testing.T) {
	timeNow := time.Date(2017, time.October, 2, 1, 0, 0, 0, testTimeZone) // JST: Mon 2 Oct 2017 00:00:00
	setNow(timeNow)

	testFilterEval(t, "due before: 10/2/2017", todoist.Item{Due: due("Sun 1 Oct 2017 15:00:00 +0000")}, false)      // JST: Mon 2 Oct 2017 00:00:00
	testFilterEval(t, "due before: 10/2/2017", todoist.Item{Due: due("Sun 1 Oct 2017 14:59:59 +0000")}, true)       // JST: Sun 1 Oct 2017 23:59:59
	testFilterEval(t, "due before: 10/2/2017 13:00", todoist.Item{Due: due("Mon 2 Oct 2017 4:00:00 +0000")}, false) // JST: Mon 2 Oct 2017 13:00:00
	testFilterEval(t, "due before: 10/2/2017 13:00", todoist.Item{Due: due("Mon 2 Oct 2017 3:59:00 +0000")}, true)  // JST: Mon 2 Oct 2017 12:59:00

	testFilterEval(t, "due before: 10/2/2017 13:00", todoist.Item{Due: nil}, false) // JST: Mon 2 Oct 2017 12:59:00
}

func TestOverDueEval(t *testing.T) {
	timeNow := time.Date(2017, time.October, 2, 12, 0, 0, 0, testTimeZone) // JST: Mon 2 Oct 2017 12:00:00
	setNow(timeNow)

	testFilterEval(t, "over due", todoist.Item{Due: due("Mon 2 Oct 2017 2:59:00 +0000")}, true)  // JST: Mon 2 Oct 2017 11:59:00
	testFilterEval(t, "over due", todoist.Item{Due: due("Mon 2 Oct 2017 3:00:00 +0000")}, false) // JST: Mon 2 Oct 2017 12:00:00
	testFilterEval(t, "od", todoist.Item{Due: due("Mon 2 Oct 2017 2:59:00 +0000")}, true)        // JST: Mon 2 Oct 2017 11:59:00
	testFilterEval(t, "od", todoist.Item{Due: due("Mon 2 Oct 2017 3:00:00 +0000")}, false)       // JST: Mon 2 Oct 2017 12:00:00

	testFilterEval(t, "od", todoist.Item{Due: nil}, false) // JST: Mon 2 Oct 2017 12:00:00
}

func TestDueAfterEval(t *testing.T) {
	timeNow := time.Date(2017, time.October, 2, 1, 0, 0, 0, testTimeZone) // JST: Mon 2 Oct 2017 00:00:00
	setNow(timeNow)

	testFilterEval(t, "due after: 10/2/2017", todoist.Item{Due: due("Mon 2 Oct 2017 14:59:59 +0000")}, false)      // JST: Mon 2 Oct 2017 23:59:59
	testFilterEval(t, "due after: 10/2/2017", todoist.Item{Due: due("Mon 2 Oct 2017 15:00:00 +0000")}, true)       // JST: Tue 3 Oct 2017 00:00:00
	testFilterEval(t, "due after: 10/2/2017 13:00", todoist.Item{Due: due("Mon 2 Oct 2017 4:00:00 +0000")}, false) // JST: Mon 2 Oct 2017 13:00:00
	testFilterEval(t, "due after: 10/2/2017 13:00", todoist.Item{Due: due("Mon 2 Oct 2017 4:01:00 +0000")}, true)  // JST: Mon 2 Oct 2017 13:01:00

	testFilterEval(t, "due after: 10/2/2017 13:00", todoist.Item{Due: nil}, false) // JST: Mon 2 Oct 2017 13:01:00
}

func TestIsRecurring(t *testing.T) {
	testFilterEval(t, "recurring", todoist.Item{}, false)
	testFilterEval(t, "recurring", todoist.Item{Due: &todoist.Due{}}, false)
	testFilterEval(t, "recurring", todoist.Item{Due: &todoist.Due{IsRecurring: true}}, true)
}

func TestSubtaskEval(t *testing.T) {
	testFilterEval(t, "subtask", todoist.Item{}, false)
	p := "1234"
	testFilterEval(t, "subtask", todoist.Item{HaveParentID: todoist.HaveParentID{ParentID: &p}}, true)
}

func TestWeekdayEval(t *testing.T) {
	testFilterEval(t, "Tuesday", todoist.Item{Due: due("Tue 3 Oct 2017 00:00:00 +0000")}, true)
	testFilterEval(t, "Tuesday", todoist.Item{Due: due("Mon 2 Oct 2017 00:00:00 +0000")}, false)
	testFilterEval(t, "Tuesday", todoist.Item{}, false)
}

func TestSearchEval(t *testing.T) {
	testFilterEval(t, "search: work on that project", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that project"}}, true)
	testFilterEval(t, "search: work on that other project", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that project"}}, false)

	// partial should also work
	testFilterEval(t, "search: project", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that project"}}, true)
	testFilterEval(t, "search: work", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that project"}}, true)

	// wilcards should work
	testFilterEval(t, "search: work on*project", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that project"}}, true)
	testFilterEval(t, "search: work on*project", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that other project"}}, true)
	testFilterEval(t, "search: work on * other project", todoist.Item{BaseItem: todoist.BaseItem{Content: "work on that project"}}, false)
}

func TestPersonExprEval(t *testing.T) {
	user := todoist.User{
		ID:    "1",
		Email: "alice@example.com",
	}

	collaborators := todoist.Collaborators{
		{
			ID:       "1",
			Email:    "alice@example.com",
			FullName: "Alice",
		},
		{
			ID:       "2",
			Email:    "bob@example.com",
			FullName: "Bob",
		},
		{
			ID:       "3",
			FullName: "Definitely a Name",
		},
	}

	// assigned to
	testFilterEvalWithCollaborators(t, "assigned to: me", todoist.Item{ResponsibleUID: "1"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned to: Me", todoist.Item{ResponsibleUID: "1"}, user, collaborators, true)

	testFilterEvalWithCollaborators(t, "assigned to: Me", todoist.Item{ResponsibleUID: "2"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned to: Others", todoist.Item{ResponsibleUID: "2"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned to: others", todoist.Item{ResponsibleUID: "2"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned to: Others", todoist.Item{ResponsibleUID: "1"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned to: Alice", todoist.Item{ResponsibleUID: "1"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned to: Bob", todoist.Item{ResponsibleUID: "1"}, user, collaborators, false)
	testFilterEvalWithCollaborators(t, "assigned to: Bob", todoist.Item{ResponsibleUID: "1"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned to: bob@example.com", todoist.Item{ResponsibleUID: "2"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned to: alice@example.com", todoist.Item{ResponsibleUID: "2"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned to: Definitely a Name", todoist.Item{ResponsibleUID: "3"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned to: Definitely a Name", todoist.Item{ResponsibleUID: "2"}, user, collaborators, false)

	// assigned by
	testFilterEvalWithCollaborators(t, "assigned by: me", todoist.Item{AssignedByUID: "1"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned by: Me", todoist.Item{AssignedByUID: "1"}, user, collaborators, true)

	testFilterEvalWithCollaborators(t, "assigned by: Me", todoist.Item{AssignedByUID: "2"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned by: Others", todoist.Item{AssignedByUID: "2"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned by: others", todoist.Item{AssignedByUID: "2"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned by: Others", todoist.Item{AssignedByUID: "1"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned by: Alice", todoist.Item{AssignedByUID: "1"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned by: Bob", todoist.Item{AssignedByUID: "1"}, user, collaborators, false)
	testFilterEvalWithCollaborators(t, "assigned by: Bob", todoist.Item{AssignedByUID: "1"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned by: bob@example.com", todoist.Item{AssignedByUID: "2"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned by: alice@example.com", todoist.Item{AssignedByUID: "2"}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "assigned by: Definitely a Name", todoist.Item{AssignedByUID: "3"}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "assigned by: Definitely a Name", todoist.Item{AssignedByUID: "2"}, user, collaborators, false)

	// added by
	testFilterEvalWithCollaborators(t, "added by: me", todoist.Item{BaseItem: todoist.BaseItem{UserID: "1"}}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "added by: Me", todoist.Item{BaseItem: todoist.BaseItem{UserID: "1"}}, user, collaborators, true)

	testFilterEvalWithCollaborators(t, "added by: Me", todoist.Item{BaseItem: todoist.BaseItem{UserID: "2"}}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "added by: Others", todoist.Item{BaseItem: todoist.BaseItem{UserID: "2"}}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "added by: others", todoist.Item{BaseItem: todoist.BaseItem{UserID: "2"}}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "added by: Others", todoist.Item{BaseItem: todoist.BaseItem{UserID: "1"}}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "added by: Alice", todoist.Item{BaseItem: todoist.BaseItem{UserID: "1"}}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "added by: Bob", todoist.Item{BaseItem: todoist.BaseItem{UserID: "1"}}, user, collaborators, false)
	testFilterEvalWithCollaborators(t, "added by: Bob", todoist.Item{BaseItem: todoist.BaseItem{UserID: "1"}}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "added by: bob@example.com", todoist.Item{BaseItem: todoist.BaseItem{UserID: "2"}}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "added by: alice@example.com", todoist.Item{BaseItem: todoist.BaseItem{UserID: "2"}}, user, collaborators, false)

	testFilterEvalWithCollaborators(t, "added by: Definitely a Name", todoist.Item{BaseItem: todoist.BaseItem{UserID: "3"}}, user, collaborators, true)
	testFilterEvalWithCollaborators(t, "added by: Definitely a Name", todoist.Item{BaseItem: todoist.BaseItem{UserID: "2"}}, user, collaborators, false)
}

func TestWildcardProject(t *testing.T) {
	projects := todoist.Projects{
		todoist.Project{
			HaveID: todoist.HaveID{ID: "1"},
			Name:   "baseball",
		},
		todoist.Project{
			HaveID:       todoist.HaveID{ID: "2"},
			HaveParentID: todoist.HaveParentID{ParentID: &[]string{"1"}[0]},
			Name:         "other",
		},
	}

	item1 := todoist.Item{}
	item1.ProjectID = "1"

	item2 := todoist.Item{}
	item2.ProjectID = "2"

	testFilterEvalWithProject(t, "#baseball", item1, projects, true)
	testFilterEvalWithProject(t, "#*ball", item1, projects, true)
	testFilterEvalWithProject(t, "#base*", item1, projects, true)
	testFilterEvalWithProject(t, "#b*all", item1, projects, true)

	testFilterEvalWithProject(t, "#baseball", item2, projects, false)
	testFilterEvalWithProject(t, "#*ball", item2, projects, false)
	testFilterEvalWithProject(t, "#base*", item2, projects, false)
	testFilterEvalWithProject(t, "#b*ll", item2, projects, false)
}
