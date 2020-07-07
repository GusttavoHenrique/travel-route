package route

import (
	"sync"
)

var lock = &sync.Mutex{}

type Db struct {
	tableRoutes []*Route
}

func (db *Db) ConfigDb() {
	lock.Lock()
	defer lock.Unlock()

	if db.tableRoutes == nil {
		db.tableRoutes = make([]*Route, 0)
	}
}

func (db *Db) SaveRoute(item *Route) {
	db.tableRoutes = append(db.tableRoutes, item)
}

func (db *Db) FindRoutes() []*Route {
	return db.tableRoutes
}
