package route

import (
	"testing"
)

const (
	origin      = "SAO"
	destination = "NAT"
	price       = float64(50)
)

func TestRouteCreation(t *testing.T) {
	route, err := createAndValidateRoute(origin, destination)
	validateInstanceRoute(route, err, t)
	if route.price != 0 {
		t.Errorf("Unexpected price.")
	}
}

func TestNewRoute(t *testing.T) {
	newRoute, err := NewRoute(origin, destination, price)
	validateInstanceRoute(newRoute, err, t)

	if newRoute.initialPoint == nil {
		t.Errorf("Error initial point is nil.")
	} else if newRoute.finalPoint == nil {
		t.Errorf("Error final point is nil.")
	} else if newRoute.price != price {
		t.Errorf("Error actual price = %v, and expected price = %v.", newRoute.price, price)
	}
}

func TestNewRouteWithInvalidParams(t *testing.T) {
	routeWithoutOrigin, err := NewRoute("", destination, price)
	validateError(routeWithoutOrigin, err, t)

	newRouteWithoutDestination, err := NewRoute(origin, "", price)
	validateError(newRouteWithoutDestination, err, t)

	routeWithoutPrice, err := NewRoute(origin, "", price)
	validateError(routeWithoutPrice, err, t)
}

func TestNewBestRoute(t *testing.T) {
	newRoute, err := NewBestRoute(origin, destination)
	validateInstanceRoute(newRoute, err, t)

	if newRoute.initialPoint == nil {
		t.Errorf("Error initial point is nil.")
	} else if newRoute.finalPoint == nil {
		t.Errorf("Error final point is nil.")
	} else if newRoute.price >= 0 {
		t.Errorf("Unexpected positive price.")
	}
}

func TestNewBestRouteWithInvalidParams(t *testing.T) {
	routeWithoutOrigin, err := NewBestRoute("", destination)
	validateError(routeWithoutOrigin, err, t)

	newRouteWithoutDestination, err := NewBestRoute(origin, "")
	validateError(newRouteWithoutDestination, err, t)
}

func validateError(route *Route, err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error.")
	} else if route != nil {
		t.Error("Unexpected new route instance.")
	}
}

func validateInstanceRoute(route *Route, err error, t *testing.T) {
	if err != nil {
		t.Error("Unexpected error.")
	} else if route == nil {
		t.Error("Expected new route instance.")
	}
}
