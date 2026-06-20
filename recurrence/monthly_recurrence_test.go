package recurrence

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestMonthlyRecurrenceOnDay(t *testing.T) {
	goodJSON := []byte(`{
	"period": "MONTHLY",
	"rules": {
			"day": 31
	}
}`)

	var rule Rule
	err := json.Unmarshal(goodJSON, &rule)
	assert.NoError(t, err)

	janOccur := date.New(2026, time.January, 31)
	testJanOccur := rule.Recurrence().NextOccurrence(testDateJan)
	assert.True(t, janOccur.Equal(*testJanOccur))

	febOccur := date.New(2026, time.February, 28)
	testFebOccur := rule.Recurrence().NextOccurrence(testDateFeb)
	assert.True(t, febOccur.Equal(*testFebOccur))

	marOccur := date.New(2026, time.March, 31)
	testMarOccur := rule.Recurrence().NextOccurrence(testDateMar)
	assert.True(t, marOccur.Equal(*testMarOccur))

	aprOccur := date.New(2026, time.April, 30)
	testAprOccur := rule.Recurrence().NextOccurrence(testDateApr)
	assert.True(t, aprOccur.Equal(*testAprOccur))

	mayOccur := date.New(2026, time.May, 31)
	testMayOccur := rule.Recurrence().NextOccurrence(testDateMay)
	assert.True(t, mayOccur.Equal(*testMayOccur))

	junOccur := date.New(2026, time.June, 30)
	testJunOccur := rule.Recurrence().NextOccurrence(testDateJun)
	assert.True(t, junOccur.Equal(*testJunOccur))

	julOccur := date.New(2026, time.July, 31)
	testJulOccur := rule.Recurrence().NextOccurrence(testDateJul)
	assert.True(t, julOccur.Equal(*testJulOccur))

	augOccur := date.New(2026, time.August, 31)
	testAugOccur := rule.Recurrence().NextOccurrence(testDateAug)
	assert.True(t, augOccur.Equal(*testAugOccur))

	sepOccur := date.New(2026, time.September, 30)
	testSepOccur := rule.Recurrence().NextOccurrence(testDateSep)
	assert.True(t, sepOccur.Equal(*testSepOccur))

	octOccur := date.New(2026, time.October, 31)
	testOctOccur := rule.Recurrence().NextOccurrence(testDateOct)
	assert.True(t, octOccur.Equal(*testOctOccur))

	novOccur := date.New(2026, time.November, 30)
	testNovOccur := rule.Recurrence().NextOccurrence(testDateNov)
	assert.True(t, novOccur.Equal(*testNovOccur))

	decOccur := date.New(2026, time.December, 31)
	testDecOccur := rule.Recurrence().NextOccurrence(testDateDec)
	assert.True(t, decOccur.Equal(*testDecOccur))
}

func TestMonthlyRecurrenceOnWeekday(t *testing.T) {
	badJSON := []byte(`{
	"period": "MONTHLY",
	"rules": {
		"week": 5,
		"weekday": 5,
		"day": 15
	}
}`)

	var badJSONRule Rule
	err := json.Unmarshal(badJSON, &badJSONRule)
	assert.Error(t, err)

	goodJSON := []byte(`{
	"period": "MONTHLY",
	"rules": {
		"week": 2,
		"weekday": 7
	}
	}`)

	var rule Rule
	err = json.Unmarshal(goodJSON, &rule)

	janOccur := date.New(2026, time.January, 11)
	testJanOccur := rule.Recurrence().NextOccurrence(testDateJan)
	assert.True(t, janOccur.Equal(*testJanOccur))

	febOccur := date.New(2026, time.February, 8)
	testFebOccur := rule.Recurrence().NextOccurrence(testDateFeb)
	assert.True(t, febOccur.Equal(*testFebOccur))

	marOccur := date.New(2026, time.March, 8)
	testMarOccur := rule.Recurrence().NextOccurrence(testDateMar)
	assert.True(t, marOccur.Equal(*testMarOccur))

	aprOccur := date.New(2026, time.April, 12)
	testAprOccur := rule.Recurrence().NextOccurrence(testDateApr)
	assert.True(t, aprOccur.Equal(*testAprOccur))

	mayOccur := date.New(2026, time.May, 10)
	testMayOccur := rule.Recurrence().NextOccurrence(testDateMay)
	assert.True(t, mayOccur.Equal(*testMayOccur))

	junOccur := date.New(2026, time.June, 14)
	testJunOccur := rule.Recurrence().NextOccurrence(testDateJun)
	assert.True(t, junOccur.Equal(*testJunOccur))

	julOccur := date.New(2026, time.July, 12)
	testJulOccur := rule.Recurrence().NextOccurrence(testDateJul)
	assert.True(t, julOccur.Equal(*testJulOccur))

	augOccur := date.New(2026, time.August, 9)
	testAugOccur := rule.Recurrence().NextOccurrence(testDateAug)
	assert.True(t, augOccur.Equal(*testAugOccur))

	sepOccur := date.New(2026, time.September, 13)
	testSepOccur := rule.Recurrence().NextOccurrence(testDateSep)
	assert.True(t, sepOccur.Equal(*testSepOccur))

	octOccur := date.New(2026, time.October, 11)
	testOctOccur := rule.Recurrence().NextOccurrence(testDateOct)
	assert.True(t, octOccur.Equal(*testOctOccur))

	novOccur := date.New(2026, time.November, 8)
	testNovOccur := rule.Recurrence().NextOccurrence(testDateNov)
	assert.True(t, novOccur.Equal(*testNovOccur))

	decOccur := date.New(2026, time.December, 13)
	testDecOccur := rule.Recurrence().NextOccurrence(testDateDec)
	assert.True(t, decOccur.Equal(*testDecOccur))
}

func TestMonthlyRecurrence_FifthWeek(t *testing.T) {
	fifthWkJSON := []byte(`{
		"period": "MONTHLY",
		"rules": {
			"week": 5,
			"weekday": 4
		}
		}`)

	var fifthWkRule Rule
	err := json.Unmarshal(fifthWkJSON, &fifthWkRule)
	assert.NoError(t, err)

	janFifthWkOccur := date.New(2026, time.January, 29)
	testJanFifthWkOccur := fifthWkRule.Recurrence().NextOccurrence(testDateJan)
	assert.True(t, janFifthWkOccur.Equal(*testJanFifthWkOccur))

	aprFifthWkOccur := date.New(2026, time.April, 30)
	testAprFifthWkOccur := fifthWkRule.Recurrence().NextOccurrence(testDateFeb)
	assert.True(t, aprFifthWkOccur.Equal(*testAprFifthWkOccur))

	julFifthWkOccur := date.New(2026, time.July, 30)
	testJulFifthWkOccur := fifthWkRule.Recurrence().NextOccurrence(testDateMay)
	assert.True(t, julFifthWkOccur.Equal(*testJulFifthWkOccur))

	octFifthWkOccur := date.New(2026, time.October, 29)
	testOctFifthWkOccur := fifthWkRule.Recurrence().NextOccurrence(testDateAug)
	assert.True(t, octFifthWkOccur.Equal(*testOctFifthWkOccur))

	decFifthWkOccur := date.New(2026, time.December, 31)
	testDecFifthWkOccur := fifthWkRule.Recurrence().NextOccurrence(testDateNov)
	assert.True(t, decFifthWkOccur.Equal(*testDecFifthWkOccur))
}
