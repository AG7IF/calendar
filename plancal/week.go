package plancal

import (
	"fmt"
	"time"

	cards "github.com/ag7if/playing-cards"
	"github.com/fxtlabs/date"
	"github.com/pkg/errors"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/ag7if"
	"github.com/derhabicht/planning-tools/pkg/calendar/cap"
)

type Week struct {
	calendar calendar.Calendar
	monday   date.Date
	thursday date.Date
}

func NewWeek(cal calendar.Calendar, d date.Date) Week {
	var monday date.Date

	if d.Weekday() == time.Sunday {
		monday = d.Add(-6)
	} else {
		monday = d.Add(-1 * (int(d.Weekday()) - int(time.Monday)))
	}

	thursday := monday.Add(3)

	return Week{
		calendar: cal,
		monday:   monday,
		thursday: thursday,
	}
}

func NewFromISOWeek(cal calendar.Calendar, year, isoweek int) Week {
	// 4 Jan is guaranteed to be a part of week 1, so we use 4 Jan of year as our reference date.
	refDate := date.New(year, time.January, 4)

	// Advance the ref date by isoweek-1 weeks to give us a date within the week that we want.
	d := refDate.Add(7 * (isoweek - 1))

	return NewWeek(cal, d)
}

func (w Week) StartDay() calendar.Day {
	return NewDay(w.calendar, w.monday)
}

func (w Week) Short() string {
	_, wk := w.monday.ISOWeek()
	return fmt.Sprintf("W%02d", wk)
}

func (w Week) String() string {
	yr, wk := w.monday.ISOWeek()
	return fmt.Sprintf("CY%04dW%02d", yr, wk)
}

func (w Week) Full() string {
	cy, cywk := w.monday.ISOWeek()
	fy, fyweek := cap.ComputeFiscalWeek(w.monday)

	return fmt.Sprintf("CY%04dW%02d/FY%04dW%02d", cy, cywk, fy, fyweek)
}

func (w Week) Trimester() calendar.Trimester {
	yr := cap.ComputeFiscalYear(w.monday)
	tri := cap.ComputeFiscalTrimester(w.monday)

	return NewTrimester(w.calendar, yr, tri)
}

func (w Week) FiscalQuarter() calendar.Quarter {
	yr := cap.ComputeFiscalYear(w.monday)
	qtr := cap.ComputeFiscalQuarter(w.monday)

	return NewQuarter(w.calendar, yr, qtr, FiscalQuarter)
}

func (w Week) FyWeek() (int, int) {
	return cap.ComputeFiscalWeek(w.thursday)
}

func (w Week) CalendarQuarter() calendar.Quarter {
	yr, qtr := ag7if.ComputeQuarter(w.monday)
	return NewQuarter(w.calendar, yr, qtr, CalendarQuarter)
}

func (w Week) Sprint() calendar.Sprint {
	yr, s := ag7if.ComputeSprint(w.monday)
	return NewSprint(w.calendar, yr, s)
}

func (w Week) ISOWeek() (int, int, cards.Card) {
	yr, wk := w.monday.ISOWeek()
	card, err := ag7if.ComputeWeekPlayingCard(wk)
	if err != nil {
		panic(errors.WithMessage(err, "unexpected error encountered while looking up a week playing card"))
	}

	return yr, wk, card
}

func (w Week) Next() calendar.Week {
	nextMonday := w.monday.Add(7)

	return NewWeek(w.calendar, nextMonday)
}
