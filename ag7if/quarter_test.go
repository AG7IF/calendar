package ag7if

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"

	"github.com/ag7if/calendar/calendar"
)

func TestComputeQuarterStartDate(t *testing.T) {
	// 2026 quarter start dates

	q1Start := date.New(2025, time.December, 29)
	q1Test := ComputeQuarterStartDate(2026, calendar.Q1)
	assert.Equal(t, q1Start, q1Test)

	q2Start := date.New(2026, time.March, 30)
	q2Test := ComputeQuarterStartDate(2026, calendar.Q2)
	assert.Equal(t, q2Start, q2Test)

	q3Start := date.New(2026, time.June, 29)
	q3Test := ComputeQuarterStartDate(2026, calendar.Q3)
	assert.Equal(t, q3Start, q3Test)

	q4Start := date.New(2026, time.September, 28)
	q4Test := ComputeQuarterStartDate(2026, calendar.Q4)
	assert.Equal(t, q4Start, q4Test)
}
