package main

import (
	"context"
	"flag"
	"fmt"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"
	"github/erickmaria/glooe-envoy-extauthz/internal/database"
	"github/erickmaria/glooe-envoy-extauthz/internal/entity"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"
	"github/erickmaria/glooe-envoy-extauthz/internal/types"
	"log"
)

var (
	migrate = database.Migrate{}
	ctx     = context.Background()
)

func init() {

	// Parsing Command-Line Flag
	var profile string
	flag.StringVar(&profile, "profile", "development", "applation profile acceped: development, production, test")
	flag.Parse()

	// Initializing applacation Profile
	config.Init("../../configs", "application.yaml", profile)
	logging.Logger(ctx).Infof("Application profile: %s", config.AppConfig.Profile)

}

func main() {

	conn := database.NewConnection()
	db := conn.Dial(ctx)
	defer db.Close()

	domain := entity.Domain{}
	app := entity.App{}
	token := entity.Token{}

	db.Find(&domain, entity.Domain{Url: "domain1.test.com"})
	if domain.Url == "" {
		log.Printf("Domain %s not exist on database\n", domain.Name)
		return
	}

	db.Find(&app, entity.App{Code: "apptest1", DomainID: domain.ID})
	if app.Code == "" {
		log.Printf("Not Fould Code on App with Domain %s\n", domain.Name)
		return
	}
	if app.Status == types.REVOKED || app.Status == types.DEACTIVATE {
		log.Printf("App %s on database\n", app.Status)
		return
	}

	db.Find(&token, entity.Token{Code: "tokentest4", AppID: app.ID})
	if token.Code == "" {
		log.Printf("Not Fould Code on Token with App %s\n", app.Name)
		return
	}
	if token.Status == types.REVOKED || token.Status == types.DEACTIVATE {
		log.Printf("Token %s on database\n", token.Status)
		return
	}

	fmt.Println("result: ", token.Code)
}

// func main() {
// 	// create a TCP
// 	addr := config.AppConfig.HTTP.Host + ":" + config.AppConfig.HTTP.Port
// 	lis, err := net.Listen("tcp", addr)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	log.Printf("listening on %s", lis.Addr())

// 	grpcServer := grpc.NewServer()

// 	implAuthServer := &authz.ImplAuthorizationServer{}
// 	auth.RegisterAuthorizationServer(grpcServer, implAuthServer)

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("Failed to start server: %v", err)
// 	}
// }
