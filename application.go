package main

import (
	"fmt"
	"log"
	"travel-route/terminal"
)

func main() {
	//err := terminal.LoadRoutesFromFile()
	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	terminalCommandsListener()

	//server := &http.Server {}
	//server.Config()
}

func terminalCommandsListener() {
	for {
		bestRoute, err := terminal.GetInputRoute()
		if err != nil {
			log.Fatalf(err.Error())
		}

		bestRouteStr, err := bestRoute.GetBestRouteStr()
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("best route: %s > %f\n", bestRouteStr, bestRoute.Price)
		}
	}
}
