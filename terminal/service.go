package terminal

import (
	"bufio"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
	"travel-route/route"
)

type Service struct {
	routeRepo *route.Repository
}

const (
	errEmptyFile                     = "Empty file"
	errConvertingInformedPrice       = "Error converting the informed price"
	errReadingRoutesFromFile         = "Error reading routes from informed file"
	errLoadingRoutesFromFile         = "Error loading routes from informed file"
	errScanningInformedFile          = "Error scanning routes from informed file"
	errSaveAllRoutesFromInformedFile = "Error saving all routes from informed file"
	errSavingInformedFile            = "Error saving routes from informed file"
)

func (s *Service) saveRoutesFromFile(filePath string) ([]*route.Route, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, errors.New(errReadingRoutesFromFile)
	}

	routes, err := s.saveRoutes(file)
	if err != nil {
		return nil, errors.New(errSavingInformedFile)
	} else if len(routes) == 0 {
		return nil, errors.New(errEmptyRoutes)
	}

	return routes, nil
}

func (s *Service) saveRoutes(file *os.File) ([]*route.Route, error) {
	var result []*route.Route

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")

		price, err := strconv.ParseFloat(splitLine[2], 64)
		if err != nil {
			return nil, errors.Wrap(err, errConvertingInformedPrice)
		}

		result, err = s.routeRepo.Save(splitLine[0], splitLine[1], price)
		if err != nil {
			return nil, errors.New(errSaveAllRoutesFromInformedFile)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, errScanningInformedFile)
	}

	if len(result) == 0 {
		return nil, errors.New(errEmptyFile)
	}

	return result, nil
}
