package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type DefaultResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

const (
	messageNotFound         = "Not found"
	messageInternalError    = "Internal error"
	messageMethodNotAllowed = "Method not allowed"
)

func NewServer() *Server {
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		httpServer: s,
	}
}

// RegisterEndpoint register handler to endpoint
func (s *Server) RegisterEndpoint(method string, path string, handler func(writer http.ResponseWriter, reader *http.Request)) {
	s.Handler(method, path, handler)
}

// RegisterEndpoint switch to handler endpoint by http method
func (s *Server) Handler(method string, path string, handler func(writer http.ResponseWriter, reader *http.Request)) {
	switch method {
	case "GET":
		http.HandleFunc(path, handler)
	case "POST":
		http.HandleFunc(path, handler)
	default:
		http.HandleFunc(path, handler)
	}
}

// DefaultHandler generate the default response handler
func (s *Server) DefaultHandler(writer http.ResponseWriter, reader *http.Request) {
	s.SetResponseError(writer, messageNotFound, http.StatusNotFound)
}

// Start start the http server listener
func (s *Server) Start() {
	log.Fatal(s.httpServer.ListenAndServe())
}

// ValidPost validate request http POST method by handler
func (s *Server) ValidPost(writer http.ResponseWriter, reader *http.Request) error {
	return s.validate(writer, reader, "POST")
}

// ValidGet validate request http GET method by handler
func (s *Server) ValidGet(writer http.ResponseWriter, reader *http.Request) error {
	return s.validate(writer, reader, "GET")
}

func (s *Server) validate(writer http.ResponseWriter, reader *http.Request, method string) error {
	if reader.Method != method {
		s.SetResponseError(writer, messageMethodNotAllowed, http.StatusMethodNotAllowed)
		return errors.New(messageMethodNotAllowed)
	}
	s.setHeaders(writer)
	return nil
}

func (s *Server) setHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

// SetResponse defines the
func (s *Server) SetResponse(writer http.ResponseWriter, response interface{}, status int) {
	if response != nil {
		var body []byte

		body, err := convert(response)
		if err != nil {
			status = http.StatusInternalServerError
			response = &DefaultResponse{
				Message: messageInternalError,
				Code:    status,
			}

			body, _ = convert(response)
			s.setStatus(writer, status)
		} else {
			s.setStatus(writer, status)
		}
		_, _ = writer.Write(body)
	} else {
		s.setStatus(writer, status)
	}

	_, _ = fmt.Fprint(writer)
}

func (s *Server) SetResponseCreated(writer http.ResponseWriter) {
	s.SetResponse(writer, nil, http.StatusCreated)
}

func (s *Server) SetResponseBadRequest(writer http.ResponseWriter, message string) {
	s.SetResponseError(writer, message, http.StatusBadRequest)
}

func (s *Server) SetResponseInternalError(writer http.ResponseWriter) {
	s.SetResponseError(writer, messageInternalError, http.StatusInternalServerError)
}

func (s *Server) SetResponseError(writer http.ResponseWriter, message string, status int) {
	response := &DefaultResponse{
		Message: message,
		Code:    status,
	}
	s.SetResponse(writer, response, status)
}

func (s *Server) StatusOk() int {
	return http.StatusOK
}

func (s *Server) StatusInternalError() int {
	return http.StatusInternalServerError
}

func (s *Server) StatusBadRequest() int {
	return http.StatusBadRequest
}

func (s *Server) setStatus(writer http.ResponseWriter, status int) {
	writer.WriteHeader(status)
}

func convert(response interface{}) ([]byte, error) {
	body, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return body, nil
}
