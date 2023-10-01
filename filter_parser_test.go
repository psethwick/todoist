package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, []Expression{VoidExpr{}}, Filter(""), "they should be equal")
}

func TestPriorityFilter(t *testing.T) {
	assert.Equal(t, []Expression{StringExpr{words: []string{"p1"}}}, Filter("p1"), "they should be equal")
	assert.Equal(t, []Expression{NoPriorityExpr{}}, Filter("No priority"), "they should be equal")
}

func TestProjectFilter(t *testing.T) {
	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  "Work",
		}},
		Filter("#Work"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  `Work\s?work`,
		}},
		Filter("#Work work"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  `Work\s?work\s?work`,
		}},
		Filter("#Work work work"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: true,
			name:  "Work",
		}},
		Filter("##Work"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  `One\s?&\s?Two`,
		}},
		Filter(`#One \& Two`), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  `One\s?\|\s?Two`,
		}},
		Filter(`#One \| Two`), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  `One\s?\|\s?Two`,
		}},
		Filter(`#One\|Two`), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  "Exclam\\s?!",
		}},
		Filter("#Exclam\\!"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  "Or\\s?\\|",
		}},
		Filter("#Or\\|"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  ".*\\s?ball",
		}},
		Filter("#*ball"), "they should be equal")

	assert.Equal(t,
		[]Expression{ProjectExpr{
			isAll: false,
			name:  "base\\s?.*",
		}},
		Filter("#base*"), "they should be equal")
}

func TestLabelFilter(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			LabelExpr{
				name: "Test",
			},
		},
		Filter("@Test"), "they should be equal")
	assert.Equal(t,
		[]Expression{
			LabelExpr{
				name: "",
			},
		},
		Filter("no labels"), "they should be equal")
}

func TestBoolInfixFilter(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			BoolInfixOpExpr{
				left:     StringExpr{words: []string{"p1"}},
				operator: '|',
				right:    StringExpr{words: []string{"p2"}},
			},
		},
		Filter("p1 | p2"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			BoolInfixOpExpr{
				left:     StringExpr{words: []string{"p1"}},
				operator: '&',
				right:    StringExpr{words: []string{"p2"}},
			},
		},
		Filter("p1 & p2"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			BoolInfixOpExpr{
				left:     StringExpr{words: []string{"p1"}},
				operator: '&',
				right: BoolInfixOpExpr{
					left:     StringExpr{words: []string{"p2"}},
					operator: '|',
					right:    StringExpr{words: []string{"p3"}},
				},
			},
		},
		Filter("p1 & (p2 | p3 )"), "they should be equal")
}

func setNow(t time.Time) {
	now = func() time.Time { return t }
}

func TestDateTimeFilter(t *testing.T) {
	timeNow := time.Date(2017, time.January, 2, 1, 0, 0, 0, testTimeZone)
	setNow(timeNow)

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2017, time.October, 5, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("10/5/2017"), "they should be equal")

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), time.January, 3, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("Jan 3"),
		"they should be equal",
	)

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), time.August, 8, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("8 August"),
		"they should be equal",
	)

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2020, time.February, 10, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("10 Feb 2020"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), time.May, 16, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("16/05"), "they should be equal")

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 16, 0, 0, 0, testTimeZone),
				allDay:    false,
			},
		},
		Filter("16:00"),
		"they should be equal",
	)

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 16, 10, 3, 0, testTimeZone),
				allDay:    false,
			},
		},
		Filter("16:10:03"),
		"they should be equal",
	)

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 15, 0, 0, 0, testTimeZone),
				allDay:    false,
			},
		},
		Filter("3pm"),
		"they should be equal",
	)

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 7, 0, 0, 0, testTimeZone),
				allDay:    false,
			},
		},
		Filter("7am"),
		"they should be equal",
	)

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2020, time.February, 10, 15, 0, 0, 0, testTimeZone), allDay: false},
		},
		Filter("10 Feb 2020 3pm"), "they should be equal")

	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 7, 0, 0, 0, testTimeZone),
				allDay:    false,
			},
		},
		Filter("7am"),
		"they should be equal",
	)
}

