package calc

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

func ComputeFirstGivenWeekdayOfMonth(d date.Date, wd time.Weekday) date.Date {
	fom := date.New(d.Year(), d.Month(), 1)
	dist := int(wd) - int(fom.Weekday())
	fom = fom.Add(dist)

	if d.Month() != fom.Month() {
		fom = fom.Add(7)
	}

	return fom
}

func ComputeNearestMonday(d date.Date) date.Date {
	var dd int
	wd := d.Weekday()
	switch {
	case time.Tuesday <= wd && wd <= time.Thursday:
		dd = -1 * (int(wd) - 1)
	case wd >= time.Friday:
		dd = 8 - int(wd)
	case wd == time.Sunday:
		dd = 1
	default:
		dd = 0
	}

	return d.Add(dd)
}

func ComputeNearestThursday(d date.Date) date.Date {
	monday := ComputeNearestMonday(d)
	return monday.Add(3)
}

func ComputeLastDayOfMonth(d date.Date) int {
	switch d.Month() {
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	case time.February:
		if date.New(d.Year(), time.December, 31).YearDay() == 366 {
			return 29
		}

		return 28
	default:
		panic(errors.Errorf("not a valid time.Month value: %d", d.Month()))
	}
}

func TimeToLocalDate(t time.Time, timezone *time.Location) date.Date {
	local := t.In(timezone)
	return date.New(local.Year(), local.Month(), local.Day())
}

func DateToLocalTime(d date.Date, timezone *time.Location) time.Time {
	return d.In(timezone)
}

func DayOfWeekGTE(l time.Weekday, r time.Weekday) bool {
	li := int(l)
	if li == 0 {
		li = 7
	}

	ri := int(r)
	if ri == 0 {
		ri = 7
	}

	return li >= ri
}
