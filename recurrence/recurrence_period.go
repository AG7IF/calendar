package recurrence

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type RP int

const (
	Yearly RP = iota
	Quarterly
	Monthly
	Weekly
)

func ParseRP(s string) (RP, error) {
	switch strings.ToUpper(s) {
	case "YEARLY":
		return Yearly, nil
	case "QUARTERLY":
		return Quarterly, nil
	case "MONTHLY":
		return Monthly, nil
	case "WEEKLY":
		return Weekly, nil
	default:
		return -1, fmt.Errorf("invalid RP: %s", s)
	}
}

func (r RP) String() string {
	switch r {
	case Yearly:
		return "YEARLY"
	case Quarterly:
		return "QUARTERLY"
	case Monthly:
		return "MONTHLY"
	case Weekly:
		return "WEEKLY"
	default:
		panic(fmt.Errorf("invalid RP value: %d", r))
	}
}

func (rp RP) Value() (driver.Value, error) {
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
