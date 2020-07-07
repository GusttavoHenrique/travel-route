package route

import (
	"sync"
)

var lock = &sync.Mutex{}

type Db struct {
	filePath    string
	tableRoutes []*Route
}

// ConfigDb init the database
func (db *Db) ConfigDb() {
	lock.Lock()
	defer lock.Unlock()

	db.filePath = ""

	if db.tableRoutes == nil {
		db.tableRoutes = make([]*Route, 0)
	}
}

func (db *Db) saveFileRoute(filePath string) {
	db.filePath = filePath
}

func (db *Db) saveRoute(item *Route) {
	db.tableRoutes = append(db.tableRoutes, item)
}

func (db *Db) findRoutes() []*Route {
	return db.tableRoutes
}
