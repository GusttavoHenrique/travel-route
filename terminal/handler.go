package terminal

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"travel-route/route"
)

const (
	errMissingArgs      = "The args is missing or invalid"
	errFileNotInformed  = "The file path was not informed"
	errEmptyRoutes      = "The informed file is empty"
	errMissingBestRoute = "The request best route is missing or invalid"
)

type Terminal struct {
	service      *Service
	routeService *route.Service
}

// NewTerminal create a new instance of terminal service
func NewTerminal(routeService *route.Service) *Terminal {
	return &Terminal{
		service:      &Service{routeRepo: routeService.Repo},
		routeService: routeService,
	}
}

// GetInputRoute take the input request form terminal command line
func (t *Terminal) GetInputRoute() (*route.Route, error) {
	fmt.Print("\nplease enter the route: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	points, err := validateInput(scan.Text())
	if err != nil {
		return nil, err
	}

	origin := points[0]
	destination := points[1]
	if err = route.ValidateRoute(origin, destination, 0, false); err != nil {
		return nil, err
	}

	bestRoute, err := route.NewBestRoute(origin, destination)
	if bestRoute == nil {
		return nil, errors.New(errMissingBestRoute)
	}

	return bestRoute, err
}

func validateInput(input string) ([]string, error) {
	if input == "" {
		return nil, errors.New(errMissingArgs)
	}

	points := strings.Split(strings.TrimSpace(input), "-")
	if len(points) < 2 || points[0] == "" || points[1] == "" {
		return nil, errors.New(errMissingArgs)
	}

	return points, nil
}

// LoadRoutesFromFile load routes from informed file
func (t *Terminal) LoadRoutesFromFile(pathFile string) error {
	if pathFile == "" {
		return errors.New(errFileNotInformed)
	}

	_, err := t.service.saveRoutesFromFile(pathFile)
	if err != nil {
		return errors.New(errLoadingRoutesFromFile)
	}

	return nil
}

// TerminalListener create a listener for terminal command line
func (t *Terminal) TerminalListener() {
	for {
		inputRoute, err := t.GetInputRoute()

		if err != nil {
			fmt.Printf("%s\n", err)
		} else {
			bestRoute, err := t.routeService.FindBestRoute(inputRoute)
			if err != nil {
				fmt.Printf("%s\n", err)
			} else {
				fmt.Printf("best route: %s > %f\n", bestRoute.BestRoute, bestRoute.Price)
			}
		}
	}
}
