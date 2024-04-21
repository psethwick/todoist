package filter

import (
	"regexp"
	"strings"
	"time"

	todoist "github.com/psethwick/todoist/lib"
)

func Eval(e Expression, item todoist.AbstractItem, store *todoist.Store) (result bool, err error) {
	result = false
	switch e.(type) {
	case BoolInfixOpExpr:
		e := e.(BoolInfixOpExpr)
		lr, err := Eval(e.left, item, store)
		if err != nil {
			return false, nil
		}
		rr, err := Eval(e.right, item, store)
		if err != nil {
			return false, nil
		}
		switch e.operator {
		case '&':
			return lr && rr, nil
		case '|':
			return lr || rr, nil
		}
	case AssignedExpr:
		return item.GetResponsibleUID() != "", nil
	case NoPriorityExpr:
		return item.HasPriority(), nil
	case RecurringExpr:
		return item.IsRecurring(), nil
	case SubtaskExpr:
		return item.HasParent(), nil
	case WeekdayExpr:
		due := item.DateTime()
		if (due == time.Time{}) {
			return false, nil
		}
		return due.Weekday() == e.(WeekdayExpr).day, nil
	case SearchExpr:
		keywords := e.(SearchExpr).keyword
		return item.SearchMatches(keywords), nil
	case PersonExpr:
		e := e.(PersonExpr)
		return EvalPerson(e, item, store.User, store.Collaborators), err
	case ProjectExpr:
		e := e.(ProjectExpr)
		return EvalProject(e, item.GetProjectID(), store.Projects), err
	case LabelExpr:
		e := e.(LabelExpr)
		return EvalLabel(e, item.GetLabelNames()), err
	case PriorityExpr:
		switch item.(type) {
		case *todoist.Item:
			item := item.(*todoist.Item)
			return e.(PriorityExpr).priority == todoist.PriorityMapping[item.Priority], err
		default:
			// CompletedItem
			return false, nil
		}
	case DateExpr:
		e := e.(DateExpr)
		return EvalDate(e, item), err
	case NotOpExpr:
		e := e.(NotOpExpr)
		r, err := Eval(e.expr, item, store)
		if err != nil {
			return false, nil
		}
		return !r, nil
	case ViewAllExpr:
		return true, nil
	default:
		return true, err
	}
	return
}

func doesPersonMatch(e PersonExpr, user todoist.User, collaborators todoist.Collaborators, id string) bool {
	switch strings.ToLower(e.person) {
	case "me":
		return id == user.ID
	case "others":
		return id != "" && id != user.ID
	default:
		for _, c := range collaborators {
			matchEmail, _ := regexp.MatchString(e.person, c.Email)
			if matchEmail {
				return id == c.ID
			}
			matchName, _ := regexp.MatchString(e.person, c.FullName)
			if matchName {
				return id == c.ID
			}
		}
		return false
	}
}

func EvalPerson(e PersonExpr, item todoist.AbstractItem, user todoist.User, collaborators todoist.Collaborators) (result bool) {
	switch e.operation {
	case ASSIGNED_BY:
		return doesPersonMatch(e, user, collaborators, item.GetAssignedByUID())
	case ASSIGNED_TO:
		id := item.GetResponsibleUID()
		if id == "" {
			return false
		}
		return doesPersonMatch(e, user, collaborators, id)
	case ADDED_BY:
		return doesPersonMatch(e, user, collaborators, item.GetUserID())
	}
	return true
}

func EvalDate(e DateExpr, item todoist.AbstractItem) (result bool) {
	itemDate := item.DateTime()
	if e.operation == NO_TIME {
		return item.HasTime()
	}

	if (itemDate == time.Time{}) {
		return e.operation == NO_DUE_DATE
	}
	allDay := e.allDay
	theDate := e.datetime
	switch e.operation {
	case CREATED_BEFORE:
		if _, ok := item.(todoist.CompletedItem); ok {
			return false
		}
		return item.CreatedTime().Before(theDate)
	case CREATED_AFTER:
		if _, ok := item.(todoist.CompletedItem); ok {
			return false
		}
		return item.CreatedTime().After(theDate)
	case DUE_ON:
		var startDate, endDate time.Time
		if allDay {
			startDate = theDate
			endDate = theDate.AddDate(0, 0, 1)
			if itemDate.Equal(startDate) || (itemDate.After(startDate) && itemDate.Before(endDate)) {
				return true
			}
		}
		return false
	case DUE_BEFORE:
		return itemDate.Before(theDate)
	case DUE_AFTER:
		endDateTime := theDate
		if allDay {
			endDateTime = theDate.AddDate(0, 0, 1).Add(-time.Duration(time.Microsecond))
		}
		return itemDate.After(endDateTime)
	default:
		return false
	}
}

func EvalProject(e ProjectExpr, projectID string, projects todoist.Projects) bool {
	for _, id := range projects.GetIDsByName(e.name, e.isAll) {
		if id == projectID {
			return true
		}
	}
	return false
}

func EvalLabel(e LabelExpr, labelNames []string) bool {
	if e.name == "" {
		return len(labelNames) == 0
	}

	for _, name := range labelNames {
		if name == e.name {
			return true
		}
	}

	return false
}
