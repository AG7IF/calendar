package recurrence

import (
	"time"

	"github.com/fxtlabs/date"
)

type YearlyRecurrence struct {
	month time.Month
	day   int
	until *date.Date
}

func NewYearlyRecurrence(month time.Month, day int, until *date.Date) YearlyRecurrence {
	return YearlyRecurrence{
		month: month,
		day:   day,
		until: until,
	}
}

func (yr YearlyRecurrence) NextOccurrence(now date.Date) *date.Date {

	var next date.Date
	if now.Month() > yr.month || now.Day() >= yr.day {
		next = date.New(now.Year()+1, yr.month, yr.day)
	} else {
		next = date.New(now.Year(), yr.month, yr.day)
	}

	if yr.until != nil && next.After(*yr.until) {
		return nil
	}

	return &next
}
