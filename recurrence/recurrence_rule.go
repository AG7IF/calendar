package recurrence

import (
	"encoding/json"
	"maps"
	"slices"
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

const (
	dayKey     = "day"
	monthKey   = "month"
	onDayKey   = "onDay"
	weekKey    = "week"
	weekdayKey = "weekday"
)

type Rule struct {
	every  int
	period RP
	start  date.Date
	until  *date.Date
	rules  map[string]int
}

func (r *Rule) Recurrence() Recurrence {
	switch r.period {
	case RPYears:
		month := time.Month(r.rules[monthKey])
		day := r.rules[dayKey]
		return NewYearlyRecurrence(month, day, r.until)

	case RPQuarters:
		// because time.Sunday == 0, and the standard I am using is the ISO 8601 standard
		// (where Sunday == 7), we simply mod 7 to force the ISO sunday into the time.Sunday value.
		weekday := time.Weekday(r.rules[weekdayKey] % 7)
		week := r.rules[weekKey]

		return NewQuarterlyRecurrence(week, weekday, r.until)
	case RPMonths:
		if r.rules[onDayKey] == 1 {
			day := r.rules[dayKey]
			return NewOnDayMonthlyRecurrence(day, r.until)
		}

		week := r.rules[weekKey]
		weekday := time.Weekday(r.rules[weekdayKey] % 7)

		return NewWeekdayMonthlyRecurrence(week, weekday, r.until)
	case RPWeeks:
		return nil
	}
	return nil
}

func (r *Rule) MarshalJSON() ([]byte, error) {
	_, ok := r.rules[onDayKey]
	if ok {
		delete(r.rules, onDayKey)
	}

	marshal := struct {
		Every  int            `json:"every,omitempty"`
		Period RP             `json:"period"`
		Start  date.Date      `json:"start"`
		Until  *date.Date     `json:"until,omitempty"`
		Rules  map[string]int `json:"rules"`
	}{
		Every:  r.every,
		Period: r.period,
		Start:  r.start,
		Until:  r.until,
		Rules:  r.rules,
	}

	return json.Marshal(marshal)
}

func (r *Rule) UnmarshalJSON(raw []byte) error {
	unmarshal := struct {
		Every  int            `json:"every,omitempty"`
		Period RP             `json:"period"`
		Start  date.Date      `json:"start"`
		Until  *date.Date     `json:"until,omitempty"`
		Rules  map[string]int `json:"rules"`
	}{}

	err := json.Unmarshal(raw, &unmarshal)
	if err != nil {
		return errors.WithStack(err)
	}

	rulesKeys := slices.Collect(maps.Keys(unmarshal.Rules))

	switch unmarshal.Period {
	case RPYears:
		if !(slices.Contains(rulesKeys, monthKey) && slices.Contains(rulesKeys, dayKey)) {
			return errors.Errorf("rules has the following keys %v, but [\"%s\",\"%s\"] are required.", rulesKeys, monthKey, dayKey)
		}
	case RPQuarters:
		if !(slices.Contains(rulesKeys, weekKey) && slices.Contains(rulesKeys, weekdayKey)) {
			return errors.Errorf("rules has the following keys %v, but [\"%s\",\"%s\"] are required.", rulesKeys, weekKey, weekdayKey)
		}
	case RPMonths:
		if slices.Contains(rulesKeys, dayKey) {
			if slices.Contains(rulesKeys, weekKey) || slices.Contains(rulesKeys, weekdayKey) {
				return errors.New("monthly recurrences can specify a day or a week/weekday combination, not both.")
			}

			unmarshal.Rules[onDayKey] = 1

		} else if slices.Contains(rulesKeys, weekKey) && slices.Contains(rulesKeys, weekdayKey) {
			if slices.Contains(rulesKeys, dayKey) {
				return errors.New("monthly recurrences can specify a day or a week/weekday combination, not both.")
			}

			unmarshal.Rules[onDayKey] = 0

		} else {
			return errors.Errorf("rules has the following keys %v, but [\"%s\"] or [\"%s\",\"%s\"] are required.", rulesKeys, dayKey, weekKey, weekdayKey)
		}
	case RPWeeks:
		return errors.New("weekly recurrence not implemented")
	}

	r.period = unmarshal.Period
	r.rules = unmarshal.Rules
	r.until = unmarshal.Until

	return nil
}
