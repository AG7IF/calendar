package plancal

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/snabb/isoweek"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/moonphase"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/holidays"
	"github.com/derhabicht/planning-tools/pkg/calendar/natyr"
	"github.com/derhabicht/planning-tools/pkg/calendar/solstice_table"
)

const julianPeriodOffest = 4713

type Calendar struct {
	fiscalYear           int
	naturalYear          natyr.NaturalYear
	lunarCalibrationDate date.Date
	solsticeTable        calendar.SolsticeTable
	holidayCalendar      calendar.HolidayCalendar
}

func NewCalendar(fiscalYear int, birthday date.Date) *Calendar {
	st := solstice_table.NewSolsticeTable(fiscalYear)
	hc := holidays.NewHolidayCalendar()

	return &Calendar{
		fiscalYear:           fiscalYear,
		naturalYear:          natyr.NewNaturalYear(birthday),
		lunarCalibrationDate: computeLunarCalibrationDate(fiscalYear),
		solsticeTable:        st,
		holidayCalendar:      hc,
	}
}

func computeLunarCalibrationDate(fiscalYear int) date.Date {
	// April is the middle month of the 15 months rendered in this planning_calendar.
	// This makes the full moon closest to April 15 an ideal candidate for the lunar calibration date.
	refDate := date.New(fiscalYear, time.April, 15)

	ref := float64(fiscalYear) + (float64(refDate.YearDay()) / float64(date.New(fiscalYear, time.December, 31).Year()))
	jde := moonphase.New(ref)
	lcd := julian.JDToTime(jde)

	return date.New(lcd.Year(), lcd.Month(), lcd.Day())
}

func (c *Calendar) FiscalYear() int {
	return c.fiscalYear
}

func (c *Calendar) JulianPeriod() int {
	return c.fiscalYear + julianPeriodOffest
}

func (c *Calendar) NaturalYearDecade() int {
	return c.naturalYear.Decade(date.New(c.fiscalYear-1, time.October, 1))
}

func (c *Calendar) NaturalYearLustrum() int {
	return c.naturalYear.Lustrum(date.New(c.fiscalYear-1, time.October, 1))
}

func (c *Calendar) NaturalYearTriad() int {
	return c.naturalYear.Triad(date.New(c.fiscalYear-1, time.October, 1))
}

func (c *Calendar) NaturalYear() int {
	ny, _ := c.naturalYear.ToNYDate(date.New(c.fiscalYear-1, time.October, 1))
	return ny
}

func (c *Calendar) LunarCalibrationDate() date.Date {
	return c.lunarCalibrationDate
}

func (c *Calendar) SolsticeTable() calendar.SolsticeTable {
	return c.solsticeTable
}

func (c *Calendar) HolidayCalendar() calendar.HolidayCalendar {
	return c.holidayCalendar
}

func (c *Calendar) FirstTrimester() calendar.Trimester {
	return NewTrimester(c, c.fiscalYear, calendar.T1)
}

func (c *Calendar) FirstCalendarQuarter() calendar.Quarter {
	return NewQuarter(c, c.fiscalYear-1, calendar.Q4, CalendarQuarter)
}

func (c *Calendar) FirstFiscalQuarter() calendar.Quarter {
	return NewQuarter(c, c.fiscalYear, calendar.Q1, FiscalQuarter)
}

func (c *Calendar) FirstMonth() calendar.Month {
	return NewMonth(c, c.fiscalYear-1, time.October)
}

func (c *Calendar) FirstSprint() calendar.Sprint {
	panic("implement me")
}

func (c *Calendar) FirstWeek() calendar.Week {
	return c.FirstMonth().FirstWeek()
}

func (c *Calendar) FetchWeek(year, week int) (calendar.Week, error) {
	y, m, d := isoweek.StartDate(year, week)

	wk := NewWeek(c, date.New(y, m, d))

	return wk, nil
}
