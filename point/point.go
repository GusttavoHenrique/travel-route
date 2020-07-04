package point

import (
	"github.com/pkg/errors"
	"math"
)

type Point struct {
	Name     string
	estimate float64
}

const errMissingName = "The attribute 'name' is missing or invalid"

// NewPoint create a new point instance
func NewPoint(name string) (*Point, error) {
	if name == "" {
		return nil, errors.New(errMissingName)
	}

	return &Point{
		Name:     name,
		estimate: math.Inf(1),
	}, nil
}
