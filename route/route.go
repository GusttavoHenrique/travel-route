package route

import (
	"github.com/pkg/errors"
	"math"
	"travel-route/point"
)

type Route struct {
	InitialPoint *point.Point
	FinalPoint   *point.Point
	Price        float64
	BestRoute    []*point.Point
}

const (
	errSameOriginAndDestination = "The 'origin' field cannot be the same of the 'destination' field"
	errOriginIsMissing          = "The 'origin' field is mandatory"
	errDestinationIsMissing     = "The 'destination' field is mandatory"
	errPriceIsMissing           = "The attribute 'price' is missing or invalid"
	errBestRouteNotFound        = "The best route not found"
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

	route.Price = price
	return route, nil
}

// NewBestRoute create a new best route instance
func NewBestRoute(origin string, destination string) (*Route, error) {
	route, err := createAndValidateRoute(origin, destination)
	if err != nil {
		return nil, err
	}

	route.Price = math.Inf(-1)
	return route, nil
}

func createAndValidateRoute(origin string, destination string) (*Route, error) {
	if origin == destination {
		return nil, errors.New(errSameOriginAndDestination)
	}

	initialPoint, err := point.NewPoint(origin)
	if err != nil {
		return nil, errors.Wrap(err, errOriginIsMissing)
	}

	finalPoint, err := point.NewPoint(destination)
	if err != nil {
		return nil, errors.Wrap(err, errDestinationIsMissing)
	}

	return &Route{
		InitialPoint: initialPoint,
		FinalPoint:   finalPoint,
	}, nil
}

func (route *Route) GetBestRouteStr() (string, error) {
	list := route.BestRoute
	if len(list) != 0 {
		return "", errors.New(errBestRouteNotFound)
	}

	bestRoute := ""
	for _, r := range list {
		if bestRoute != "" {
			bestRoute = bestRoute + " - "
		}

		bestRoute = bestRoute + r.Name
	}

	return bestRoute, nil
}
