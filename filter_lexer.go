package main

import (
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

type Expression interface{}
type Token struct {
	token   int
	literal string
}

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
	isAll bool
	name  string
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

type DateExpr struct {
	operation dateOperation
	datetime  time.Time
	allDay    bool
}

func atoi(a string) (i int) {
	i, _ = strconv.Atoi(a)
	return
}

var now = time.Now
var today = func() time.Time {
	return time.Date(now().Year(), now().Month(), now().Day(), 0, 0, 0, 0, now().Location())
}
var timezone = func() *time.Location {
	return now().Location()
}

type Lexer struct {
	scanner.Scanner
	result Expression
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
			token = TWELVE_CLOCK_IDENT
		} else if _, ok := TodayIdentHash[lowerToken]; ok {
			token = TODAY_IDENT
		} else if _, ok := TomorrowIdentHash[lowerToken]; ok {
			token = TOMORROW_IDENT
		} else if _, ok := OverDueHash[lowerToken]; ok {
			token = OVERDUE
		} else if lowerToken == "yesterday" {
			token = YESTERDAY_IDENT
		} else if lowerToken == "due" {
			token = DUE
		} else if lowerToken == "next" {
			token = NEXT
		} else if lowerToken == "before" {
			token = BEFORE
		} else if lowerToken == "after" {
			token = AFTER
		} else if lowerToken == "over" {
			token = OVER
		} else if lowerToken == "no" {
			token = NO
		} else if lowerToken == "time" {
			token = TIME
		} else if lowerToken == "date" {
			token = DATE
		} else if lowerToken == "labels" {
			token = LABELS
		} else if lowerToken == "view" {
			token = VIEW
		} else if lowerToken == "all" {
			token = ALL
		} else if lowerToken == "to" {
			token = TO
		} else if lowerToken == "by" {
			token = BY
		} else if lowerToken == "added" {
			token = ADDED
		} else if lowerToken == "assigned" {
			token = ASSIGNED
		} else if lowerToken == "days" {
			token = DAYS
		} else if lowerToken == "hours" {
			token = HOURS
		} else if lowerToken == "shared" {
			token = SHARED
		} else if lowerToken == "created" {
			token = CREATED
		} else if lowerToken == "subtask" {
			token = SUBTASK
		} else if lowerToken == "priority" {
			token = PRIORITY
		} else if lowerToken == "recurring" {
			token = RECURRING
		} else if lowerToken == "search" {
			token = SEARCH
		} else {
			token = STRING
		}
	case scanner.Int:
		token = NUMBER
	}
	lval.token = Token{token: token, literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	l.result = ErrorExpr{e}
}

func Filter(f string) (e Expression) {
	l := new(Lexer)
	l.Init(strings.NewReader(f))
	// important to exclude scanner.ScanFloats because afternoon times in am/pm format trigger float parsing
	l.Mode = scanner.ScanIdents | scanner.ScanInts
	yyParse(l)
	return l.result
}
