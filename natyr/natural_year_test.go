package natyr

import (
	"testing"
	"time"

	"github.com/fxtlabs/date"
	"github.com/stretchr/testify/assert"
)

func TestNaturalYear_Decade(t *testing.T) {
	ny := NewNaturalYear(date.New(1988, time.September, 27))

	decade := ny.Decade(date.New(2025, time.October, 1))

	assert.Equal(t, 4, decade)
}

func TestNaturalYear_Lustrum(t *testing.T) {
	ny := NewNaturalYear(date.New(1988, time.September, 27))

	lustrum := ny.Lustrum(date.New(2025, time.October, 1))

	assert.Equal(t, 8, lustrum)
}

func TestNaturalYear_Triad(t *testing.T) {
	ny := NewNaturalYear(date.New(1988, time.September, 27))

	triad := ny.Triad(date.New(2025, time.October, 1))

	assert.Equal(t, 13, triad)
}

func TestNaturalYear_ToNYDate(t *testing.T) {
	ny := NewNaturalYear(date.New(1988, time.September, 27))

	year, day := ny.ToNYDate(date.New(2025, time.October, 1))

	assert.Equal(t, 38, year, "wrong year")
	assert.Equal(t, 5, day, "wrong day")

	year, day = ny.ToNYDate(date.New(2025, time.September, 27))

	assert.Equal(t, 38, year, "wrong year")
	assert.Equal(t, 1, day, "wrong day")

	year, day = ny.ToNYDate(date.New(2024, time.September, 27))

	assert.Equal(t, 37, year, "wrong year")
	assert.Equal(t, 1, day, "wrong day")

	year, day = ny.ToNYDate(date.New(2000, time.September, 27))

	assert.Equal(t, 13, year, "wrong year")
	assert.Equal(t, 1, day, "wrong day")
}
