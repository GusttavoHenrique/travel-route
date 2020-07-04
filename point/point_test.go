package point

import "testing"

const name = "SAO"

func TestNewPoint(t *testing.T) {
	newPoint, err := NewPoint(name)
	validateInstancePoint(newPoint, err, t)

	if newPoint.Name != name {
		t.Errorf("Error actual name = %v, and expected name = %v.", newPoint.Name, name)
	} else if newPoint.estimate <= 0 {
		t.Errorf("Unexpected negative estimate.")
	}
}

func TestNewPointWithoutName(t *testing.T) {
	newPoint, err := NewPoint("")
	validateError(newPoint, err, t)

	expected := errMissingName
	actual := err.Error()
	if actual != expected {
		t.Errorf("Error actual = %v, and expected = %v.", actual, expected)
	}
}

func validateError(route *Point, err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error.")
	} else if route != nil {
		t.Error("Unexpected new point instance.")
	}
}

func validateInstancePoint(route *Point, err error, t *testing.T) {
	if err != nil {
		t.Error("Unexpected error.")
	} else if route == nil {
		t.Error("Expected new point instance.")
	}
}
