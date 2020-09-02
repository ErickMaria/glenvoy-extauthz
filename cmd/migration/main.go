package main

import (
	"context"
	"flag"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"
	"github/erickmaria/glooe-envoy-extauthz/internal/database"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"
)

var (
	profile, migration string
	migrate            = database.Migrate{}
	ctx                = context.Background()
)

func init() {

	// Parsing Command-Line Flag
	flag.StringVar(&profile, "profile", "", "get profile allows in configs/profile.yaml")
	flag.StringVar(&migration, "migration", "create", "migration mode: acceped: create, delete")
	flag.Parse()

	// Initializing applacation Profile
	config.Init(profile, ctx)
	logging.Init(config.AppConfig.Glenvoy.App.Name)
	logging.Logger(ctx).Infof("Application profile: %s", config.AppConfig.Profile)
}

func main() {
	conn := database.NewConnection()
	db := conn.Dial(ctx)
	defer db.Close()

	db.LogMode(true)

	if migration == "create" {
		logging.Logger(ctx).Infow("creating migration")
		migrate.Create(ctx, db)
	} else if migration == "delete" {
		logging.Logger(ctx).Infow("deleting migration")
		migrate.Delete(ctx, db)
	} else {
		logging.Logger(ctx).Fatalf("migration value not acceped")
	}

	logging.Logger(ctx).Infow("Finished")
}
