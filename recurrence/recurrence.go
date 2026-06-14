package recurrence

import (
	"github.com/fxtlabs/date"
)

type Recurrence interface {
	NextOccurrence(now date.Date) *date.Date
}
