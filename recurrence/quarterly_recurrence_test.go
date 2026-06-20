package recurrence

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestQuarterlyRecurrence(t *testing.T) {
	badJSON := []byte(`{
	"period": "QUARTERS",
	"start": "1988-09-27",
	"rules": {
		"month": 9,
		"day": 27
	}
}`)
	var badJSONRule Rule
	err := json.Unmarshal(badJSON, &badJSONRule)
	assert.Error(t, err)

	goodJSON := []byte(`{
	"period": "QUARTERS",
	"start": "2026-01-29",
	"rules": {
		"week": 5,
		"weekday": 4
	}
}`)
	var goodJSONRule Rule
	err = json.Unmarshal(goodJSON, &goodJSONRule)
	assert.NoError(t, err)

	qr := goodJSONRule.Recurrence()

	q1Occur := date.New(2026, time.January, 29) // 2026W05-4
	q1Test := qr.NextOccurrence(testDateJan)
	assert.True(t, q1Occur.Equal(*q1Test))

	q2Occur := date.New(2026, time.April, 30) // 2026W18-4
	q2Test := qr.NextOccurrence(testDateApr)
	assert.True(t, q2Occur.Equal(*q2Test))

	q3Occur := date.New(2026, time.July, 30) // 2026W31-4
	q3Test := qr.NextOccurrence(testDateJul)
	assert.True(t, q3Occur.Equal(*q3Test))

	q4Occur := date.New(2026, time.October, 29) // 2026W44-4
	q4Test := qr.NextOccurrence(testDateOct)
	assert.True(t, q4Occur.Equal(*q4Test))
}
