package route

import (
	"github.com/pkg/errors"
)

type Repository struct {
	Db *Db
}

const (
	errDuplicatedRoutesSamePrice     = "Found duplicated routes with same prices"
	errDuplicatedRoutesDistinctPrice = "Found duplicated routes with distinct prices"
	errInitRoute                     = "Error initializing route"
	errRoutesNotFound                = "Routes not found"
)

// Save save route in database
func (repo *Repository) Save(origin string, destination string, price float64) ([]*Route, error) {
	err := repo.validateAndSave(origin, destination, price)
	if err != nil {
		return nil, err
	}

	routes := repo.Db.FindRoutes()
	return routes, nil
}

func (repo *Repository) validateAndSave(origin string, destination string, price float64) error {
	routes := repo.Db.FindRoutes()

	err := func(rs []*Route) error {
		for _, r := range rs {
			if r.InitialPoint.Name == origin && r.FinalPoint.Name == destination {
				if r.Price == price {
					return errors.New(errDuplicatedRoutesSamePrice)
				}
				return errors.New(errDuplicatedRoutesDistinctPrice)
			}
		}
		return nil
	}(routes)

	if err != nil {
		return err
	}

	newRoute, err := NewRoute(origin, destination, price)
	if err != nil {
		return errors.New(errInitRoute)
	}
	repo.Db.SaveRoute(newRoute)
	return nil
}

// FindAll retrieves saved routes
func (repo *Repository) FindAll() []*Route {
	return repo.Db.FindRoutes()
}

// Find all saved routes in database
func (repo *Repository) Find(origin string, destination string, price float64) (*[]route, error) {
	result := make([]route, 0)

	tables := repo.Db.FindRoutes()
	for _, r := range tables {
		success := true
		validateField(r.InitialPoint.Name, origin, &success, origin != "")
		validateField(r.FinalPoint.Name, destination, &success, destination != "")
		validateField(r.Price, price, &success, price > 0)

		if success {
			route := route{
				Origin:      r.InitialPoint.Name,
				Destination: r.FinalPoint.Name,
				Price:       r.Price,
			}
			result = append(result, route)
		}
	}

	if len(result) == 0 {
		return nil, errors.New(errRoutesNotFound)
	}

	return &result, nil
}

func validateField(tableField interface{}, field interface{}, success *bool, validate bool) {
	if validate {
		*success = *success && tableField == field
	}
}
