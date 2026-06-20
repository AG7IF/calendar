package recurrence

import (
	"time"

	"github.com/fxtlabs/date"
)

type MonthlyRecurrence struct {
	day     int
	week    int
	weekday time.Weekday
	until   *date.Date
	onDay   bool
}

func NewOnDayMonthlyRecurrence(day int, until *date.Date) MonthlyRecurrence {
	return MonthlyRecurrence{
		day:   day,
		until: until,
		onDay: true,
	}
}

func NewWeekdayMonthlyRecurrence(week int, weekday time.Weekday, until *date.Date) MonthlyRecurrence {
	return MonthlyRecurrence{
		week:    week,
		weekday: weekday,
		until:   until,
		onDay:   false,
	}
}

func (mr MonthlyRecurrence) determineNextOnDayOccurrence(now date.Date) *date.Date {
	year := now.Year()
	month := now.Month()

	if now.Day() >= mr.day {
		if month == time.December {
			year += 1
			month = time.January
		} else {
			month += 1
		}
	}

	next := date.New(year, month, mr.day)
	return &next
}

func (mr MonthlyRecurrence) determineNextWeekdayOccurrence(now date.Date) *date.Date {
	panic("implement me!")
}

func (mr MonthlyRecurrence) NextOccurrence(now date.Date) *date.Date {
	if mr.onDay {
		return mr.determineNextOnDayOccurrence(now)
	}

	return mr.determineNextWeekdayOccurrence(now)
}
