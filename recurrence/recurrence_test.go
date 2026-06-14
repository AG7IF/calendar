package recurrence

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestYearlyRecurrence(t *testing.T) {
	dead := date.New(2068, time.September, 27)
	rhbd := NewYearlyRecurrence(time.September, 27, nil)
	rhbdud := NewYearlyRecurrence(time.September, 27, &dead)

	now := date.New(1988, time.September, 27)

	next := rhbd.NextOccurrence(now)

	assert.NotNil(t, next)
	assert.True(t, next.Equal(date.New(1989, time.September, 27)))

	now = date.New(2069, time.September, 27)
	next = rhbdud.NextOccurrence(now)

	assert.Nil(t, next)
}
