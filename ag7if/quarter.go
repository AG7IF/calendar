package ag7if

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/ag7if/calendar/calendar"
)

func ComputeQuarter(d date.Date) (int, calendar.Q) {
	year, week := d.ISOWeek()

	switch {
	case (1 <= week) && (week <= 13):
		return year, calendar.Q1
	case (14 <= week) && (week <= 26):
		return year, calendar.Q2
	case (27 <= week) && (week <= 39):
		return year, calendar.Q3
	case (40 <= week) && (week <= 53):
		return year, calendar.Q4
	default:
		panic(errors.Errorf("invalid week number: %d", week))
	}
}

func ComputeQuarterStartWeek(qtr calendar.Q) int {
	switch qtr {
	case calendar.Q1:
		return 1
	case calendar.Q2:
		return 14
	case calendar.Q3:
		return 27
	case calendar.Q4:
		return 40
	default:
		panic(errors.Errorf("invalid quarter value: %d", qtr))
	}
}

func ComputeQuarterStartMonth(qtr calendar.Q) time.Month {
	switch qtr {
	case calendar.Q1:
		return time.January
	case calendar.Q2:
		return time.April
	case calendar.Q3:
		return time.July
	case calendar.Q4:
		return time.October
	default:
		panic(errors.Errorf("invalid quarter value: %d", qtr))
	}
}

func ComputeQuarterStartDate(year int, qtr calendar.Q) date.Date {
	start := ComputeWeek1StartDate(year)

	switch qtr {
	case calendar.Q1:
		return start
	case calendar.Q2:
		return start.Add(13 * 7)
	case calendar.Q3:
		return start.Add(26 * 7)
	case calendar.Q4:
		return start.Add(39 * 7)
	default:
		panic(errors.Errorf("invalid quarter value: %d", qtr))
	}
}
