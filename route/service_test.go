package route

import "testing"

func Test(t *testing.T) {
	var routes []*Route

	route1 := &Route{
		Price: 10,
	}
	routes = append(routes, route1)

	route2 := &Route{
		Price: 20,
	}
	routes = append(routes, route2)

	routes = pop(routes)
	if len(routes) > 1 {
		t.Error("Expected only 1 route.")
	} else if routes[0].Price != 10 {
		t.Error("Expected item from stack bottom.")
	}
}
