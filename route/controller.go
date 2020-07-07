package route

import (
	"encoding/json"
	"log"
	http2 "net/http"
	"strconv"
	"travel-route/http"
)

type Controller struct {
	HttpServer *http.Server
	service    *Service
}

type route struct {
	Origin      string  `json:"origin"`
	Destination string  `json:"destination"`
	Price       float64 `json:"price"`
}

type bestRoute struct {
	BestRoute string  `json:"best-route"`
	Price     float64 `json:"price"`
}

// NewController create a new route service instance
func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}

// PostRoute the handler to get route
func (c *Controller) GetRoutes(w http2.ResponseWriter, r *http2.Request) {
	if err := c.HttpServer.ValidGet(w, r); err != nil {
		return
	}

	request := c.getQueryParams(r)
	if routes, err := c.service.Find(request.Origin, request.Destination, request.Price); err == nil {
		c.HttpServer.SetResponse(w, routes, c.HttpServer.StatusOk())
	} else {
		c.HttpServer.SetResponseNotFound(w, err.Error())
	}
}

// GetBestRoute the handler to get best route
func (c *Controller) GetBestRoute(w http2.ResponseWriter, r *http2.Request) {
	if err := c.HttpServer.ValidGet(w, r); err != nil {
		return
	}

	request := c.getQueryParams(r)
	bestRoute, err := NewBestRoute(request.Origin, request.Destination)
	if err != nil {
		c.HttpServer.SetResponseBadRequest(w, err.Error())
		return
	}

	if routes, err := c.service.FindBestRoute(bestRoute); err == nil {
		c.HttpServer.SetResponse(w, routes, c.HttpServer.StatusOk())
	} else {
		c.HttpServer.SetResponseNotFound(w, err.Error())
	}
}

// PostRoute the handler to post route
func (c *Controller) PostRoute(w http2.ResponseWriter, r *http2.Request) {
	if err := c.HttpServer.ValidPost(w, r); err != nil {
		return
	}

	request, err := c.getBody(r)
	if err != nil {
		c.HttpServer.SetResponseBadRequest(w, err.Error())
		return
	}

	if err := c.service.Save(request.Origin, request.Destination, request.Price); err == nil {
		c.HttpServer.SetResponseCreated(w)
	} else {
		c.HttpServer.SetResponseBadRequest(w, err.Error())
	}
}

func (c *Controller) getQueryParams(r *http2.Request) *route {
	origin := c.getQueryParam(r, "origin")
	destination := c.getQueryParam(r, "destination")
	priceStr := c.getQueryParam(r, "price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		price = 0
	}

	result := &route{
		Origin:      origin,
		Destination: destination,
		Price:       price,
	}

	return result
}

func (c *Controller) getQueryParam(r *http2.Request, key string) string {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return ""
	}

	return keys[0]
}

func (c *Controller) getBody(r *http2.Request) (*route, error) {
	var request *route

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Print(err.Error(), c.HttpServer.StatusBadRequest())
		return nil, err
	}

	err = ValidateRoute(request.Origin, request.Destination, request.Price, true)
	if err != nil {
		return nil, err
	}

	return request, nil
}
