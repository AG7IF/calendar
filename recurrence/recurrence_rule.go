package recurrence

import (
	"encoding/json"
	"time"

	"github.com/fxtlabs/date"
	"github.com/pkg/errors"
)

const (
	dayKey     = "day"
	monthKey   = "month"
	weekKey    = "week"
	weekdayKey = "weekday"
)

type Rule struct {
	period RP
	rules  map[string]int
	until  *date.Date
}

func (r *Rule) Recurrence() Recurrence {
	switch r.period {
	case Yearly:
		month := time.Month(r.rules[monthKey])
		day := r.rules[dayKey]
		return NewYearlyRecurrence(month, day, r.until)

	case Quarterly:
		// because time.Sunday == 0, and the standard I am using is the ISO 8601 standard
		// (where Sunday == 7), we simply mod 7 to force the ISO sunday into the time.Sunday value.
		weekday := time.Weekday(r.rules[weekdayKey] % 7)
		week := r.rules[weekKey]

		return NewQuarterlyRecurrence(week, weekday, r.until)
	case Monthly:
		return nil
	case Weekly:
		return nil
	}
	return nil
}

func (r *Rule) MarshalJSON() ([]byte, error) {
	marshal := struct {
		Period RP             `json:"period"`
		Rules  map[string]int `json:"rules"`
		Until  *date.Date     `json:"until,omitempty"`
	}{
		Period: r.period,
		Rules:  r.rules,
		Until:  r.until,
	}

	return json.Marshal(marshal)
}

func (r *Rule) UnmarshalJSON(raw []byte) error {
	unmarshal := struct {
		Period RP             `json:"period"`
		Rules  map[string]int `json:"rules"`
		Until  *date.Date     `json:"until,omitempty"`
	}{}

	err := json.Unmarshal(raw, &unmarshal)
	if err != nil {
		return errors.WithStack(err)
	}

	switch unmarshal.Period {
	case Yearly:
		_, ok := unmarshal.Rules[monthKey]
		if !ok {
			return errors.Errorf("yearly recurrence rules do not define key: '%s'", monthKey)
		}

		_, ok = unmarshal.Rules[dayKey]
		if !ok {
			return errors.Errorf("yearly recurrence rules do not define key: '%s'", dayKey)
		}
	case Quarterly:
		_, ok := unmarshal.Rules[weekKey]
		if !ok {
			return errors.Errorf("yearly recurrence rules do not define key: '%s'", weekKey)
		}
		_, ok = unmarshal.Rules[weekdayKey]
		if !ok {
			return errors.Errorf("yearly recurrence rules do not define key: '%s'", weekKey)
		}
	case Monthly:
		return errors.New("monthly recurrence not implemented")
	case Weekly:
		return errors.New("weekly recurrence not implemented")
	}

	r.period = unmarshal.Period
	r.rules = unmarshal.Rules
	r.until = unmarshal.Until

	return nil
}
