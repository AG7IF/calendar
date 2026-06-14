package plancal

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/ag7if"
	"github.com/derhabicht/planning-tools/pkg/calendar/cap"
)

type QT int

const (
	CalendarQuarter QT = iota
	FiscalQuarter
)

type Quarter struct {
	calendar    calendar.Calendar
	quarterType QT
	year        int
	quarter     calendar.Q
	startWeek   Week
}

func NewQuarter(cal calendar.Calendar, year int, qtr calendar.Q, qt QT) Quarter {
	switch qt {
	case CalendarQuarter:
		weekNo := ag7if.ComputeQuarterStartWeek(qtr)

		return Quarter{
			calendar:    cal,
			quarterType: qt,
			year:        year,
			quarter:     qtr,
			startWeek:   NewFromISOWeek(cal, year, weekNo),
		}
	case FiscalQuarter:
		startDate := cap.ComputeFiscalQuarterStartDate(year, qtr)

		return Quarter{
			calendar:    cal,
			quarterType: qt,
			year:        year,
			quarter:     qtr,
			startWeek:   NewWeek(cal, startDate),
		}
	default:
		panic(errors.Errorf("invalid quarter type value: %d", qt))
	}
}

func (q Quarter) Year() int {
	return q.year
}

func (q Quarter) Quarter() calendar.Q {
	return q.quarter
}

func (q Quarter) StartDay() calendar.Day {
	return q.startWeek.StartDay()
}

func (q Quarter) Short() string {
	return q.quarter.String()
}

func (q Quarter) String() string {
	switch q.quarterType {
	case CalendarQuarter:
		return fmt.Sprintf("CY%d%s", q.year, q.quarter)
	case FiscalQuarter:
		return fmt.Sprintf("FY%d%s", q.year, q.quarter)
	default:
		panic(errors.Errorf("invalid value for quarterType: %d", q.quarterType))
	}
}

func (q Quarter) Full() string {
	switch q.quarterType {
	case CalendarQuarter:
		return fmt.Sprintf("CY%d, %s Quarter", q.year, humanize.Ordinal(int(q.quarter)))
	case FiscalQuarter:
		return fmt.Sprintf("FY%d, %s Quarter", q.year, humanize.Ordinal(int(q.quarter)))
	default:
		panic(errors.Errorf("invalid value for quarterType: %d", q.quarterType))
	}
}

func (q Quarter) FirstMonth() calendar.Month {
	switch q.quarterType {
	case CalendarQuarter:
		mo := ag7if.ComputeQuarterStartMonth(q.quarter)
		return NewMonth(q.calendar, q.year, mo)
	case FiscalQuarter:
		mo := cap.ComputeFiscalQuarterStartMonth(q.quarter)
		if q.quarter == calendar.Q1 {
			return NewMonth(q.calendar, q.year-1, mo)
		}
		return NewMonth(q.calendar, q.year, mo)
	default:
		panic(errors.Errorf("invalid value for quarterType: %d", q.quarterType))
	}
}

func (q Quarter) Next() calendar.Quarter {
	switch q.quarter {
	case calendar.Q1:
		return NewQuarter(q.calendar, q.year, calendar.Q2, q.quarterType)
	case calendar.Q2:
		return NewQuarter(q.calendar, q.year, calendar.Q3, q.quarterType)
	case calendar.Q3:
		return NewQuarter(q.calendar, q.year, calendar.Q4, q.quarterType)
	case calendar.Q4:
		return NewQuarter(q.calendar, q.year+1, calendar.Q1, q.quarterType)
	default:
		panic(errors.Errorf("invalid value for quarter: %d", q.quarter))
	}
}
