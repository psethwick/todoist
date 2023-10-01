package main

import (
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

type (
	Expression interface{}
	Token      struct {
		token   int
		literal string
	}
)

type ErrorExpr struct {
	error string
}

type VoidExpr struct{}

type ViewAllExpr struct{}

type NoPriorityExpr struct{}

type AssignedExpr struct{}

type RecurringExpr struct{}

type SubtaskExpr struct{}

type SharedExpr struct{}

type SearchExpr struct {
	keyword string
}

type StringExpr struct {
	words []string
}

func NewStringExpr(strs ...string) StringExpr {
	return StringExpr{
		words: strs,
	}
}

func (se StringExpr) Add(other StringExpr) StringExpr {
	se.words = append(se.words, other.words...)
	return se
}

func (se StringExpr) String() string {
	return strings.Join(se.words, "\\s?")
}

type personOperation int

const (
	ASSIGNED_TO personOperation = iota
	ASSIGNED_BY
	ADDED_BY
)

// A collaborator can be identified by:
//
//	The person’s email
//	The person’s full name
//	“Me” (referring to yourself)
//	“Others” (referring to all users other than yourself)
type PersonExpr struct {
	person    string
	operation personOperation
}

type BoolInfixOpExpr struct {
	left     Expression
	operator rune
	right    Expression
}

type ProjectExpr struct {
	isAll   bool
	name    string
	section string
}

type LabelExpr struct {
	name string
}

type NotOpExpr struct {
	expr Expression
}

type dateOperation int

const (
	DUE_ON dateOperation = iota
	DUE_BEFORE
	DUE_AFTER
	NO_DUE_DATE
	NO_TIME
	CREATED_BEFORE
)

type WeekdayExpr struct {
	day time.Weekday
}

type DateExpr struct {
	operation dateOperation
	datetime  time.Time
	allDay    bool
}

func atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

var (
	now   = time.Now
	today = func() time.Time {
		return time.Date(now().Year(), now().Month(), now().Day(), 0, 0, 0, 0, now().Location())
	}
)

var timezone = func() *time.Location {
	return now().Location()
}

type Lexer struct {
	scanner.Scanner
	input  string
	last   int
	result []Expression
}

var MonthIdentHash = map[string]time.Month{
	"jan":  time.January,
	"feb":  time.February,
	"mar":  time.March,
	"apr":  time.April,
	"may":  time.May,
	"june": time.June,
	"july": time.July,
	"aug":  time.August,
	"sept": time.September,
	"oct":  time.October,
	"nov":  time.November,
	"dec":  time.December,

	"january":   time.January,
	"february":  time.February,
	"march":     time.March,
	"april":     time.April,
	"august":    time.August,
	"september": time.September,
	"october":   time.October,
	"november":  time.November,
	"december":  time.December,
}

var WeekdayHash = map[string]time.Weekday{
	"sunday":    time.Sunday,
	"monday":    time.Monday,
	"tuesday":   time.Tuesday,
	"wednesday": time.Wednesday,
	"thursday":  time.Thursday,
	"friday":    time.Friday,
	"saturday":  time.Saturday,
}

var OrdinalHash = map[string]bool{
	"st": true,
	"nd": true,
	"rd": true,
	"th": true,
}

var TwelveClockIdentHash = map[string]bool{
	"am": false,
	"pm": true,
}

var TodayIdentHash = map[string]bool{
	"today": true,
	"tod":   true,
}

var TomorrowIdentHash = map[string]bool{
	"tomorrow": true,
	"tom":      true,
}

var OverDueHash = map[string]bool{
	"overdue": true,
	"od":      true,
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	switch token {
	case scanner.Ident:
		lowerToken := strings.ToLower(l.TokenText())
		if _, ok := MonthIdentHash[lowerToken]; ok {
			token = MONTH_IDENT
		} else if _, ok := TwelveClockIdentHash[lowerToken]; ok {
			if l.last == NUMBER {
				token = TWELVE_CLOCK_IDENT
			} else {
				token = STRING
			}
		} else if _, ok := TodayIdentHash[lowerToken]; ok {
			token = TODAY_IDENT
		} else if _, ok := TomorrowIdentHash[lowerToken]; ok {
			token = TOMORROW_IDENT
		} else if _, ok := OverDueHash[lowerToken]; ok {
			token = OVERDUE
		} else if _, ok := OrdinalHash[lowerToken]; ok {
			if l.last == NUMBER {
				token = ORDINAL
			} else {
				token = STRING
			}
		} else if _, ok := WeekdayHash[lowerToken]; ok {
			token = WEEKDAY
		} else if lowerToken == "yesterday" {
			token = YESTERDAY_IDENT
		} else if lowerToken == "due" {
			token = DUE
		} else if lowerToken == "next" {
			token = NEXT
		} else if lowerToken == "before" {
			if l.last == DUE || l.last == CREATED {
				token = BEFORE
			} else {
				token = STRING
			}
		} else if lowerToken == "after" {
			if l.last == DUE {
				token = AFTER
			} else {
				token = STRING
			}
		} else if lowerToken == "over" {
			token = l.NextMatches(`due`, OVER, STRING)
		} else if lowerToken == "no" {
			token = l.NextMatches(`date|time|due date|labels|priority`, NO, STRING)
		} else if lowerToken == "time" {
			if l.last == NO {
				token = TIME
			} else {
				token = STRING
			}
		} else if lowerToken == "date" {
			if l.last == NO || l.last == DUE {
				token = DATE
			} else {
				token = STRING
			}
		} else if lowerToken == "labels" {
			if l.last == NO {
				token = LABELS
			} else {
				token = STRING
			}
		} else if lowerToken == "view" {
			token = l.NextMatches(`all`, VIEW, STRING)
		} else if lowerToken == "all" {
			if l.last == VIEW {
				token = ALL
			} else {
				token = STRING
			}
		} else if lowerToken == "added" {
			token = l.NextMatches(`by`, ADDED, STRING)
		} else if lowerToken == "assigned" {
			token = ASSIGNED
		} else if lowerToken == "to" {
			if l.last == ASSIGNED {
				token = TO
			} else {
				token = STRING
			}
		} else if lowerToken == "by" {
			if l.last == ASSIGNED || l.last == ADDED {
				token = BY
			} else {
				token = STRING
			}
		} else if lowerToken == "days" {
			token = DAYS
		} else if lowerToken == "hours" {
			token = HOURS
		} else if lowerToken == "shared" {
			token = SHARED
		} else if lowerToken == "created" {
			token = l.NextMatches(`before`, CREATED, STRING)
		} else if lowerToken == "subtask" {
			token = SUBTASK
		} else if lowerToken == "priority" {
			if l.last == NO {
				token = PRIORITY
			} else {
				token = STRING
			}
		} else if lowerToken == "recurring" {
			token = RECURRING
		} else if lowerToken == "search" {
			token = SEARCH
		} else {
			token = STRING
		}
	case int('@'):
		if l.IsInString() {
			token = STRING
		} else {
			token = AT
		}
	case scanner.Int:
		i := atoi(l.TokenText())
		if i > 2000 {
			token = YEAR_NUMBER
		} else {
			token = NUMBER
		}
	}
	lval.token = Token{token: token, literal: l.TokenText()}
	l.last = token
	return token
}

func (l *Lexer) IsInString() bool {
	if l.last != STRING {
		return false
	}
	offset := l.Pos().Offset
	if offset == 0 {
		return false
	}
	return l.input[offset-1] != ' '
}

func (l *Lexer) NextMatches(s string, pass int, fail int) int {
	offset := l.Pos().Offset
	match, _ := regexp.MatchString(s, l.input[offset:])
	if match {
		return pass
	} else {
		return fail
	}
}

func (l *Lexer) Error(e string) {
	l.result = append(l.result, ErrorExpr{e})
}

func Filter(f string) (e []Expression) {
	l := new(Lexer)
	l.Init(strings.NewReader(f))
	l.input = f
	// exclude scanner.ScanFloats because afternoon times in am/pm format trigger float parsing
	l.Mode = scanner.ScanIdents | scanner.ScanInts
	yyParse(l)
	return l.result
}
