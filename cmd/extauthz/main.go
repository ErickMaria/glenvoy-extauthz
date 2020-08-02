package main

import (
	"flag"
	"github/erickmaria/glenvoy-extauthz/internal/pkg/config"
)

func init() {

	// Parsing Command-Line Flag
	var profile string
	flag.StringVar(&profile, "profile", "development", "applation profile")
	flag.Parse()

	// Initializing applacation Profile
	config.Init("../../../configs", "application.yaml", profile)
}

func main() {}