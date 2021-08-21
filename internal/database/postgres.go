package database

import (
	"github.com/xdorro/golang-fiber-project/ent"
	"github.com/xdorro/golang-fiber-project/ent/migrate"
	"log"

	_ "github.com/lib/pq"
)

func (conn *DBConnection) Postgres() *ent.Client {
	client, err := ent.Open("postgres", conn.Url)

	if err != nil {
		log.Fatal(err)
	}

	// Run migration.
	if err = client.Schema.Create(
		conn.Ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
