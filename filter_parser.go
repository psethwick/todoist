// Code generated by goyacc -o filter_parser.go filter_parser.y. DO NOT EDIT.

//line filter_parser.y:2
package main

import __yyfmt__ "fmt"

//line filter_parser.y:2

import (
	"strings"
	"time"
)

//line filter_parser.y:9
type yySymType struct {
	yys   int
	token Token
	expr  Expression
}

const BY = 57346
const TO = 57347
const ADDED = 57348
const ASSIGNED = 57349
const SUBTASK = 57350
const SHARED = 57351
const STRING = 57352
const NUMBER = 57353
const NEXT = 57354
const MONTH_IDENT = 57355
const TWELVE_CLOCK_IDENT = 57356
const HOURS = 57357
const PRIORITY = 57358
const RECURRING = 57359
const TODAY_IDENT = 57360
const TOMORROW_IDENT = 57361
const YESTERDAY_IDENT = 57362
const DAYS = 57363
const VIEW = 57364
const ALL = 57365
const DUE = 57366
const CREATED = 57367
const BEFORE = 57368
const AFTER = 57369
const OVER = 57370
const OVERDUE = 57371
const NO = 57372
const DATE = 57373
const TIME = 57374
const LABELS = 57375
const SEARCH = 57376
const ORDINAL = 57377
const WEEKDAY = 57378
const YEAR_NUMBER = 57379

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"BY",
	"TO",
	"ADDED",
	"ASSIGNED",
	"SUBTASK",
	"SHARED",
	"STRING",
	"NUMBER",
	"NEXT",
	"MONTH_IDENT",
	"TWELVE_CLOCK_IDENT",
	"HOURS",
	"PRIORITY",
	"RECURRING",
	"TODAY_IDENT",
	"TOMORROW_IDENT",
	"YESTERDAY_IDENT",
	"DAYS",
	"VIEW",
	"ALL",
	"DUE",
	"CREATED",
	"BEFORE",
	"AFTER",
	"OVER",
	"OVERDUE",
	"NO",
	"DATE",
	"TIME",
	"LABELS",
	"SEARCH",
	"ORDINAL",
	"WEEKDAY",
	"YEAR_NUMBER",
	"'#'",
	"'@'",
	"'\\\\'",
	"'&'",
	"'*'",
	"'/'",
	"','",
	"'.'",
	"'|'",
	"'('",
	"')'",
	"'!'",
	"':'",
	"'-'",
	"'+'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line filter_parser.y:409

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 171

var yyAct = [...]int8{
	38, 2, 23, 24, 84, 50, 51, 77, 124, 52,
	53, 54, 55, 107, 104, 82, 84, 63, 65, 35,
	34, 39, 103, 80, 102, 71, 41, 42, 43, 87,
	90, 100, 99, 68, 89, 91, 92, 88, 69, 48,
	83, 64, 49, 49, 47, 81, 96, 69, 44, 93,
	94, 95, 83, 71, 71, 71, 71, 36, 48, 33,
	44, 49, 72, 47, 123, 73, 126, 101, 45, 125,
	70, 97, 105, 46, 106, 111, 114, 45, 113, 70,
	57, 32, 46, 98, 62, 74, 61, 45, 60, 37,
	109, 108, 46, 115, 127, 59, 56, 58, 112, 110,
	86, 116, 117, 118, 85, 119, 71, 120, 121, 79,
	78, 122, 75, 76, 67, 66, 25, 71, 9, 8,
	22, 71, 71, 71, 30, 19, 18, 17, 36, 35,
	34, 39, 7, 6, 10, 14, 41, 42, 43, 12,
	13, 31, 16, 15, 40, 1, 28, 29, 11, 0,
	0, 0, 20, 0, 21, 0, 26, 27, 45, 0,
	37, 5, 0, 46, 0, 3, 0, 4, 0, 33,
	44,
}

