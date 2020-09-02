package config

import (
	"bytes"
	"context"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
	"github/erickmaria/glooe-envoy-extauthz/internal/pkg/logging"
)

// Profile for configuration files
type Profile struct {
	Path 	string 	 `yaml:"path"`
	File 	string   `yaml:"file"`
	Suffixs []string `yaml:"suffix"`
	Parent string	 `yaml:"parent"`	
	Default bool 	 `yaml:"default"`
}

var ProfileConfig *Profile

func Load(profileFlag string, ctx context.Context) (error, string) {

	data, err := ioutil.ReadFile("configs/profile.yaml")
	if err != nil {
		logging.Logger(ctx).Fatalf("cannot found profile file %s", err)
	}	
	
	var opts map[string]Profile
	bdata := bytes.NewBuffer(data).Bytes()

	if err = yaml.Unmarshal(bdata, &opts); err != nil {
		logging.Logger(ctx).Fatalf("[ProfileConfig] Error reading profile file, %s", err)
	}

	if profileFlag == "" {
		for opt := range(opts){
			profileOpt := opts[opt]
			if profileOpt.Default == true {
				profileFlag = opt
				break
			}
		}
	}

	pconfig := opts[profileFlag]

	if pconfig.File == "" {
		logging.Logger(ctx).Fatalf("profile type not exists")
	}

	ProfileConfig = &pconfig

	return nil, profileFlag
}