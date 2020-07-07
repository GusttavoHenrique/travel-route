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
		c.HttpServer.SetResponseInternalError(w)
	}
}

// PostRoute the handler to post route
func (c *Controller) PostRoute(w http2.ResponseWriter, r *http2.Request) {
	if err := c.HttpServer.ValidPost(w, r); err != nil {
		return
	}

	request := c.getBody(r)
	if err := c.service.Save(request.Origin, request.Destination, request.Price); err != nil {
		c.HttpServer.SetResponseCreated(w)
	} else {
		c.HttpServer.SetResponseInternalError(w)
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

func (c *Controller) getBody(r *http2.Request) *route {
	var request *route

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Print(err.Error(), c.HttpServer.StatusBadRequest())
		return nil
	}

	return request
}
