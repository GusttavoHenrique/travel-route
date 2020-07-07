package terminal

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGetInputRoute(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"nil input", ""},
		{"informed invalid input", "SAO"},
		{"informed destination", " - NAT"},
		{"informed origin", "SAO - "},
		{"informed origin and destination", "SAO - NAT"},
		{"informed hyphen", " - "},
		{"informed space", "  "},
		{"informed any thing", " any thing "},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			content := []byte(test.input)
			tmpfile, err := ioutil.TempFile("", "")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write(content); err != nil {
				log.Fatal(err)
			}

			if _, err := tmpfile.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			oldStdin := os.Stdin
			os.Stdin = tmpfile
			defer func() { os.Stdin = oldStdin }()

			terminal := &Terminal{}
			inputRoute, err := terminal.GetInputRoute()
			if err == nil && inputRoute == nil {
				t.Fatal("Expected error.")
			}

			if inputRoute != nil {
				if inputRoute.FinalPoint == nil {
					t.Error("Expected initial point in new route instance.")
				} else if inputRoute.FinalPoint == nil {
					t.Error("Expected final point in new route instance.")
				}
			}

			if err := tmpfile.Close(); err != nil {
				log.Fatal(err)
			}
		})
	}
}

func TestLoadBestRoute(t *testing.T) {
	file := "./resource/input-routes.csv"

	tests := []struct {
		name     string
		input    string
		filePath string
	}{
		{"empty file", "", file},
		{"one route", "SAO,NAT,10", file},
		{"duplicated routes", "SAO,NAT,10\nSAO,NAT,10", file},
		{"duplicated routes", "SAO,NAT,10\nSAO,NAT,10\nSAO,NAT,10", file},
		{"same route with distinct prices", "SAO,NAT,10\nSAO,NAT,20", file},
		{"distinct routes", "NAT,SAO,10\nNAT,SAO,20", file},
		{"distinct routes without file path", "NAT,SAO,10\nNAT,SAO,20", ""},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			content := []byte(test.input)
			tmpfile, err := ioutil.TempFile("", "")
			if err != nil {
				log.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write(content); err != nil {
				log.Fatal(err)
			}

			if _, err := tmpfile.Seek(0, 0); err != nil {
				log.Fatal(err)
			}

			oldStdin := os.Stdin
			os.Stdin = tmpfile
			defer func() { os.Stdin = oldStdin }()

			terminal := &Terminal{}
			err = terminal.LoadRoutesFromFile(test.filePath)
			if err == nil {
				t.Fatal("Expected error.")
			}

			if err := tmpfile.Close(); err != nil {
				log.Fatal(err)
			}
		})
	}
}
