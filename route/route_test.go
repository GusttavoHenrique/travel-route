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
	validateNewRoute(route, err, t)

	if route.initialPoint.Name == route.finalPoint.Name {
		t.Error(errSameOriginAndDestination)
	}
}

func TestNewRoute(t *testing.T) {
	tests := []struct {
		name        string
		origin      string
		destination string
		price       float64
	}{
		{"all params", origin, destination, price},
		{"only origin", origin, "", 0},
		{"only destination", "", destination, 0},
		{"only price", "", "", price},
		{"only origin and destination", origin, destination, 0},
		{"only origin and price", origin, "", price},
		{"only destination and price", "", destination, price},
		{"no params", "", "", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newRoute, err := NewRoute(test.origin, test.destination, test.price)

			if err != nil {
				validateError(newRoute, err, t)
			} else {
				validateNewRoute(newRoute, err, t)

				if newRoute.initialPoint.Name != origin {
					t.Errorf("Error actual = %v, and expected = %v.", newRoute.initialPoint.Name, origin)
				} else if newRoute.finalPoint.Name != destination {
					t.Errorf("Error actual = %v, and expected = %v.", newRoute.finalPoint.Name, destination)
				} else if newRoute.price == 0 {
					t.Fatal("Expected price in new route instance.")
				} else if newRoute.price != price {
					t.Errorf("Error actual = %v, and expected = %v.", newRoute.price, price)
				}
			}
		})
	}
}

func TestBestRoute(t *testing.T) {
	tests := []struct {
		name        string
		origin      string
		destination string
	}{
		{"all params", origin, destination},
		{"only origin", origin, ""},
		{"only destination", "", destination},
		{"no params", "", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newRoute, err := NewBestRoute(test.origin, test.destination)

			if err != nil {
				validateError(newRoute, err, t)
			} else {
				validateNewRoute(newRoute, err, t)

				if newRoute.initialPoint.Name != origin {
					t.Errorf("Error actual = %v, and expected = %v.", newRoute.initialPoint.Name, origin)
				} else if newRoute.finalPoint.Name != destination {
					t.Errorf("Error actual = %v, and expected = %v.", newRoute.finalPoint.Name, destination)
				} else if newRoute.price == 0 {
					t.Fatal("Expected price in new route instance.")
				} else if newRoute.price >= 0 {
					t.Errorf("Unexpected positive price.")
				}
			}
		})
	}
}

func validateError(route *Route, err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error.")
	} else if route != nil {
		t.Fatal("Unexpected new route instance.")
	}
}

func validateNewRoute(route *Route, err error, t *testing.T) {
	if err != nil {
		t.Error("Unexpected error.")
	} else if route == nil {
		t.Fatal("Expected new route instance.")
	} else if route.initialPoint == nil {
		t.Fatal("Expected final point in new route instance.")
	} else if route.finalPoint == nil {
		t.Fatal("Expected initial point in new route instance.")
	}
}
