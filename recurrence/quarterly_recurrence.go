package recurrence

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/snabb/isoweek"

	"github.com/ag7if/calendar/ag7if"
	"github.com/ag7if/calendar/calendar"
)

type QuarterlyRecurrence struct {
	week    int
	weekday time.Weekday
	until   *date.Date
}

func NewQuarterlyRecurrence(week int, weekday time.Weekday, until *date.Date) QuarterlyRecurrence {
	return QuarterlyRecurrence{
		week:    week,
		weekday: weekday,
		until:   until,
	}
}

func DOWGTE(l time.Weekday, r time.Weekday) bool {
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

func (qr QuarterlyRecurrence) NextOccurrence(now date.Date) *date.Date {
	_, qtr := ag7if.ComputeQuarter(now)

	var nextWeek int
	switch qtr {
	case calendar.Q1:
		nextWeek = qr.week
	case calendar.Q2:
		nextWeek = qr.week + 13
	case calendar.Q3:
		nextWeek = qr.week + 26
	case calendar.Q4:
		nextWeek = qr.week + 39
	}

	yr, nowWeek := isoweek.FromDate(now.Year(), now.Month(), now.Day())
	if nowWeek == nextWeek && DOWGTE(now.Weekday(), qr.weekday) {
		nextWeek += 13
		if nowWeek > 52 {
			nextWeek = qr.week
			yr += 1
		}
	}

	nextYr, nextMo, nextDy := isoweek.StartDate(yr, nextWeek)

	if qr.weekday == time.Sunday {
		nextDy += 6
	} else {
		nextDy += int(qr.weekday) - 1
	}

	next := date.New(nextYr, nextMo, nextDy)

	if qr.until != nil && next.After(*qr.until) {
		return nil
	}

	return &next
}
