package main

import (
	"context"
	"flag"
	"net"

	"github/erickmaria/glooe-envoy-extauthz/internal/authz"
	"github/erickmaria/glooe-envoy-extauthz/internal/config"
	"github/erickmaria/glooe-envoy-extauthz/internal/database"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	ctx = context.Background()
)

func init() {

	// Parsing Command-Line Flag
	var profile string
	flag.StringVar(&profile, "profile", "", "get profile allows in configs/profile.yaml")
	flag.Parse()

	// Initializing applacation Profile
	config.Init(profile, ctx)
	logging.Init(config.AppConfig.Glenvoy.App.Name)

	logging.Logger(ctx).Infof("loading Application profile: %s", config.AppConfig.Profile)

}

func main() {

	conn := database.NewConnection()
	db := conn.Dial(ctx)
	defer db.Close()

	// create a TCP server
	addr := config.AppConfig.Glenvoy.HTTP.Host + ":" + config.AppConfig.Glenvoy.HTTP.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logging.Logger(ctx).Fatalf("failed to listen: %v", err)
	}
	logging.Logger(ctx).Infof("listening on %s", lis.Addr())
	grpcServer := grpc.NewServer()

	implAuthServer := &authz.ImplAuthorizationServer{
		DB: db,
	}
	auth.RegisterAuthorizationServer(grpcServer, implAuthServer)
	
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		logging.Logger(ctx).Fatalf("failed to start server: %v", err)
	}
}
