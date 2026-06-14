package cap

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/pkg/calendar"
)

func ComputeFiscalYear(d date.Date) int {
	if d.Month() == time.October || d.Month() == time.November || d.Month() == time.December {
		return d.Year() + 1
	}

	return d.Year()
}

func ComputeFiscalYearStartDate(fy int) date.Date {
	d := date.New(fy-1, time.October, 1)

	d = calendar.ComputeNearestMonday(d)

	return d
}

func ComputeFiscalTrimester(d date.Date) calendar.T {
	switch d.Month() {
	case time.October, time.November, time.December, time.January:
		return calendar.T1
	case time.February, time.March, time.April, time.May:
		return calendar.T2
	case time.June, time.July, time.August, time.September:
		return calendar.T3
	default:
		panic(errors.Errorf("invalid month value: %d", d.Month()))
	}
}

func ComputeFiscalTrimesterStartDate(fy int, trimester calendar.T) date.Date {
	switch trimester {
	case calendar.T1:
		return date.New(fy-1, time.October, 1)
	case calendar.T2:
		return date.New(fy, time.February, 1)
	case calendar.T3:
		return date.New(fy, time.June, 1)
	default:
		panic(errors.Errorf("invalid trimester value: %d", trimester))
	}
}

func ComputeFiscalQuarter(d date.Date) calendar.Q {
	switch d.Month() {
	case time.October, time.November, time.December:
		return calendar.Q1
	case time.January, time.February, time.March:
		return calendar.Q2
	case time.April, time.May, time.June:
		return calendar.Q3
	case time.July, time.August, time.September:
		return calendar.Q4
	default:
		panic(errors.Errorf("invalid month value: %d", d.Month()))
	}
}

func ComputeFiscalQuarterStartMonth(q calendar.Q) time.Month {
	switch q {
	case calendar.Q1:
		return time.October
	case calendar.Q2:
		return time.January
	case calendar.Q3:
		return time.April
	case calendar.Q4:
		return time.July
	default:
		panic(errors.Errorf("invalid quarter value: %d", q))
	}
}

func ComputeFiscalQuarterStartDate(fy int, q calendar.Q) date.Date {
	month := ComputeFiscalQuarterStartMonth(q)

	if q == calendar.Q1 {
		return date.New(fy-1, month, 1)
	}

	return date.New(fy, month, 1)
}

func ComputeFiscalWeek(d date.Date) (int, int) {
	fy := ComputeFiscalYear(d)
	fyStart := ComputeFiscalYearStartDate(fy)

	return fy, (d.Sub(fyStart) / 7) + 1
}
