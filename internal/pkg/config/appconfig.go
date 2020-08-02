package config

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gobuffalo/packr"
	yaml "gopkg.in/yaml.v2"
)

// Option for configurations
type Option struct {
	App struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
		Version     string `yaml:"version"`
	} `yaml:"app"`
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Redis struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		DB          string `yaml:"db"`
		DialTimeout string `yaml:"dialtimeout"`
	} `yaml:"redis"`

	Profile string
	// OperatorSet map[string]bool
}

// AppConfig is the configs for the whole application
var AppConfig *Option

// Init is using to initialize the configs
func Init(path, file, profile string) error {

	box := packr.NewBox(path)
	s, err := box.Find(file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var options map[string]Option
	data := bytes.NewBuffer(s).Bytes()

	if err = yaml.Unmarshal(data, &options); err != nil {
		log.Fatalf("[AppCofing] Error reading configuration files, %s", err)
	}

	opt := options[profile]
	opt.Profile = profile
	AppConfig = &opt
	return nil
}
