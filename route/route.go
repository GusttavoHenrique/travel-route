package route

import (
	"github.com/pkg/errors"
	"math"
	"travel-route/point"
)

type Route struct {
	initialPoint *point.Point
	finalPoint   *point.Point
	price        float64
}

const (
	errOriginIsMissing      = "The 'origin' field is mandatory"
	errDestinationIsMissing = "The 'destination' field is mandatory"
	errPriceIsMissing       = "The attribute 'price' is missing or invalid"
)

// NewRoute create a new route instance
func NewRoute(origin string, destination string, price float64) (*Route, error) {
	route, err := createAndValidateRoute(origin, destination)
	if err != nil {
		return nil, err
	}

	if price <= 0 {
		return nil, errors.New(errPriceIsMissing)
	}

	route.price = price
	return route, nil
}

// NewBestRoute create a new best route instance
func NewBestRoute(origin string, destination string) (*Route, error) {
	route, err := createAndValidateRoute(origin, destination)
	if err != nil {
		return nil, err
	}

	route.price = math.Inf(-1)
	return route, nil
}

func createAndValidateRoute(origin string, destination string) (*Route, error) {
	initialPoint, err := point.NewPoint(origin)
	if err != nil {
		return nil, errors.Wrap(err, errOriginIsMissing)
	}

	finalPoint, err := point.NewPoint(destination)
	if err != nil {
		return nil, errors.Wrap(err, errDestinationIsMissing)
	}

	return &Route{
		initialPoint: initialPoint,
		finalPoint:   finalPoint,
	}, nil
}
