package natyr

import (
	"github.com/fxtlabs/date"
)

const avgDaysInYear float64 = 365.2425

type NaturalYear struct {
	birthday date.Date
}

func NewNaturalYear(birthday date.Date) NaturalYear {
	return NaturalYear{birthday: birthday}
}

func (ny NaturalYear) Decade(d date.Date) int {
	dDy := d.Sub(ny.birthday)

	return int(float64(dDy)/(avgDaysInYear*10)) + 1
}

func (ny NaturalYear) Lustrum(d date.Date) int {
	dDy := d.Sub(ny.birthday)

	return int(float64(dDy)/(avgDaysInYear*5)) + 1
}

func (ny NaturalYear) Triad(d date.Date) int {
	dDy := d.Sub(ny.birthday)

	return int(float64(dDy)/(avgDaysInYear*3)) + 1
}

// ToNYDate returns the Natural Year and NY ordinal day.
func (ny NaturalYear) ToNYDate(d date.Date) (int, int) {
	dDy := d.Sub(ny.birthday)

	year := int(float64(dDy) / avgDaysInYear)
	day := dDy - int(float64(year)*avgDaysInYear)

	return year + 1, day
}
