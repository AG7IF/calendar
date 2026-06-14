package plancal

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/nathan-osman/go-sunrise"
	"github.com/rickar/cal/v2"

	"github.com/derhabicht/planning-tools/internal/config"
	"github.com/derhabicht/planning-tools/pkg/calendar"
)

type Day struct {
	calendar calendar.Calendar
	date     date.Date
	sunrise  time.Time
	sunset   time.Time
}

func NewDay(cal calendar.Calendar, d date.Date) Day {
	lat := config.GetFloat64(config.HomeLocationLat)
	long := config.GetFloat64(config.HomeLocationLong)
	loc := calendar.GetLocation()
	sr, ss := sunrise.SunriseSunset(lat, long, d.Year(), d.Month(), d.Day())

	return Day{
		calendar: cal,
		date:     d,
		sunrise:  sr.In(loc),
		sunset:   ss.In(loc),
	}
}

func (d Day) Date() date.Date {
	return d.date
}

func (d Day) ISODate() string {
	return d.date.FormatISO(4)
}

func (d Day) IsHoliday() (bool, bool, calendar.Holiday) {
	return d.calendar.HolidayCalendar().IsHoliday(d.date)
}

func (d Day) IsSolstice() calendar.Solstice {
	return d.calendar.SolsticeTable().IsSolstice(d.date)
}

func (d Day) OrdinalDay() int {
	return d.date.YearDay()
}

func (d Day) Weekday() time.Weekday {
	return d.date.Weekday()
}

func (d Day) WeekdayOccurrenceInMonth() int {
	// You can find the occurrence of the weekday by performing an integer division by 7 and adding 1.
	// If the current day is the 7th of the month, it is the first occurrence of that weekday, but the above method
	// mistakenly identifies it as the 2nd. So, we zero-index the day first (subtract 1).
	return ((d.date.Day() - 1) / 7) + 1
}

func (d Day) MJD() int {
	return int(cal.ModifiedJulianDate(calendar.DateToLocalTime(d.date)))
}

func (d Day) Sunrise() time.Time {
	return d.sunrise
}

func (d Day) Sunset() time.Time {
	return d.sunset
}

func (d Day) Next() calendar.Day {
	next := d.date.Add(1)

	return NewDay(d.calendar, next)
}
