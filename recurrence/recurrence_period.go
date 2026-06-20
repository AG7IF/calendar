package recurrence

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type RP int

const (
	RPYears RP = iota
	RPQuarters
	RPMonths
	RPWeeks
)

func ParseRP(s string) (RP, error) {
	switch strings.TrimSpace(strings.ToUpper(s)) {
	case "YEARS":
		return RPYears, nil
	case "QUARTERS":
		return RPQuarters, nil
	case "MONTHS":
		return RPMonths, nil
	case "WEEKS":
		return RPWeeks, nil
	default:
		return -1, errors.Errorf("invalid RP: %s", s)
	}
}

func (rp *RP) String() string {
	switch *rp {
	case RPYears:
		return "YEARS"
	case RPQuarters:
		return "QUARTERS"
	case RPMonths:
		return "MONTHS"
	case RPWeeks:
		return "WEEKS"
	default:
		panic(errors.Errorf("invalid RP value: %d", rp))
	}
}

func (rp *RP) MarshalJSON() ([]byte, error) {
	return []byte(rp.String()), nil
}

func (rp *RP) UnmarshalJSON(raw []byte) error {
	val, err := ParseRP(strings.Trim(string(raw), `"`))
	if err != nil {
		return errors.WithStack(err)
	}

	*rp = val
	return nil
}

func (rp *RP) Value() (driver.Value, error) {
	return []byte(rp.String()), nil
}

func (rp *RP) Scan(raw any) error {
	s, ok := raw.(string)
	if !ok {
		return errors.Errorf("scanned value is not a string: %v", raw)
	}

	val, err := ParseRP(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*rp = val
	return nil
}
