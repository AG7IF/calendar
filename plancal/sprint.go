package plancal

import (
	"fmt"

	"github.com/fxtlabs/date"

	"github.com/derhabicht/planning-tools/pkg/calendar"
	"github.com/derhabicht/planning-tools/pkg/calendar/ag7if"
)

type Sprint struct {
	calendar  calendar.Calendar
	year      int
	sprint    ag7if.S
	firstWeek Week
}

func NewSprint(cal calendar.Calendar, yr int, sprint ag7if.S) Sprint {
	wks := ag7if.ComputeSprintWeekNumbers(sprint)
	firstWeek := NewFromISOWeek(cal, yr, wks[0])

	return Sprint{
		calendar:  cal,
		year:      yr,
		sprint:    sprint,
		firstWeek: firstWeek,
	}
}

func NewSprintFromDate(cal calendar.Calendar, d date.Date) Sprint {
	year, sprint := ag7if.ComputeSprint(d)

	return NewSprint(cal, year, sprint)
}

func (s Sprint) StartDay() calendar.Day {
	return s.firstWeek.StartDay()
}

func (s Sprint) Short() string {
	return s.sprint.String()
}

func (s Sprint) String() string {
	return fmt.Sprintf("CY%d%s", s.year, s.sprint)
}

func (s Sprint) Full() string {
	return fmt.Sprintf("CY%d, Sprint %s ", s.year, s.sprint)
}

func (s Sprint) FirstWeek() calendar.Week {
	return s.firstWeek
}

func (s Sprint) Next() calendar.Sprint {
	if s.sprint == ag7if.SP4 {
		return NewSprint(s.calendar, s.year+1, ag7if.S01)
	}

	return NewSprint(s.calendar, s.year, s.sprint+1)
}
