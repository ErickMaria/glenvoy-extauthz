package main

import (
	"flag"
	"github/erickmaria/glooe-envoy-extauthz/internal/database"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/config"
	"log"
)

var (
	conn               = database.Connection{}
	migrate            = database.Migrate{}
	profile, migration string
)

func init() {

	// Parsing Command-Line Flag
	flag.StringVar(&profile, "profile", "development", "applation profile acceped: development, production, test")
	flag.StringVar(&migration, "migration", "create", "migration mode: acceped: create, delete")
	flag.Parse()

	// Initializing applacation Profile
	config.Init("../../../configs", "application.yaml", profile)

	// Dababase Configuration
	conn = database.Connection{
		Dialect: config.AppConfig.Datasource.Dialect,
		SSLMode: "disable",
		Datasource: database.Datasource{
			Host:     config.AppConfig.Datasource.Host,
			Port:     config.AppConfig.Datasource.Port,
			Database: config.AppConfig.Datasource.Database,
			Username: config.AppConfig.Datasource.Username,
			Password: config.AppConfig.Datasource.Password,
		},
	}

}

func main() {
	log.Println("Application profile:", config.AppConfig.Profile)

	db := conn.Get()
	defer db.Close()

	db.LogMode(true)

	if migration == "create" {
		log.Println("Creating migration")
		migrate.Create(db)
	} else if migration == "delete" {
		log.Println("Deleting migration")
		migrate.Delete(db)
	} else {
		log.Fatalln("migration value not acceped")
	}

	log.Println("Finished")
}
