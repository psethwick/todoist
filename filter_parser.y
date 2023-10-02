%{
package main 

import (
    "time"
    "strings"
)
%}
%union{
    token Token
    expr Expression
    exprs []Expression
}

%type<exprs> filter exprs
%type<expr> expr 
%type<expr> s_datetime s_date s_date_year
%type<expr> s_overdue s_nodate s_project_key s_project_all_key 
%type<expr> s_time s_person s_label_key s_no_labels s_string 
%type<expr> s_special_chars
%token<token> BY TO ADDED ASSIGNED SUBTASK SHARED STRING NUMBER NEXT
%token<token> MONTH_IDENT TWELVE_CLOCK_IDENT HOURS PRIORITY RECURRING
%token<token> TODAY_IDENT TOMORROW_IDENT YESTERDAY_IDENT DAYS VIEW ALL
%token<token> DUE CREATED BEFORE AFTER OVER OVERDUE NO DATE TIME LABELS
%token<token> SEARCH ORDINAL WEEKDAY YEAR_NUMBER AT
%token<token> '#' '\\' '&' '*' '/' ',' '.'

%left STRING
%left MONTH_IDENT
%left NUMBER

%left '/'
%left '*'
%left '&' '|'
%left '\\'
%left ','
%left '(' ')'
%left '!' 

%%

filter
    : exprs
    {
        yylex.(*Lexer).result = $1
    }

exprs
    : expr
    {
        $$ = []Expression{$1}
    }
    | exprs ',' expr
    {
        $$ = append($1, $3)
    }

expr
    :
    {
        $$ = VoidExpr{}
    }
    | expr '|' expr
    {
        $$ = BoolInfixOpExpr{left: $1, operator: '|', right: $3}
    }
    | expr '&' expr
    {
        $$ = BoolInfixOpExpr{left: $1, operator: '&', right: $3}
    }
    | '(' expr ')'
    {
        $$ = $2
    }
    | '!' expr
    {
        $$ = NotOpExpr{expr: $2}
    }
    | '/' s_string
    {
        $$ = ProjectExpr{isAll: false, section: $2.(StringExpr).String()}
    }
    | s_project_key s_string
    {
        $$ = ProjectExpr{isAll: false, name: $2.(StringExpr).String()}
    }
    | s_project_key s_string '/' s_string
    {
        name := $2.(StringExpr).String()
        section := $4.(StringExpr).String()
        $$ = ProjectExpr{isAll: false, name: name, section: section}
    }
    | s_project_all_key s_string
    {
        $$ = ProjectExpr{isAll: true, name: $2.(StringExpr).String()}
    }
    | s_label_key s_string
    {
        $$ = LabelExpr{name: $2.(StringExpr).String()}
    }
    | s_no_labels
    {
        $$ = LabelExpr{name: ""}
    }
    | s_nodate
    {
        $$ = DateExpr{operation: NO_DUE_DATE}
    }
    | NO TIME
    {
        $$ = DateExpr{operation: NO_TIME}
    }
    | NO PRIORITY
    {
        $$ = NoPriorityExpr{}
    }
    | s_overdue
    {
        $$ = DateExpr{allDay: false, datetime: now(), operation: DUE_BEFORE}
    }
    | VIEW ALL
    {
        $$ = ViewAllExpr{}
    }
    | RECURRING
    {
        $$ = RecurringExpr{}
    }
    | CREATED BEFORE ':' s_datetime
    {
        e := $4.(DateExpr)
        e.operation = CREATED_BEFORE
        $$ = e
    }
    | DUE BEFORE ':' s_datetime
    {
        e := $4.(DateExpr)
        e.operation = DUE_BEFORE
        $$ = e
    }
    | DUE ':' s_datetime
    {
        e := $3.(DateExpr)
        e.operation = DUE_ON
        $$ = e
    }
    | DUE AFTER ':' s_datetime
    {
        e := $4.(DateExpr)
        e.operation = DUE_AFTER
        $$ = e
    }
    | SHARED
    {
        $$ = SharedExpr{}
    }
    | SUBTASK
    {
        $$ = SubtaskExpr{}
    }
    | ASSIGNED
    {
        $$ = AssignedExpr{}
    }
    | SEARCH ':' s_string
    {
        $$ = SearchExpr{keyword: $3.(StringExpr).String()}
    }
    | WEEKDAY
    {
        $$ = WeekdayExpr{day: WeekdayHash[strings.ToLower($1.literal)]}
    }
    | s_person
    | s_datetime
    | s_string


s_special_chars
    : '\\' '&'
    {
        $$ = "&"
    }
    | '\\' '!'
    {
        $$ = "!"
    }
    | '\\' '|'
    {
        $$ = `\|`
    }
    | '\\' '*'
    {
        $$ = `\*`
    }
    | '\\' '('
    {
        $$ = `\(`
    }
    | '\\' ')'
    {
        $$ = `\)`
    }
    | '.'
    {
        $$ =`\.`
    }

s_string
    : STRING
    {
        $$ = NewStringExpr($1.literal)
    }
    | '*'
    {
        $$ = NewStringExpr(".*")
    }
    | s_special_chars
    {
        $$ = NewStringExpr($1.(string))
    }
    | s_string STRING 
    {
        $$ = $1.(StringExpr).Add(NewStringExpr($2.literal))
    }
    | s_string '*' 
    {
        $$ = $1.(StringExpr).Add(NewStringExpr(".*"))
    }
    | s_string s_special_chars
    {
        $$ = $1.(StringExpr).Add(NewStringExpr($2.(string)))
    }


