package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ctx = context.Background()
)

type Connection struct {
	Datasource Datasource
	SSLMode    string
	Dialect    string
}

func (conn *Connection) Get() *gorm.DB {

	db, err := gorm.Open(conn.Dialect, conn.getDialectConnection())
	if err != nil {
		log.Fatal("Error to create connection: ", err.Error())
	}

	log.Println("Database Connection Established")

	return db

}

func (conn *Connection) Ping(db *gorm.DB) error {
	if err := db.DB().PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func (conn *Connection) getDialectConnection() string {
	if conn.Dialect == "" {
		log.Fatal("Dialect cannot is empty")
	}
	dialect := strings.ToLower(conn.Dialect)
	switch dialect {
	case "postgres":
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			conn.Datasource.Host, conn.Datasource.Port, conn.Datasource.Database, conn.Datasource.Username, conn.Datasource.Password, conn.SSLMode)
	default:
		log.Fatal("Dialect not permitted")
	}

	return ""
}
