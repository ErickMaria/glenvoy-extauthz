package main

import (
	"flag"
	"github/erickmaria/glenvoy-extauthz/internal/authz"
	"github/erickmaria/glenvoy-extauthz/internal/pkg/config"
	"log"
	"net"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"

	"google.golang.org/grpc"
)

func init() {

	// Parsing Command-Line Flag
	var profile string
	flag.StringVar(&profile, "profile", "development", "applation profile")
	flag.Parse()

	// Initializing applacation Profile
	config.Init("../../../configs", "application.yaml", profile)
}

func main() {
	// create a TCP
	addr := config.AppConfig.HTTP.Host + ":" + config.AppConfig.HTTP.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", lis.Addr())

	grpcServer := grpc.NewServer()

	implAuthServer := &authz.ImplAuthorizationServer{}
	auth.RegisterAuthorizationServer(grpcServer, implAuthServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
