package main

import (
	"bufio"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
	"strings"
	"travel-route/route"
)

const (
	errMissingArgs             = "The args of best route is missing or invalid."
	errReadingInformedFile     = "Error reading routes from informed file"
	errConvertingInformedPrice = "Error converting the informed price"
	errLoadingRoutesFromFile   = "Error loading routes from informed file"
	errScanningInformedFile    = "Error scanning routes from informed file"
	errEmptyRoutes             = "The informed file is empty"
	errMissingBestRoute        = "The request best route is missing or invalid"
)

func main() {
	bestRoute, err := loadBestRoute()
	if err != nil {
		log.Fatalf(err.Error())
	}

	routes, err := validateAndLoadInputRoute()
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("%v", routes)
	log.Printf("%v", bestRoute)
}

func loadBestRoute() (*route.Route, error) {
	args := os.Args
	if len(args) < 3 || args[0] != "" || args[1] != "" || args[2] != "" {
		return nil, errors.New(errMissingArgs)
	}

	bestRoute, err := route.NewBestRoute(args[0], args[1])
	if bestRoute == nil {
		return nil, errors.New(errMissingBestRoute)
	}

	return bestRoute, err
}

func validateAndLoadInputRoute() ([]*route.Route, error) {
	f, err := os.Open("./resource/input-file.txt")
	defer f.Close()
	if err != nil {
		return nil, errors.New(errReadingInformedFile)
	}

	var result []*route.Route

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		price, err := strconv.ParseFloat(splitLine[2], 64)
		if err != nil {
			return nil, errors.Wrap(err, errConvertingInformedPrice)
		}

		newRoute, err := route.NewRoute(splitLine[0], splitLine[1], price)
		if err != nil {
			return nil, errors.New(errLoadingRoutesFromFile)
		}
		result = append(result, newRoute)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, errScanningInformedFile)
	}

	if len(result) == 0 {
		return nil, errors.New(errEmptyRoutes)
	}

	return result, nil
}
