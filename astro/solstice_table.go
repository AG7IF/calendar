package astro

import (
	"time"

	"github.com/fxtlabs/date"
	"github.com/soniakeys/meeus/v3/julian"
	"github.com/soniakeys/meeus/v3/solstice"

	"github.com/ag7if/calendar/calc"
	"github.com/ag7if/calendar/location"
)

type SolsticeTable struct {
	firstWinterSolstice  time.Time
	vernalEquinox        time.Time
	summerSolstice       time.Time
	autumnalEquinox      time.Time
	secondWinterSolstice time.Time
	timezone             location.TZ
}

func NewSolsticeTable(fy int, timezone location.TZ) SolsticeTable {
	winter1 := julian.JDToTime(solstice.December(fy - 1))
	vernal := julian.JDToTime(solstice.March(fy))
	summmer := julian.JDToTime(solstice.June(fy))
	autumnal := julian.JDToTime(solstice.September(fy))
	winter2 := julian.JDToTime(solstice.December(fy))

	return SolsticeTable{
		firstWinterSolstice:  winter1,
		vernalEquinox:        vernal,
		summerSolstice:       summmer,
		autumnalEquinox:      autumnal,
		secondWinterSolstice: winter2,
		timezone:             timezone,
	}
}

func (s SolsticeTable) IsSolstice(date date.Date) Solstice {
	switch date.Month() {
	case time.March:
		if date == calc.TimeToLocalDate(s.vernalEquinox, s.timezone) {
			return VernalEquinox
		}
	case time.June:
		if date == calc.TimeToLocalDate(s.summerSolstice, s.timezone) {
			return SummerSolstice
		}
	case time.September:
		if date == calc.TimeToLocalDate(s.autumnalEquinox, s.timezone) {
			return AutumnalEquinox
		}
	case time.December:
		if date == calc.TimeToLocalDate(s.firstWinterSolstice, s.timezone) || date == calc.TimeToLocalDate(s.secondWinterSolstice, s.timezone) {
			return WinterSolstice
		}
	default:
		break
	}
	return NoSolstice
}

func (s SolsticeTable) FirstWinterSolstice() time.Time {
	return s.firstWinterSolstice
}

func (s SolsticeTable) VernalEquinox() time.Time {
	return s.vernalEquinox
}

func (s SolsticeTable) SummerSolstice() time.Time {
	return s.summerSolstice
}

func (s SolsticeTable) AutumnalEquinox() time.Time {
	return s.autumnalEquinox
}

func (s SolsticeTable) SecondWinterSolstice() time.Time {
	return s.secondWinterSolstice
}
