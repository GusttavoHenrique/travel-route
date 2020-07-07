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
	errBestRouteNotFound        = "The best route not found"
	errOriginIsMissing          = "The 'origin' field is mandatory"
	errDestinationIsMissing     = "The 'destination' field is mandatory"
	errPriceIsMissing           = "The attribute 'price' is missing or invalid"
	errSameOriginAndDestination = "The 'origin' field cannot be the same of the 'destination' field"
	errPriceCannotBeNegative    = "The 'price' cannot be negative"
)

// NewRoute create a new route instance
func NewRoute(origin string, destination string, price float64) (*Route, error) {
	route, err := createAndValidateRoute(origin, destination, price, true)
	if err != nil {
		return nil, err
	}

	route.Price = price
	return route, nil
}

// NewBestRoute create a new best route instance
func NewBestRoute(origin string, destination string) (*Route, error) {
	route, err := createAndValidateRoute(origin, destination, 0, false)
	if err != nil {
		return nil, err
	}

	route.Price = math.Inf(-1)
	return route, nil
}

func createAndValidateRoute(origin string, destination string, price float64, validatePrice bool) (*Route, error) {
	err := ValidateRoute(origin, destination, price, validatePrice)
	if err != nil {
		return nil, err
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

// NewBestRoute retrieve a string with the best route
func (route *Route) GetBestRouteStr() (string, error) {
	list := route.BestRoute
	if len(list) == 0 {
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

// ValidateRoute validate the route fields
func ValidateRoute(origin string, destination string, price float64, validatePrice bool) error {
	if origin == "" {
		return errors.Errorf(errOriginIsMissing)
	} else if destination == "" {
		return errors.Errorf(errDestinationIsMissing)
	} else if origin == destination {
		return errors.New(errSameOriginAndDestination)
	} else if price == 0 && validatePrice {
		return errors.Errorf(errPriceIsMissing)
	} else if price <= 0 && validatePrice {
		return errors.New(errPriceCannotBeNegative)
	}
	return nil
}
