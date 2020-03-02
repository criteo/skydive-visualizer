package ipam

import "network/skydive-visualizer-go/graph"

type Config struct {
	Host     string `yaml:"host,omitempty" env:"IPAM_HOST"`
	Port     int    `yaml:"port,omitempty" env:"IPAM_PORT"`
	Database string `yaml:"database,omitempty" env:"IPAM_DATABASE"`
	User     string `yaml:"user,omitempty" env:"IPAM_USER"`
	Password string `yaml:"password,omitempty" env:"IPAM_PASSWORD"`
	Mapping  []struct {
		Match string          `yaml:"match,omitempty"`
		Value string          `yaml:"value,omitempty"`
		ID    graph.Attribute `yaml:"id,omitempty"`
		Name  string          `yaml:"name,omitempty"`
	} `yaml:"mapping,omitempty"`
}
