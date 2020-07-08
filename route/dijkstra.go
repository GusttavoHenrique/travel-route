package route

import (
	"github.com/pkg/errors"
	"travel-route/point"
)

func calculateBestRoute(initialRoute *Route, routes []*Route) (*bestRoute, error) {
	// Todo calculate the best route
	initialRoute.InitialPoint.Name = "SAO"
	initialRoute.FinalPoint.Name = "NAT"
	initialRoute.BestRoute = append(initialRoute.BestRoute, initialRoute.InitialPoint)
	otherPoint := &point.Point{Name: "RIO"}
	initialRoute.BestRoute = append(initialRoute.BestRoute, otherPoint)
	initialRoute.BestRoute = append(initialRoute.BestRoute, initialRoute.InitialPoint)
	initialRoute.Price = 10

	if 1 != 1 {
		return nil, errors.New(errBestRouteNotFound)
	}

	bestRouteStr, err := initialRoute.GetBestRouteStr()
	if err != nil {
		return nil, err
	}

	return &bestRoute{
		BestRoute: bestRouteStr,
		Price:     initialRoute.Price,
	}, nil
}

func pop(routes []*Route) []*Route {
	length := len(routes) - 1
	routes = routes[:length]
	return routes
}

func push(routes []*Route, route *Route) []*Route {
	return append(routes, route)
}

func neighbors(routes []*Route, route *Route) {
	// Todo calculate all neighbors for all routes
}
