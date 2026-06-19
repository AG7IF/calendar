package location

import (
	"regexp"
	"time"

	"github.com/golang/geo/s2"
	"github.com/pkg/errors"
	"github.com/soniakeys/unit"
	"github.com/tzneal/coordconv"
)

type Location struct {
	name      string
	latitude  unit.Angle
	longitude unit.Angle
	timezone  *time.Location
}

func NewLocation(name string, latitude, longitude unit.Angle, timezone *time.Location) Location {
	return Location{
		name:      name,
		latitude:  latitude,
		longitude: longitude,
		timezone:  timezone,
	}
}

func FromMGRS(name, mgrs string, timezone *time.Location) (Location, error) {
	l, err := coordconv.DefaultMGRSConverter.ConvertToGeodetic(mgrs)
	if err != nil {
		return Location{}, err
	}

	return Location{
		name:      name,
		latitude:  unit.Angle(l.Lat.Radians()),
		longitude: unit.Angle(l.Lng.Radians()),
		timezone:  timezone,
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

func (l Location) Timezone() *time.Location {
	return l.timezone
}

func mgrsPrecision(mgrs string) (int, error) {
	r := regexp.MustCompile(`\d{1,2}[[:alpha:]]{3}(\d+)`)
	match := r.FindStringSubmatch(mgrs)

	if len(match) < 2 {
		return -1, errors.Errorf("attempt to determine precision of invalid MGRS string: %s", mgrs)
	}

	return len(match[1]) / 2, nil
}