var yyPact = [...]int16{
	118, -1000, 17, 118, 118, 47, 47, 47, 47, -1000,
	-1000, 64, -1000, 63, -1000, 58, -9, -1000, -1000, 110,
	-17, -1000, -1000, -1000, 37, 18, 27, -1000, 61, -1000,
	108, -4, -1000, 99, 98, 2, -1000, -1000, -1000, 93,
	-1000, -1000, -1000, -1000, 89, -12, -1000, 118, 118, 118,
	-2, -1000, 37, 28, 37, 37, -1000, -1000, -1000, -1000,
	52, -1000, -18, -19, 8, -26, -28, -36, 47, -1000,
	-1000, -1000, 118, -1000, -1000, -37, -1000, -10, 70, 69,
	-1000, 88, 38, 87, -1000, 41, 78, -1000, -1000, -1000,
	-1000, -1000, -1000, -1, -1, -1000, -1000, 47, -1000, 8,
	8, -1000, 8, 47, 47, 37, -1000, 47, -1000, -1000,
	21, -1000, -42, -1000, 32, -1000, 37, -1000, -1000, -1000,
	37, 37, 37, 29, 83, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 145, 1, 2, 144, 141, 139, 134, 133, 132,
	81, 120, 119, 118, 3, 0, 116,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 16, 16, 15, 15, 15, 15, 15, 15,
	15, 14, 14, 14, 14, 14, 14, 11, 11, 11,
	9, 8, 12, 13, 7, 7, 6, 6, 3, 3,
	3, 3, 3, 3, 5, 5, 5, 5, 5, 5,
	5, 5, 4, 4, 4, 4, 10, 10, 10, 10,
}

var yyR2 = [...]int8{
	0, 0, 1, 3, 3, 3, 2, 2, 2, 4,
	2, 2, 1, 1, 2, 2, 1, 2, 1, 4,
	4, 3, 4, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 3, 3, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 2, 2, 2, 4, 4, 4,
	2, 1, 1, 2, 2, 3, 2, 1, 2, 1,
	1, 3, 3, 2, 5, 3, 4, 3, 1, 1,
	1, 1, 2, 3, 2, 3, 3, 5, 3, 2,
}

var yyChk = [...]int16{
	-1000, -1, -2, 47, 49, 43, -8, -9, -12, -13,
	-7, 30, -6, 22, 17, 25, 24, 9, 8, 7,
	34, 36, -11, -3, -14, -16, 38, 39, 28, 29,
	6, -5, -10, 51, 12, 11, 10, 42, -15, 13,
	-4, 18, 19, 20, 52, 40, 45, 46, 41, 44,
	-2, -2, -14, -14, -14, -14, 32, 16, 33, 31,
	24, 23, 26, 26, 50, 27, 5, 4, 50, 10,
	42, -15, 44, 38, 24, 4, -10, 11, 11, 11,
	21, 43, 13, 50, 14, 11, 11, 41, 49, 46,
	42, 47, 48, -2, -2, -2, 48, 43, 31, 50,
	50, -3, 50, 50, 50, -14, -2, 50, 21, 21,
	11, 37, 11, 37, 35, 15, -14, -3, -3, -3,
	-14, -14, -14, 43, 50, 37, 37, 11,
}

