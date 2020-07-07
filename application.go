package main

import (
	"flag"
	"log"
	"travel-route/http"
	"travel-route/route"
	"travel-route/terminal"
)

func main() {
	db := &route.Db{}
	db.ConfigDb()

	routeService := &route.Service{
		Repo: &route.Repository{
			Db: db,
		},
	}
	routeController := route.NewController(routeService)
	cmdTerminal := terminal.NewTerminal(routeService)

	var pathFile string
	flag.StringVar(&pathFile, "file", "", "The file flag is mandatory")
	flag.Parse()

	err := cmdTerminal.LoadRoutesFromFile(pathFile)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// terminal command listener
	go cmdTerminal.TerminalListener()

	// Init http server
	httpServer := http.NewServer()
	httpServer.RegisterEndpoint("GET", "/routes", routeController.GetRoutes)
	httpServer.RegisterEndpoint("POST", "/route", routeController.PostRoute)
	httpServer.RegisterEndpoint("*", "/", httpServer.DefaultHandler)
	httpServer.Start()
}
