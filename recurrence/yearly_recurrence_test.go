package recurrence

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestYearlyRecurrence(t *testing.T) {
	badJSON := []byte(`{
	"period": "YEARLY",
	"rules": {
		"month": 9,
		"weekday": 5
	}
}`)
	var badJSONRule Rule
	err := json.Unmarshal(badJSON, &badJSONRule)
	assert.Error(t, err)

	jsonNoDeath := []byte(`{
	"period": "YEARLY",
	"rules": {
		"month": 9,
		"day": 27
	}
}`)
	var noDeathRule Rule
	err = json.Unmarshal(jsonNoDeath, &noDeathRule)
	assert.NoError(t, err)

	next := noDeathRule.Recurrence().NextOccurrence(testDateJan)
	assert.NotNil(t, next)
	assert.True(t, next.Equal(date.New(2026, time.September, 27)))

	jsonDeath := []byte(`{
	"period": "YEARLY",
	"rules": {
		"month": 9,
		"day": 27
	},
	"until": "2068-09-27"
}`)
	var deathRule Rule
	err = json.Unmarshal(jsonDeath, &deathRule)
	assert.NoError(t, err)

	next = deathRule.Recurrence().NextOccurrence(date.New(2069, time.September, 27))

	assert.Nil(t, next)
}
