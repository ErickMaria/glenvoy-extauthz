package database

import (
	"github/erickmaria/glooe-envoy-extauthz/internal/entity"
	"log"

	"github.com/jinzhu/gorm"
)

type Migrate struct{}

func (mssql *Migrate) Create(db *gorm.DB) {

	create := db.AutoMigrate(

		&entity.Token{},
		&entity.App{},
		&entity.Domain{},
	).
		Model(&entity.Token{}).AddForeignKey("app_id", "apps(id)", "CASCADE", "CASCADE").
		Model(&entity.App{}).AddForeignKey("domain_id", "domains(id)", "CASCADE", "CASCADE")

	if err := create.GetErrors(); len(err) > 0 {
		log.Fatalf("Create Migration Errors: %v", err)
	}
}

func (mssql *Migrate) Delete(db *gorm.DB) {

	delete := db.Model(&entity.Token{}).RemoveForeignKey("app_id", "apps(id)").
		Model(&entity.App{}).RemoveForeignKey("domain_id", "domains(id)").
		DropTableIfExists(
			&entity.Domain{},
			&entity.App{},
			&entity.Token{},
		)

	if err := delete.GetErrors(); len(err) > 0 {
		log.Fatalf("Delete Migration Errors: %v", err)
	}

}
