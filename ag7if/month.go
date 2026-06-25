package ag7if

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

// ComputeMonth identifies in which *fiscal* month a given date lands, where
// fiscal months are defined by a 4-4-5 cadence on the ISO 8601 week year.
func ComputeMonth(d date.Date) (int, time.Month) {
	year, week := d.ISOWeek()

	switch {
	case (1 <= week) && (week <= 4):
		return year, time.January
	case (5 <= week) && (week <= 8):
		return year, time.February
	case (9 <= week) && (week <= 13):
		return year, time.March
	case (14 <= week) && (week <= 17):
		return year, time.April
	case (18 <= week) && (week <= 21):
		return year, time.May
	case (22 <= week) && (week <= 26):
		return year, time.June
	case (27 <= week) && (week <= 30):
		return year, time.July
	case (31 <= week) && (week <= 34):
		return year, time.August
	case (35 <= week) && (week <= 39):
		return year, time.September
	case (40 <= week) && (week <= 43):
		return year, time.October
	case (44 <= week) && (week <= 47):
		return year, time.November
	case (48 <= week) && (week <= 53):
		return year, time.December
	default:
		panic(errors.Errorf("invalid week number: %d", week))
	}
}

// ComputeMonthStartWeek computes the starting ISO 8601 week number of a *fiscal*
// month as defined by a 4-4-5 cadence of the ISO 8601 week year.
func ComputeMonthStartWeek(m time.Month) int {
	switch m {
	case time.January:
		return 1
	case time.February:
		return 5
	case time.March:
		return 9
	case time.April:
		return 14
	case time.May:
		return 18
	case time.June:
		return 22
	case time.July:
		return 27
	case time.August:
		return 31
	case time.September:
		return 35
	case time.October:
		return 40
	case time.November:
		return 44
	case time.December:
		return 48
	default:
		panic(errors.Errorf("invalid month value: %d", m))
	}
}

// ComputeMonthStartDate computes the start of a *fiscal* month as defined by
// a 4-4-5 cadence of the ISO 8601 week year.
func ComputeMonthStartDate(year int, month time.Month) date.Date {
	start := ComputeWeek1StartDate(year)

	startWeek := ComputeMonthStartWeek(month)

	return start.Add((startWeek - 1) * 7)
}
