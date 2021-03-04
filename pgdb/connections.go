package pgdb

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/go-pg/pg/v9"
)

func (c *connection) syncAllConnections() {
	db := getSysDB()
	var apps []Application
	err := db.Model(&apps).Select()
	if err != nil {
		log.Error("cannot get list of connections")
	}

	for _, app := range apps {
		if !app.IsBlocked {
			_, _ = c.getOrAddConnection(app.AppName)
		} else {
			c.removeConnection(app.AppName)
		}
	}
}

// DB purpose is global access to database
// Should be initialized in main function
type connection struct {
	DB    map[string]*pg.DB
	mutex sync.RWMutex
}

var instance *connection
var once sync.Once

// GetSysDB return authentication database
func getSysDB() *pg.DB {
	cn, err := getDBConnection("auth")
	if err != nil {
		panic("System database not available")
	}
	return cn
}

func getInstance() *connection {
	once.Do(func() {
		instance = &connection{}
		instance.DB = make(map[string]*pg.DB)
	})

	return instance
}

// getDBConnection return new connection
// if connection not available add it first
func getDBConnection(name string) (*pg.DB, error) {
	c := getInstance()
	cn, err := c.getOrAddConnection(name)
	if err != nil {
		return nil, err
	}
	return cn, nil
}

// GetConnectionList return copy of connections
func GetConnectionList() map[string]*pg.DB {
	return getInstance().DB
}

// addConnection adds new connection
func (c *connection) getOrAddConnection(name string) (*pg.DB, error) {
	con, ok := c.DB[name]

	if !ok {
		c.mutex.Lock()
		con = pg.Connect(
			&pg.Options{
				User:     "postgres",
				Database: name,
				Password: os.Getenv("DATABASE_ROOT_PASS"),
			})
		c.DB[name] = con
		c.mutex.Unlock()
	}
	return con, nil
}

func (c *connection) removeConnection(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.DB, name)
}
