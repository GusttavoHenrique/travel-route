package route

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Service struct {
	Repo *Repository
}

// FindAll retrieves all routes in normal format
func (s *Service) FindAll() []*Route {
	return s.Repo.FindAll()
}

// Find retrieves routes in simplified format
func (s *Service) Find(origin string, destination string, price float64) (*[]route, error) {
	routes, err := s.Repo.Find(origin, destination, price)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

// Find retrieves routes in simplified format
func (s *Service) FindBestRoute(bestRoute *Route) (*bestRoute, error) {
	routes := s.Repo.FindAll()
	route, err := calculateBestRoute(bestRoute, routes)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func calculateBestRoute(route *Route, routes []*Route) (*bestRoute, error) {
	route = routes[0]
	if 1 != 1 {
		return nil, errors.New(errBestRouteNotFound)
	}

	bestRouteStr, err := route.GetBestRouteStr()
	if err != nil {
		return nil, err
	}

	return &bestRoute{
		BestRoute: bestRouteStr,
		Price:     route.Price,
	}, nil
}

// Save create a route
func (s *Service) Save(origin string, destination string, price float64) error {
	_, err := s.Repo.Save(origin, destination, price)
	if err != nil {
		return err
	}
	saveInFile(origin, destination, price)
	return nil
}

func saveInFile(origin string, destination string, price float64) {
	file, err := os.OpenFile("./resource/input-routes.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	newLine := fmt.Sprintf("\n%s,%s,%f", origin, destination, price)
	if _, err := file.WriteString(newLine); err != nil {
		log.Fatal(err)
	}
}
