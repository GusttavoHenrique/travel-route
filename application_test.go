package main

import (
	"os"
	"testing"
	"travel-route/route"
)

func TestLoadBestRoute(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"nil args", []string{}},
		{"empty args", []string{"", "", ""}},
		{"only application", []string{"./application.go", "", ""}},
		{"only application and origin", []string{"./application.go", "SAO", ""}},
		{"only application and destination", []string{"./application.go", "", "NAT"}},
		{"only application and origin", []string{"", "SAO", "NAT"}},
		{"with all args", []string{"./application.go", "", "NAT"}},
		{"with all args", []string{"./application.go", "SAO", "NAT"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			os.Args = test.args
			bestRoute, err := loadBestRoute()

			if err != nil {
				validateError(bestRoute, err, t)
			} else {
				validateInstanceRoute(bestRoute, err, t)
			}
		})
	}
}

func validateError(route *route.Route, err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error.")
	} else if route != nil {
		t.Error("Unexpected new route instance.")
	}
}

func validateInstanceRoute(route *route.Route, err error, t *testing.T) {
	if err != nil {
		t.Error("Unexpected error.")
	} else if route == nil {
		t.Error("Expected new route instance.")
	}
}
