package database

import (
	"context"
	"github.com/xdorro/golang-fiber-project/ent"
	"github.com/xdorro/golang-fiber-project/internal/config"
	"log"
	"sync"
)

const (
	mariadb  string = "mariadb"
	mysql    string = "mysql"
	postgres string = "postgres"
)

var (
	once *sync.Once
	db   *ent.Client
)

type DBConnection struct {
	Driver string
	Url    string
	Ctx    context.Context
}

func Connection() *ent.Client {
	if db == nil {
		once = &sync.Once{}

		once.Do(func() {
			if db == nil {
				database := config.Database()
				log.Printf("Connect to [%s] %s", database.Driver, database.Url)

				connect := &DBConnection{
					Driver: database.Driver,
					Url:    database.Url,
					Ctx:    context.Background(),
				}

				db = connect.Handler()
			}
		})
	}

	return db
}

func Close() error {
	return db.Close()
}

func (conn *DBConnection) Handler() *ent.Client {
	switch conn.Driver {

	case mysql, mariadb:
		return conn.Mysql()

	case postgres:
		return conn.Postgres()
	}

	log.Fatalf("Can't find database driver")
	return nil
}
