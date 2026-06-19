package location

import (
	"regexp"

	"github.com/golang/geo/s2"
	"github.com/pkg/errors"
	"github.com/soniakeys/unit"
	"github.com/tzneal/coordconv"
)

type Location struct {
	name      string
	latitude  unit.Angle
	longitude unit.Angle
	tz        TZ
}

func NewLocation(name string, latitude, longitude unit.Angle, tz TZ) Location {
	return Location{
		name:      name,
		latitude:  latitude,
		longitude: longitude,
		tz:        tz,
	}
}

func FromMGRS(name, mgrs string, tz TZ) (Location, error) {
	l, err := coordconv.DefaultMGRSConverter.ConvertToGeodetic(mgrs)
	if err != nil {
		return Location{}, err
	}

	return Location{
		name:      name,
		latitude:  unit.Angle(l.Lat.Radians()),
		longitude: unit.Angle(l.Lng.Radians()),
	}, nil
}

func (l Location) Latitude() unit.Angle {
	return l.latitude
}

func (l Location) Longitude() unit.Angle {
	return l.longitude
}

func (l Location) MGRS(precision int) string {
	c := s2.LatLngFromDegrees(l.latitude.Deg(), l.longitude.Deg())
	m, _ := coordconv.DefaultMGRSConverter.ConvertFromGeodetic(c, precision)

	return m
}

func (l Location) TZ() TZ {
	return l.tz
}

func mgrsPrecision(mgrs string) (int, error) {
	r := regexp.MustCompile(`\d{1,2}[[:alpha:]]{3}(\d+)`)
	match := r.FindStringSubmatch(mgrs)

	if len(match) < 2 {
		return -1, errors.Errorf("attempt to determine precision of invalid MGRS string: %s", mgrs)
	}

	return len(match[1]) / 2, nil
}
