package terminal

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestLoadBestRoute(t *testing.T) {
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
			tmpfile, err := ioutil.TempFile("", "example")
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
			defer func() { os.Stdin = oldStdin }()

			os.Stdin = tmpfile

			inputRoute, err := GetInputRoute()
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
