package recurrence

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestQuarterlyRecurrence(t *testing.T) {
	qr := NewQuarterlyRecurrence(5, time.Thursday, nil)

	q1Occur := date.New(2026, time.January, 29) // 2026W05-4
	q1Test := qr.NextOccurrence(date.New(2026, time.January, 1))
	assert.True(t, q1Occur.Equal(*q1Test))

	q2Occur := date.New(2026, time.April, 30) // 2026W18-4
	q2Test := qr.NextOccurrence(date.New(2026, time.April, 1))
	assert.True(t, q2Occur.Equal(*q2Test))

	q3Occur := date.New(2026, time.July, 30) // 2026W31-4
	q3Test := qr.NextOccurrence(date.New(2026, time.July, 1))
	assert.True(t, q3Occur.Equal(*q3Test))

	q4Occur := date.New(2026, time.October, 29) // 2026W44-4
	q4Test := qr.NextOccurrence(date.New(2026, time.October, 1))
	assert.True(t, q4Occur.Equal(*q4Test))
}