func TestSpecialDateTimeFilter(t *testing.T) {
	timeNow := time.Date(2017, time.January, 1, 1, 0, 0, 0, testTimeZone)
	setNow(timeNow)
	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2017, time.January, 1, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("today"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2017, time.January, 1, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("tod"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2017, time.January, 1, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("Today"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2017, time.January, 2, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("tomorrow"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2017, time.January, 2, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("tom"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{operation: DUE_ON, datetime: time.Date(2016, time.December, 31, 0, 0, 0, 0, testTimeZone), allDay: true},
		},
		Filter("yesterday"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{
				operation: DUE_BEFORE,
				datetime:  time.Date(2017, time.January, 15, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("14 days"), "they should be equal")

	assert.Equal(t,
		[]Expression{
			DateExpr{
				operation: DUE_BEFORE,
				datetime:  time.Date(2017, time.January, 15, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("next 14 days"), "they should be equal")
}

func TestNoTime(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			DateExpr{operation: NO_TIME},
		},
		Filter("no time"),
	)
}

func TestSubtask(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			SubtaskExpr{},
		},
		Filter("subtask"),
	)
}

func TestShared(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			SharedExpr{},
		},
		Filter("shared"),
	)
}

func TestPerson(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			PersonExpr{operation: ASSIGNED_TO, person: "me"},
		},
		Filter("assigned to: me"),
	)

	assert.Equal(t,
		[]Expression{
			PersonExpr{operation: ASSIGNED_TO, person: `alice\s?@\s?example\s?\.\s?com`},
		},
		Filter("assigned to: alice@example.com"),
	)

	assert.Equal(t,
		[]Expression{
			PersonExpr{operation: ASSIGNED_BY, person: "me"},
		},
		Filter("assigned by: me"),
	)

	assert.Equal(t,
		[]Expression{
			PersonExpr{operation: ASSIGNED_TO, person: "Becky"},
		},
		Filter("assigned to: Becky"),
	)

	assert.Equal(t,
		[]Expression{
			PersonExpr{operation: ADDED_BY, person: "me"},
		},
		Filter("added by: me"),
	)

	assert.Equal(t,
		[]Expression{
			PersonExpr{operation: ADDED_BY, person: "Becky"},
		},
		Filter("added by: Becky"),
	)
}

func TestAssigned(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			AssignedExpr{},
		},
		Filter("assigned"),
	)
}

func TestRecurring(t *testing.T) {
	assert.Equal(t,
		[]Expression{
			RecurringExpr{},
		},
		Filter("recurring"),
	)
}

func TestDateTimeElapsedFilter(t *testing.T) {
	timeNow := time.Date(2017, time.January, 2, 18, 0, 0, 0, testTimeZone)
	setNow(timeNow)
	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day()+1, 16, 0, 0, 0, testTimeZone),
				allDay:    false,
			},
		},
		Filter("16:00"),
		"they should be equal",
	)

	timeNow = time.Date(2017, time.May, 16, 23, 59, 59, 0, testTimeZone)
	setNow(timeNow)
	assert.Equal(t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year(), time.May, 16, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("16/05"), "they should be equal")

	timeNow = time.Date(2017, time.May, 17, 0, 0, 0, 0, testTimeZone)
	setNow(timeNow)
	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				operation: DUE_ON,
				datetime:  time.Date(timeNow.Year()+1, time.May, 16, 0, 0, 0, 0, testTimeZone),
				allDay:    true,
			},
		},
		Filter("16/05"),
		"they should be equal",
	)
}

func TestParseWildcard(t *testing.T) {
	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{name: `.*\s?ball`},
		},
		Filter("#*ball"),
	)

	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{name: `base\s?.*`},
		},
		Filter("#base*"),
	)

	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{name: `b\s?.*\s?ll`},
		},
		Filter("#b*ll"),
	)

	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{name: `b\s?.*\s?all`},
		},
		Filter("#b*all"),
	)
}

func TestWeekday(t *testing.T) {
	assert.Equal(
		t,
		[]Expression{
			WeekdayExpr{day: time.Tuesday},
		},
		Filter("Tuesday"),
	)
}

func TestSections(t *testing.T) {
	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{section: "Meetings"},
		},
		Filter("/Meetings"),
	)

	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{name: "Work", section: "Meetings"},
		},
		Filter("#Work/Meetings"),
	)

	assert.Equal(
		t,
		[]Expression{
			BoolInfixOpExpr{
				left:     ProjectExpr{name: "Work"},
				operator: '&',
				right:    ProjectExpr{section: "Meetings"},
			},
		},
		Filter("#Work & /Meetings"),
	)

	assert.Equal(
		t,
		[]Expression{
			NotOpExpr{expr: ProjectExpr{section: ".*"}},
		},
		Filter("!/*"),
	)

	assert.Equal(
		t,
		[]Expression{
			BoolInfixOpExpr{
				left:     NotOpExpr{expr: ProjectExpr{section: ".*"}},
				operator: '&',
				right:    NotOpExpr{expr: ProjectExpr{name: "Inbox"}},
			},
		},
		Filter("!/* & !#Inbox"),
	)
}

