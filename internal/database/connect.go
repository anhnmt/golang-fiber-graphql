package database

import (
	"github.com/xdorro/golang-fiber-project/ent"
	"github.com/xdorro/golang-fiber-project/internal/config"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	once *sync.Once
	db   *ent.Client
)

func Connection() *ent.Client {
	if db == nil {
		once = &sync.Once{}

		once.Do(func() {
			if db == nil {
				database := config.Database()
				log.Printf("Connect to [%s] %s", database.Driver, database.Url)

				client, err := ent.Open(database.Driver, database.Url)

				if err != nil {
					log.Fatal(err)
				}

				// Run migration.
				Migration(client)

				db = client
			}
		})
	}

	return db
}

func Close() error {
	return db.Close()
}
