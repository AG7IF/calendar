package calendar

import (
	"time"

	cards "github.com/ag7if/playing-cards"
	"github.com/fxtlabs/date"
)

type Calendar interface {
	FiscalYear() int
	JulianPeriod() int
	NaturalYearDecade() int
	NaturalYearLustrum() int
	NaturalYearTriad() int
	NaturalYear() int
	LunarCalibrationDate() date.Date
	SolsticeTable() SolsticeTable
	HolidayCalendar() HolidayCalendar
	FirstTrimester() Trimester
	FirstCalendarQuarter() Quarter
	FirstFiscalQuarter() Quarter
	FirstMonth() Month
	FirstSprint() Sprint
	FirstWeek() Week
	FetchWeek(year, week int) (Week, error)
}

type Period interface {
	StartDay() Day
	Short() string
	String() string
	Full() string
}

type Trimester interface {
	Period
	Year() int
	Trimester() T
	FirstMonth() Month
	Next() Trimester
}

type Quarter interface {
	Period
	Year() int
	Quarter() Q
	FirstMonth() Month
	Next() Quarter
}

type Month interface {
	Period
	Year() int
	Month() time.Month
	FirstWeek() Week
	Prev() Month
	Next() Month
}

type Sprint interface {
	Period
	FirstWeek() Week
	Next() Sprint
}

type Week interface {
	Period
	Trimester() Trimester
	FiscalQuarter() Quarter
	FyWeek() (int, int)
	CalendarQuarter() Quarter
	Sprint() Sprint
	ISOWeek() (int, int, cards.Card)
	Next() Week
}

type Day interface {
	Date() date.Date
	ISODate() string
	IsHoliday() (bool, bool, Holiday)
	IsSolstice() Solstice
	OrdinalDay() int
	Weekday() time.Weekday
	WeekdayOccurrenceInMonth() int
	MJD() int
	Sunrise() time.Time
	Sunset() time.Time
	Next() Day
}

type SolsticeTable interface {
	IsSolstice(date date.Date) Solstice
	FirstWinterSolstice() time.Time
	VernalEquinox() time.Time
	SummerSolstice() time.Time
	AutumnalEquinox() time.Time
	SecondWinterSolstice() time.Time
}

type HolidayCalendar interface {
	IsHoliday(date.Date) (bool, bool, Holiday)
	Holidays() []Holiday
}

type Holiday interface {
	Occurs(year int) (date.Date, date.Date)
	String() string
	FullName() string
}
