package database

import (
	"context"
	"github.com/xdorro/golang-fiber-base-project/ent"
	"github.com/xdorro/golang-fiber-base-project/ent/migrate"
	"log"
)

func Migration(client *ent.Client) {
	// Run migration.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
