package plancal

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/cap"
)

type Trimester struct {
	calendar   calendar.Calendar
	trimester  calendar.T
	fiscalYear int
	startDate  date.Date
}

func NewTrimester(cal calendar.Calendar, fiscalYear int, trimester calendar.T) Trimester {
	startDate := cap.ComputeFiscalTrimesterStartDate(fiscalYear, trimester)

	return Trimester{
		calendar:   cal,
		trimester:  trimester,
		fiscalYear: fiscalYear,
		startDate:  startDate,
	}
}

func (t Trimester) Year() int {
	return t.fiscalYear
}

func (t Trimester) Trimester() calendar.T {
	return t.trimester
}

func (t Trimester) StartDay() calendar.Day {
	return NewDay(t.calendar, t.startDate)
}

func (t Trimester) Short() string {
	return t.trimester.String()
}

func (t Trimester) String() string {
	return fmt.Sprintf("FY%d%s", t.fiscalYear, t.trimester)
}

func (t Trimester) Full() string {
	return fmt.Sprintf("FY%d, %s Trimester", t.fiscalYear, humanize.Ordinal(int(t.trimester)))
}

func (t Trimester) FirstMonth() calendar.Month {
	switch t.trimester {
	case calendar.T1:
		return NewMonth(t.calendar, t.fiscalYear-1, time.October)
	case calendar.T2:
		return NewMonth(t.calendar, t.fiscalYear, time.February)
	case calendar.T3:
		return NewMonth(t.calendar, t.fiscalYear, time.June)
	default:
		panic(errors.Errorf("invalid trimester value: %d", t.trimester))
	}
}

func (t Trimester) Next() calendar.Trimester {
	switch t.trimester {
	case calendar.T1:
		return NewTrimester(t.calendar, t.fiscalYear, calendar.T2)
	case calendar.T2:
		return NewTrimester(t.calendar, t.fiscalYear, calendar.T3)
	case calendar.T3:
		return NewTrimester(t.calendar, t.fiscalYear+1, calendar.T1)
	default:
		panic(errors.Errorf("invalid trimester value: %d", t.trimester))
	}
}
