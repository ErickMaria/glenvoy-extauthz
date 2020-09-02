package config

import (
	"context"
	"strings"
	"os"

	"github.com/spf13/viper"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"
)

const (
	ymlPrefixKey = "glenvoy"
)

// Applcation for configuration files
type Application struct {
	App struct {
		Name        string `yaml:"name"`
		Environment string `yaml:"environment"`
		Version     string `yaml:"version"`
	} `yaml:"app"`
	HTTP struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"http"`
	Datasource struct {
		Dialect  string `yaml:"dialect"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	} `yaml:"datasource"`
	Redis struct {
		Host        string `yaml:"host"`
		Port        string `yaml:"port"`
		Username    string `yaml:"username"`
		Password    string `yaml:"password"`
		DB          string `yaml:"db"`
		DialTimeout string `yaml:"dialtimeout"`
	} `yaml:"redis"`
	AppKeys struct {
		Headers [2]string `yaml:"headers"`
	} `yaml:"appkeys"`

	Profile string
}

// AppConfig is the configs for the whole application
var AppConfig *Application

// Init is using to initialize the configs
func Init(profileFlag string, ctx context.Context) error {

	// Load application profile
	_, currentProfile := Load(profileFlag, ctx)
	applicationFile := findConfigFile(ProfileConfig, ctx)


	viperSetup := viper.GetViper()
	viper.AddConfigPath(ProfileConfig.Path)
	viperSetup.SetConfigName(applicationFile)
	viperSetup.SetConfigType("yaml")
	viperSetup.AllowEmptyEnv(true)
	viperSetup.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viperSetup.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logging.Logger(ctx).Fatalf("[VIPER] Error reading configuration file, %s", err)
	}

	var opt Application
	
	if err := viperSetup.UnmarshalKey(ymlPrefixKey, &opt); err != nil {
		logging.Logger(ctx).Fatalf("[AppCofing] Error reading configuration files, %s", err)
	}

	opt.Profile = currentProfile 
	AppConfig = &opt

	return nil
}

func findConfigFile(profile *Profile, ctx context.Context) string {
	pattern := "#{suffix}#"
	var configfile string
	var configfileList []string 
	var findApplicationConfig bool = false 

	for _, suffix := range(profile.Suffixs) {
		
		if suffix != "" {
			suffix = "-"+suffix
		}

		filename := strings.Replace(profile.File, pattern, suffix, strings.Count(pattern, "")-1)
		configfile = profile.Path + filename
		configfileList = append(configfileList, configfile)
		if _, err := os.Stat(configfile); os.IsNotExist(err) {
			continue
		}

		findApplicationConfig = true
		break
	}

	if !findApplicationConfig {
		logging.Logger(ctx).Fatalf("not found applcation configs: %s", configfileList)
	}

	return configfile
}