var yyDef = [...]int8{
	1, -2, 2, 0, 0, 0, 0, 0, 0, 12,
	13, 0, 16, 0, 18, 0, 0, 23, 24, 25,
	0, 27, 28, 29, 30, 31, 51, 52, 0, 57,
	0, 59, 60, 0, 0, 0, 41, 42, 43, 0,
	68, 69, 70, 71, 0, 0, 40, 0, 0, 0,
	0, 6, 7, 8, 10, 11, 14, 15, 53, 54,
	0, 17, 0, 0, 0, 0, 0, 0, 0, 44,
	45, 46, 0, 50, 56, 0, 58, 0, 0, 0,
	63, 0, 74, 0, 79, 72, 0, 34, 35, 36,
	37, 38, 39, 3, 4, 33, 5, 0, 55, 0,
	0, 21, 0, 0, 0, 26, 32, 0, 61, 62,
	75, 67, 76, 65, 73, 78, 9, 19, 20, 22,
	47, 48, 49, 0, 0, 66, 64, 77,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 49, 3, 38, 3, 3, 41, 3,
	47, 48, 42, 52, 44, 51, 45, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 50, 3,
	3, 3, 3, 3, 39, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 40, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 46,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line filter_parser.y:42
		{
			yyVAL.expr = VoidExpr{}
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:46
		{
			yylex.(*Lexer).result = yyDollar[1].expr
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:52
		{
			yyVAL.expr = BoolInfixOpExpr{left: yyDollar[1].expr, operator: '|', right: yyDollar[3].expr}
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:56
		{
			yyVAL.expr = BoolInfixOpExpr{left: yyDollar[1].expr, operator: '&', right: yyDollar[3].expr}
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:60
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:64
		{
			yyVAL.expr = NotOpExpr{expr: yyDollar[2].expr}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:68
		{
			yyVAL.expr = ProjectExpr{isAll: false, section: yyDollar[2].expr.(StringExpr).String()}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:72
		{
			yyVAL.expr = ProjectExpr{isAll: false, name: yyDollar[2].expr.(StringExpr).String()}
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:76
		{
			name := yyDollar[2].expr.(StringExpr).String()
			section := yyDollar[4].expr.(StringExpr).String()
			yyVAL.expr = ProjectExpr{isAll: false, name: name, section: section}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:82
		{
			yyVAL.expr = ProjectExpr{isAll: true, name: yyDollar[2].expr.(StringExpr).String()}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:86
		{
			yyVAL.expr = LabelExpr{name: yyDollar[2].expr.(StringExpr).String()}
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:90
		{
			yyVAL.expr = LabelExpr{name: ""}
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:94
		{
			yyVAL.expr = DateExpr{operation: NO_DUE_DATE}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:98
		{
			yyVAL.expr = DateExpr{operation: NO_TIME}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:102
		{
			yyVAL.expr = NoPriorityExpr{}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:106
		{
			yyVAL.expr = DateExpr{allDay: false, datetime: now(), operation: DUE_BEFORE}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:110
		{
			yyVAL.expr = ViewAllExpr{}
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:114
		{
			yyVAL.expr = RecurringExpr{}
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:118
		{
			e := yyDollar[4].expr.(DateExpr)
			e.operation = CREATED_BEFORE
			yyVAL.expr = e
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:124
		{
			e := yyDollar[4].expr.(DateExpr)
			e.operation = DUE_BEFORE
			yyVAL.expr = e
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:130
		{
			e := yyDollar[3].expr.(DateExpr)
			e.operation = DUE_ON
			yyVAL.expr = e
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:136
		{
			e := yyDollar[4].expr.(DateExpr)
			e.operation = DUE_AFTER
			yyVAL.expr = e
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:142
		{
			yyVAL.expr = SharedExpr{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:146
		{
			yyVAL.expr = SubtaskExpr{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:150
		{
			yyVAL.expr = AssignedExpr{}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:154
		{
			yyVAL.expr = SearchExpr{keyword: yyDollar[3].expr.(StringExpr).String()}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:158
		{
			yyVAL.expr = WeekdayExpr{day: WeekdayHash[strings.ToLower(yyDollar[1].token.literal)]}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:168
		{
			l := yyDollar[1].expr.(ListExpr)
			l.exprs = append(l.exprs, yyDollar[3].expr)
			yyVAL.expr = l
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:174
		{
			yyVAL.expr = ListExpr{exprs: []Expression{yyDollar[1].expr, yyDollar[3].expr}}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:180
		{
			yyVAL.expr = "&"
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:184
		{
			yyVAL.expr = "!"
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:188
		{
			yyVAL.expr = `\|`
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:192
		{
			yyVAL.expr = `\*`
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:196
		{
			yyVAL.expr = `\(`
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:200
		{
			yyVAL.expr = `\)`
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:204
		{
			yyVAL.expr = `\.`
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:210
		{
			yyVAL.expr = NewStringExpr(yyDollar[1].token.literal)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:214
		{
			yyVAL.expr = NewStringExpr(".*")
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:218
		{
			yyVAL.expr = NewStringExpr(yyDollar[1].expr.(string))
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:222
		{
			yyVAL.expr = yyDollar[1].expr.(StringExpr).Add(NewStringExpr(yyDollar[2].token.literal))
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:226
		{
			yyVAL.expr = yyDollar[1].expr.(StringExpr).Add(NewStringExpr(".*"))
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:230
		{
			yyVAL.expr = yyDollar[1].expr.(StringExpr).Add(NewStringExpr(yyDollar[2].expr.(string)))
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:237
		{
			yyVAL.expr = PersonExpr{operation: ASSIGNED_TO, person: yyDollar[4].expr.(StringExpr).String()}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:241
		{
			yyVAL.expr = PersonExpr{operation: ASSIGNED_BY, person: yyDollar[4].expr.(StringExpr).String()}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:245
		{
			yyVAL.expr = PersonExpr{operation: ADDED_BY, person: yyDollar[4].expr.(StringExpr).String()}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:251
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:257
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:263
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:269
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:275
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:279
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:285
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:289
		{
			yyVAL.expr = yyDollar[1].token
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:296
		{
			date := yyDollar[1].expr.(time.Time)
			time := yyDollar[2].expr.(time.Duration)
			yyVAL.expr = DateExpr{allDay: false, datetime: date.Add(time)}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:302
		{
			yyVAL.expr = DateExpr{allDay: true, datetime: yyDollar[1].expr.(time.Time)}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:306
		{
			nd := now().Sub(today())
			d := yyDollar[1].expr.(time.Duration)
			if d <= nd {
				d = d + time.Duration(int64(time.Hour)*24)
			}
			yyVAL.expr = DateExpr{allDay: false, datetime: today().Add(d)}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:315
		{
			date := today().AddDate(0, 0, -atoi(yyDollar[2].token.literal))
			yyVAL.expr = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:320
		{
			date := today().AddDate(0, 0, atoi(yyDollar[2].token.literal))
			yyVAL.expr = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:325
		{
			date := today().AddDate(0, 0, atoi(yyDollar[1].token.literal))
			yyVAL.expr = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
//line filter_parser.y:332
		{
			yyVAL.expr = time.Date(atoi(yyDollar[5].token.literal), time.Month(atoi(yyDollar[1].token.literal)), atoi(yyDollar[3].token.literal), 0, 0, 0, 0, timezone())
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:336
		{
			yyVAL.expr = time.Date(atoi(yyDollar[3].token.literal), MonthIdentHash[strings.ToLower(yyDollar[1].token.literal)], atoi(yyDollar[2].token.literal), 0, 0, 0, 0, timezone())
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
//line filter_parser.y:340
		{
			yyVAL.expr = time.Date(atoi(yyDollar[4].token.literal), MonthIdentHash[strings.ToLower(yyDollar[1].token.literal)], atoi(yyDollar[2].token.literal), 0, 0, 0, 0, timezone())
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:344
		{
			yyVAL.expr = time.Date(atoi(yyDollar[3].token.literal), MonthIdentHash[strings.ToLower(yyDollar[2].token.literal)], atoi(yyDollar[1].token.literal), 0, 0, 0, 0, timezone())
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:348
		{
			tod := today()
			date := yyDollar[1].expr.(time.Time)
			if date.Before(tod) {
				date = date.AddDate(1, 0, 0)
			}
			yyVAL.expr = date
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:357
		{
			yyVAL.expr = today()
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:361
		{
			yyVAL.expr = today().AddDate(0, 0, 1)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line filter_parser.y:365
		{
			yyVAL.expr = today().AddDate(0, 0, -1)
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:371
		{
			yyVAL.expr = time.Date(today().Year(), MonthIdentHash[strings.ToLower(yyDollar[1].token.literal)], atoi(yyDollar[2].token.literal), 0, 0, 0, 0, timezone())
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:375
		{
			yyVAL.expr = time.Date(today().Year(), MonthIdentHash[strings.ToLower(yyDollar[1].token.literal)], atoi(yyDollar[2].token.literal), 0, 0, 0, 0, timezone())
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:379
		{
			yyVAL.expr = time.Date(today().Year(), MonthIdentHash[strings.ToLower(yyDollar[2].token.literal)], atoi(yyDollar[1].token.literal), 0, 0, 0, 0, timezone())
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:383
		{
			yyVAL.expr = time.Date(now().Year(), time.Month(atoi(yyDollar[3].token.literal)), atoi(yyDollar[1].token.literal), 0, 0, 0, 0, timezone())
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:389
		{
			yyVAL.expr = time.Duration(int64(time.Hour)*int64(atoi(yyDollar[1].token.literal)) + int64(time.Minute)*int64(atoi(yyDollar[3].token.literal)))
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
//line filter_parser.y:393
		{
			yyVAL.expr = time.Duration(int64(time.Hour)*int64(atoi(yyDollar[1].token.literal)) + int64(time.Minute)*int64(atoi(yyDollar[3].token.literal)) + int64(time.Second)*int64(atoi(yyDollar[5].token.literal)))
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line filter_parser.y:397
		{
			yyVAL.expr = time.Duration(int64(time.Hour) * int64(atoi(yyDollar[2].token.literal)))
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line filter_parser.y:401
		{
			hour := atoi(yyDollar[1].token.literal)
			if TwelveClockIdentHash[yyDollar[2].token.literal] {
				hour = hour + 12
			}
			yyVAL.expr = time.Duration(int64(time.Hour) * int64(hour))
		}
	}
	goto yystack /* stack new state and value */
}
