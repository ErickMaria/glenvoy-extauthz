package database

import (
	"context"
	"fmt"
	"github/erickmaria/glooe-envoy-extauthz/internal/entity"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var (
	ctx = context.Background()
)

type MSSql struct {
	Server   string
	Port     int
	Database string
	Password string
	Username string
}

func (mssql *MSSql) Connect() *gorm.DB {

	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		mssql.Username, mssql.Password, mssql.Server, mssql.Port, mssql.Database)

	db, err := gorm.Open("mssql", connectionString)
	if err != nil {
		log.Fatal("Error to create connection: ", err.Error())
	}

	return db

}

func (mssql *MSSql) CreateMigration(db *gorm.DB) {

	createMigration := db.AutoMigrate(
		&entity.Token{},
		&entity.App{},
		&entity.Group{},
		&entity.Domain{},
	).
		Model(&entity.Token{}).AddForeignKey("app_id", "APPS(id)", "CASCADE", "CASCADE").
		// Model(&entity.Token{}).AddForeignKey("group_id", "GROUPS(id)", "CASCADE", "CASCADE").
		Model(&entity.App{}).AddForeignKey("group_id", "GROUPS(id)", "CASCADE", "CASCADE").
		Model(&entity.Group{}).AddForeignKey("domain_id", "DOMAINS(id)", "CASCADE", "CASCADE")

	if err := createMigration.GetErrors(); len(err) > 0 {
		log.Fatalf("Create Migration Errors: %v", err)
	}
}

func (mssql *MSSql) DeleteMigration(db *gorm.DB) {

	db.Model(&entity.Token{}).RemoveForeignKey("app_id", "APPS(id)")
	// db.Model(&entity.Token{}).RemoveForeignKey("group_id", "GROUPS(id)")
	db.Model(&entity.App{}).RemoveForeignKey("group_id", "GROUPS(id)")
	db.Model(&entity.Group{}).RemoveForeignKey("domain_id", "DOMAINS(id)")

	deleteMigration := db.DropTableIfExists(
		&entity.Group{},
		&entity.Domain{},
		&entity.App{},
		&entity.Token{},
	)

	if err := deleteMigration.GetErrors(); len(err) > 0 {
		log.Fatalf("Delete Migration Errors: %v", err)
	}

}

func (mssql *MSSql) Ping(db *gorm.DB) error {
	if err := db.DB().PingContext(ctx); err != nil {
		return err
	}
	return nil
}
