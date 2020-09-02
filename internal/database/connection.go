package database

import (
	"context"
	"fmt"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Connection struct {
	Datasource Datasource
	SSLMode    string
	Dialect    string
}

func NewConnection() Connection {
	return Connection{
		Dialect: config.AppConfig.Glenvoy.Datasource.Dialect,
		SSLMode: "disable",
		Datasource: Datasource{
			Host:     config.AppConfig.Glenvoy.Datasource.Host,
			Port:     config.AppConfig.Glenvoy.Datasource.Port,
			Database: config.AppConfig.Glenvoy.Datasource.Database,
			Username: config.AppConfig.Glenvoy.Datasource.Username,
			Password: config.AppConfig.Glenvoy.Datasource.Password,
		},
	}

}

func (conn Connection) Dial(ctx context.Context) *gorm.DB {

	db, err := gorm.Open(conn.Dialect, conn.getDialectConnection(ctx))
	if err != nil {
		logging.Logger(ctx).Fatalf("error to create connection: %v", err.Error())
	}

	logging.Logger(ctx).Infow("database connection established")

	return db

}

func (conn *Connection) Ping(ctx context.Context, db *gorm.DB) error {
	if err := db.DB().PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func (conn *Connection) getDialectConnection(ctx context.Context) string {

	// conn := conn.newConnection()

	if conn.Dialect == "" {
		logging.Logger(ctx).Fatal("dialect cannot is empty")
	}
	dialect := strings.ToLower(conn.Dialect)
	switch dialect {
	case "postgres":
		return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
			conn.Datasource.Host, conn.Datasource.Port, conn.Datasource.Database, conn.Datasource.Username, conn.Datasource.Password, conn.SSLMode)
	default:
		logging.Logger(ctx).Fatal("dialect not permitted")
	}

	return ""
}
