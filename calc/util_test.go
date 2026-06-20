package calc

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

var testDateJan = date.New(2026, time.January, 1)
var testDateFeb = date.New(2026, time.February, 1)
var testDateMar = date.New(2026, time.March, 1)
var testDateApr = date.New(2026, time.April, 1)
var testDateMay = date.New(2026, time.May, 1)
var testDateJun = date.New(2026, time.June, 1)
var testDateJul = date.New(2026, time.July, 1)
var testDateAug = date.New(2026, time.August, 1)
var testDateSep = date.New(2026, time.September, 1)
var testDateOct = date.New(2026, time.October, 1)
var testDateNov = date.New(2026, time.November, 1)
var testDateDec = date.New(2026, time.December, 1)

func TestComputeFirstGivenWeekdayOfMonth(t *testing.T) {
	firstFebThu := date.New(2026, time.February, 5)
	firstFebThuTest := ComputeFirstGivenWeekdayOfMonth(testDateFeb, time.Thursday)
	assert.True(t, firstFebThu.Equal(firstFebThuTest))

	firstMayThu := date.New(2026, time.May, 7)
	firstMayThuTest := ComputeFirstGivenWeekdayOfMonth(testDateMay, time.Thursday)
	assert.True(t, firstMayThu.Equal(firstMayThuTest))

	firstJunThu := date.New(2026, time.June, 4)
	firstJunThuTest := ComputeFirstGivenWeekdayOfMonth(testDateJun, time.Thursday)
	assert.True(t, firstJunThu.Equal(firstJunThuTest))
}
