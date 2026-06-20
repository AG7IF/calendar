package recurrence

import (
	"time"

	"github.com/fxtlabs/date"

	"github.com/ag7if/calendar/calc"
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
	day := mr.day
	if day > calc.ComputeLastDayOfMonth(now) {
		day = calc.ComputeLastDayOfMonth(now)
	}

	year := now.Year()
	month := now.Month()

	if now.Day() >= mr.day {
		if month == time.December {
			year += 1
			month = time.January
		} else {
			month += 1
		}

		if day != mr.day {
			day = calc.ComputeLastDayOfMonth(date.New(year, month, 1))
		}
	}

	next := date.New(year, month, day)
	if mr.until != nil && next.After(*mr.until) {
		return nil
	}

	return &next
}

func (mr MonthlyRecurrence) determineNextWeekdayOccurrence(now date.Date) *date.Date {
	next := now

	for {
		yr := next.Year()
		mo := next.Month()
		next = calc.ComputeFirstGivenWeekdayOfMonth(next, mr.weekday)

		next = next.Add(7 * (mr.week - 1))

		if next.Month() != mo {
			continue
		}

		if now.Before(next) {
			break
		}

		if mo == time.December {
			yr++
			mo = time.January
		} else {
			mo++
		}

		next = date.New(yr, mo, 1)
	}

	if mr.until != nil && next.After(*mr.until) {
		return nil
	}

	return &next
}

func (mr MonthlyRecurrence) NextOccurrence(now date.Date) *date.Date {
	if mr.onDay {
		return mr.determineNextOnDayOccurrence(now)
	}

	return mr.determineNextWeekdayOccurrence(now)
}
