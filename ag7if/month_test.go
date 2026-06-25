package ag7if

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestComputeMonthStartDate(t *testing.T) {
	// 2026 month start dates
	janStart := date.New(2025, time.December, 29)
	janTest := ComputeMonthStartDate(2026, time.January)
	assert.Equal(t, janStart, janTest)

	febStart := date.New(2026, time.January, 26)
	febTest := ComputeMonthStartDate(2026, time.February)
	assert.Equal(t, febStart, febTest)

	marStart := date.New(2026, time.February, 23)
	marTest := ComputeMonthStartDate(2026, time.March)
	assert.Equal(t, marStart, marTest)

	aprStart := date.New(2026, time.March, 30)
	aprTest := ComputeMonthStartDate(2026, time.April)
	assert.Equal(t, aprStart, aprTest)

	mayStart := date.New(2026, time.April, 27)
	mayTest := ComputeMonthStartDate(2026, time.May)
	assert.Equal(t, mayStart, mayTest)

	junStart := date.New(2026, time.May, 25)
	junTest := ComputeMonthStartDate(2026, time.June)
	assert.Equal(t, junStart, junTest)

	julStart := date.New(2026, time.June, 29)
	julTest := ComputeMonthStartDate(2026, time.July)
	assert.Equal(t, julStart, julTest)

	augStart := date.New(2026, time.July, 27)
	augTest := ComputeMonthStartDate(2026, time.August)
	assert.Equal(t, augStart, augTest)

	sepStart := date.New(2026, time.August, 24)
	sepTest := ComputeMonthStartDate(2026, time.September)
	assert.Equal(t, sepStart, sepTest)

	octStart := date.New(2026, time.September, 28)
	octTest := ComputeMonthStartDate(2026, time.October)
	assert.Equal(t, octStart, octTest)

	novStart := date.New(2026, time.October, 26)
	novTest := ComputeMonthStartDate(2026, time.November)
	assert.Equal(t, novStart, novTest)

	decStart := date.New(2026, time.November, 23)
	decTest := ComputeMonthStartDate(2026, time.December)
	assert.Equal(t, decStart, decTest)
}
