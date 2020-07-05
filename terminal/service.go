package terminal

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
	"travel-route/route"
)

const (
	errMissingArgs         = "The args is missing or invalid"
	errMissingBestRoute    = "The request best route is missing or invalid"
	errReadingInformedFile = "Error reading routes from informed file"
	errSavingRoutes        = "Error saving routes"
)

func GetInputRoute() (*route.Route, error) {
	fmt.Print("please enter the route: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()
	if input == "" {
		return nil, errors.New(errMissingArgs)
	}

	points := strings.Split(input, " - ")
	if len(points) < 2 || points[0] == "" || points[1] == "" {
		return nil, errors.New(errMissingArgs)
	}

	bestRoute, err := route.NewBestRoute(points[0], points[1])
	if bestRoute == nil {
		return nil, errors.New(errMissingBestRoute)
	}

	return bestRoute, err
}

func LoadRoutesFromFile() error {
	var pathFile string
	flag.StringVar(&pathFile, "file", "The file flag is mandatory", "")
	flag.Parse()

	routes, err := route.GetRoutesFromFile(pathFile)
	if err != nil {
		return errors.New(errReadingInformedFile)
	}

	for _, r := range routes {
		err := route.Create(r, pathFile)
		if err != nil {
			return errors.Wrap(err, errSavingRoutes)
		}
	}

	return nil
}

func saveRoutes(routes []*route.Route) {

}
