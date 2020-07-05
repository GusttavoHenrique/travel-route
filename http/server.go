package http

import (
	"log"
	"net/http"
	"travel-route/route"
)

type Server struct {
	Writer  http.ResponseWriter
	request *http.Request
}

func (s *Server) Config() {
	s.Writer.Header().Set("Content-Type", "application/json")

	switch s.request.Method {
	case "GET":
		http.HandleFunc("/routes", route.GetRoutes)
	/*case "POST":
	http.HandleFunc("/routes", route.PostRoute)*/
	default:
		s.Writer.WriteHeader(http.StatusNotFound)
		s.Writer.Write([]byte(`{"message": "not found"}`))
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