s_person
    : ASSIGNED TO ':' s_string
    {
        $$ = PersonExpr{operation: ASSIGNED_TO, person:$4.(StringExpr).String()}
    }
    | ASSIGNED BY ':' s_string
    {
        $$ = PersonExpr{operation: ASSIGNED_BY, person:$4.(StringExpr).String()}
    }
    | ADDED BY ':' s_string
    {
        $$ = PersonExpr{operation: ADDED_BY, person:$4.(StringExpr).String()}
    }
   
s_project_all_key
    : '#' '#'
    {
        $$ = $1
    }

s_project_key
    : '#'
    {
        $$ = $1
    }

s_label_key
    : AT
    {
        $$ = $1
    }

s_no_labels
    : NO LABELS
    {
        $$ = $1
    }

s_nodate
    : NO DATE
    {
        $$ = $1
    }
    | NO DUE DATE
    {
        $$ = $1
    }

s_overdue
    : OVER DUE
    {
        $$ = $1
    }
    | OVERDUE
    {
        $$ = $1
    }


s_datetime /*  */
    : s_date_year s_time
    {
        date := $1.(time.Time)
        time := $2.(time.Duration)
        $$ = DateExpr{allDay: false, datetime: date.Add(time)}
    }
    | s_date_year
    {
        $$ = DateExpr{allDay: true, datetime: $1.(time.Time)}
    }
    | s_time
    {
        nd := now().Sub(today())
        d := $1.(time.Duration)
        if (d <= nd) {
          d = d + time.Duration(int64(time.Hour) * 24)
        }
        $$ = DateExpr{allDay: false, datetime: today().Add(d)}
    }
    | '-' NUMBER DAYS
    {
        date := today().AddDate(0, 0, -atoi($2.literal))
        $$ = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
    }
    | NEXT NUMBER DAYS
    {
        date := today().AddDate(0, 0, atoi($2.literal))
        $$ = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
    }
    | NUMBER DAYS
    {
        date := today().AddDate(0, 0, atoi($1.literal))
        $$ = DateExpr{allDay: true, datetime: date, operation: DUE_BEFORE}
    }

s_date_year
    : NUMBER '/' NUMBER '/' YEAR_NUMBER
    {
        $$ = time.Date(atoi($5.literal), time.Month(atoi($1.literal)), atoi($3.literal), 0, 0, 0, 0, timezone())
    }
    | MONTH_IDENT NUMBER YEAR_NUMBER
    {
        $$ = time.Date(atoi($3.literal), MonthIdentHash[strings.ToLower($1.literal)], atoi($2.literal), 0, 0, 0, 0, timezone())
    }
    | MONTH_IDENT NUMBER ORDINAL YEAR_NUMBER
    {
        $$ = time.Date(atoi($4.literal), MonthIdentHash[strings.ToLower($1.literal)], atoi($2.literal), 0, 0, 0, 0, timezone())
    }
    | NUMBER MONTH_IDENT YEAR_NUMBER
    {
        $$ = time.Date(atoi($3.literal), MonthIdentHash[strings.ToLower($2.literal)], atoi($1.literal), 0, 0, 0, 0, timezone())
    }
    | s_date
    {
        tod := today()
        date := $1.(time.Time)
        if date.Before(tod) {
            date = date.AddDate(1, 0, 0)
        }
        $$ = date
    }
    | TODAY_IDENT
    {
        $$ = today()
    }
    | TOMORROW_IDENT
    {
        $$ = today().AddDate(0, 0, 1)
    }
    | YESTERDAY_IDENT
    {
        $$ = today().AddDate(0, 0, -1)
    }

s_date
    : MONTH_IDENT NUMBER
    {
        $$ = time.Date(today().Year(), MonthIdentHash[strings.ToLower($1.literal)], atoi($2.literal), 0, 0, 0, 0, timezone())
    }
    | MONTH_IDENT NUMBER ORDINAL 
    {
        $$ = time.Date(today().Year(), MonthIdentHash[strings.ToLower($1.literal)], atoi($2.literal), 0, 0, 0, 0, timezone())
    }
    | NUMBER MONTH_IDENT
    {
        $$ = time.Date(today().Year(), MonthIdentHash[strings.ToLower($2.literal)], atoi($1.literal), 0, 0, 0, 0, timezone())
    }
    | NUMBER '/' NUMBER
    {
        $$ = time.Date(now().Year(), time.Month(atoi($3.literal)), atoi($1.literal), 0, 0, 0, 0, timezone())
    }

s_time
    : NUMBER ':' NUMBER
    {
        $$ = time.Duration(int64(time.Hour) * int64(atoi($1.literal)) + int64(time.Minute) * int64(atoi($3.literal)))
    }
    | NUMBER ':' NUMBER ':' NUMBER
    {
        $$ = time.Duration(int64(time.Hour) * int64(atoi($1.literal)) + int64(time.Minute) * int64(atoi($3.literal)) + int64(time.Second) * int64(atoi($5.literal)))
    }
    | '+' NUMBER HOURS
    {
        $$ = time.Duration(int64(time.Hour) * int64(atoi($2.literal)))
    }
    | NUMBER TWELVE_CLOCK_IDENT
    {
        hour := atoi($1.literal)
        if TwelveClockIdentHash[$2.literal] {
            hour = hour + 12
        }
        $$ = time.Duration(int64(time.Hour) * int64(hour))
    }

%%
