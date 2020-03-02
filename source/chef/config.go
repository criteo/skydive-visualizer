package chef

import "network/skydive-visualizer-go/graph"

type Config struct {
	Servers      []string `yaml:"servers,omitempty"`
	User         string   `yaml:"user,omitempty"`
	Key          string   `yaml:"key,omitempty" env:"CHEF_KEY"`
	AttrsMapping []struct {
		Key  string          `yaml:"key,omitempty"`
		ID   graph.Attribute `yaml:"id,omitempty"`
		Name string          `yaml:"name,omitempty"`
	} `yaml:"attrs_mapping,omitempty"`
}
