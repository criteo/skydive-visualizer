package main

import (
	"network/skydive-visualizer-go/source/chef"
	"network/skydive-visualizer-go/source/ipam"
	"os"
	"reflect"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	envPrefix = "SKYDIVEVIZ_"
)

type Config struct {
	Server *struct {
		Listen string `yaml:"listen,omitempty" env:"LISTEN"`
	} `yaml:"server,omitempty"`

	Skydive *struct {
		URL string `yaml:"url,omitempty" env:"SKYDIVE_URL"`
	} `yaml:"skydive,omitempty"`

	Chef *chef.Config `yaml:"chef,omitempty"`

	IPAM *ipam.Config `yaml:"ipam,omitempty"`
}

func GetConfig(path string) (Config, error) {
	cfg := Config{}

	f, err := os.Open(path)
	if err != nil {
		return cfg, errors.Wrap(err, "error reading config")
	}

	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return cfg, errors.Wrap(err, "error decoding config")
	}

	cfge := doEnv(&cfg).(*Config)

	return *cfge, nil
}

func doEnv(in interface{}) interface{} {
	s := reflect.ValueOf(in).Elem()
	t := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ft := t.Field(i)
		switch f.Kind() {
		case reflect.Ptr:
			v := doEnv(f.Interface())
			f.Set(reflect.ValueOf(v))

		default:
			envKey := ft.Tag.Get("env")
			if envKey == "" {
				continue
			}

			v := os.Getenv(envPrefix + envKey)
			if v != "" {
				f.SetString(v)
			}
		}
	}

	return in
}
