package solstice_table

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/solstice"

	"github.com/derhabicht/planning-tools/pkg/calendar"
)

type SolsticeTable struct {
	firstWinterSolstice  time.Time
	vernalEquinox        time.Time
	summerSolstice       time.Time
	autumnalEquinox      time.Time
	secondWinterSolstice time.Time
}

func NewSolsticeTable(fy int) *SolsticeTable {
	winter1 := julian.JDToTime(solstice.December(fy - 1))
	vernal := julian.JDToTime(solstice.March(fy))
	summmer := julian.JDToTime(solstice.June(fy))
	autumnal := julian.JDToTime(solstice.September(fy))
	winter2 := julian.JDToTime(solstice.December(fy))

	return &SolsticeTable{
		firstWinterSolstice:  winter1,
		vernalEquinox:        vernal,
		summerSolstice:       summmer,
		autumnalEquinox:      autumnal,
		secondWinterSolstice: winter2,
	}
}

func (s *SolsticeTable) IsSolstice(date date.Date) calendar.Solstice {
	switch date.Month() {
	case time.March:
		if date == calendar.TimeToLocalDate(s.vernalEquinox) {
			return calendar.VernalEquinox
		}
	case time.June:
		if date == calendar.TimeToLocalDate(s.summerSolstice) {
			return calendar.SummerSolstice
		}
	case time.September:
		if date == calendar.TimeToLocalDate(s.autumnalEquinox) {
			return calendar.AutumnalEquinox
		}
	case time.December:
		if date == calendar.TimeToLocalDate(s.firstWinterSolstice) || date == calendar.TimeToLocalDate(s.secondWinterSolstice) {
			return calendar.WinterSolstice
		}
	default:
		break
	}
	return calendar.NoSolstice
}

func (s *SolsticeTable) FirstWinterSolstice() time.Time {
	return s.firstWinterSolstice
}

func (s *SolsticeTable) VernalEquinox() time.Time {
	return s.vernalEquinox
}

func (s *SolsticeTable) SummerSolstice() time.Time {
	return s.summerSolstice
}

func (s *SolsticeTable) AutumnalEquinox() time.Time {
	return s.autumnalEquinox
}

func (s *SolsticeTable) SecondWinterSolstice() time.Time {
	return s.secondWinterSolstice
}
