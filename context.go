package calendar

import (
	"fmt"
)

type Context struct {
	Code        string    `yaml:"code"`
	Title       string    `yaml:"title"`
	Subcontexts []Context `yaml:"subcontexts,omitempty"`
}

func (c Context) String() string {
	return fmt.Sprintf("%s: %s", c.Code, c.Title)
}