func TestComplexFilter(t *testing.T) {
	timeNow := time.Date(2017, time.January, 2, 18, 0, 0, 0, testTimeZone)
	setNow(timeNow)
	assert.Equal(
		t,
		[]Expression{
			BoolInfixOpExpr{
				left: BoolInfixOpExpr{
					left: BoolInfixOpExpr{
						left: BoolInfixOpExpr{
							left: DateExpr{
								operation: DUE_ON,
								datetime:  time.Date(2017, time.January, 2, 0, 0, 0, 0, testTimeZone),
								allDay:    true,
							},
							operator: '|',
							right: DateExpr{
								operation: DUE_BEFORE,
								datetime:  timeNow,
								allDay:    false,
							},
						},
						operator: '&',
						right: ProjectExpr{
							isAll: true,
							name:  "Inbox",
						},
					},
					operator: '&',
					right: NotOpExpr{
						expr: PersonExpr{person: "others"},
					},
				}, operator: '&',
				right: NotOpExpr{expr: LabelExpr{name: "me"}},
			},
		},
		Filter("(today | overdue) & ##Inbox & (!assigned to: others) & (!@me)"),
	)
}

func TestListExpr(t *testing.T) {
	timeNow := time.Date(2017, time.January, 2, 18, 0, 0, 0, testTimeZone)
	setNow(timeNow)
	assert.Equal(
		t,
		[]Expression{
			DateExpr{
				datetime: time.Date(2017, time.January, 1, 0, 0, 0, 0, testTimeZone),
				allDay:   true,
			},
			DateExpr{
				datetime: time.Date(2017, time.January, 2, 0, 0, 0, 0, testTimeZone),
				allDay:   true,
			},
		},
		Filter("due: yesterday, today"),
	)

	assert.Equal(
		t,
		[]Expression{
			ProjectExpr{
				name: "Work",
			},
			ProjectExpr{
				name: "Play",
			},
			DateExpr{
				datetime: time.Date(2017, time.January, 2, 0, 0, 0, 0, testTimeZone),
				allDay:   true,
			},
		},
		Filter("#Work, #Play, today"),
	)
}

func TestSearchExpr(t *testing.T) {
	assert.Equal(
		t,
		[]Expression{SearchExpr{keyword: "Meeting"}},
		Filter("search: Meeting"),
	)
}

func TestNoSyntaxErrorAllOfficialExamples(t *testing.T) {
	tests := []string{
		"(today | overdue) & #Work",
		"no date",
		"#Work & no due date",
		"Saturday & @night",
		"(P1 | P2) & 14 days",
		"7 days & @waiting",
		"view all",
		"no time",
		"created before: -30 days",
		"created before: -365 days",
		"due before: +8 hours & !overdue",
		"subtask",
		"!subtask",
		"shared & !assigned",
		"assigned to: me",
		"#Work & assigned to: me",
		"assigned by: me",
		"assigned to: me",
		"assigned to: Becky",
		"added by: me",
		"added by: Becky",
		"recurring",
		"No priority",
		"next 5 days",
		"#multiple words",
		"#One \\& Two",
		"@*ball",
		"@home*",
		"assigned to: m* smith",
		"assigned to: alice@example.com",
		"#*Work",
		"Work*",

		"search: Meeting",
		"search: Meeting & today",
		"search: Meeting | search: Work",
		"search: email",
		"search: http",
		"search: http & search:*",

		"Monday",
		"Tuesday",
		"Sunday",
		"10/5/2022",
		"10/5/2022 5pm",
		"today",
		"tomorrow",
		"yesterday",
		"3 days",
		"-3 days",
		"yesterday",

		"Oct 5th 2022",
		"Oct 5th 5pm",

		"#Work",
		"##Work",
		"##School & !#Science",

		"/Meetings",
		"#Work/Meetings",
		"#Work & /Meetings",
		"!/*",
		"!/* & !#Inbox",

		"due: yesterday, today",
	}
	for _, input := range tests {
		for _, e := range Filter(input) {
			if err, ok := e.(ErrorExpr); ok {
				assert.True(t, false, input, err.error)
			}
		}
	}
}
