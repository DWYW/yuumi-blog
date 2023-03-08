package conf

import (
	"gopkg.in/yaml.v3"
)

type ConfigSource struct {
	Type   string  `yaml:"type" json:"type"`
	Source *[]byte `yaml:"source" json:"source"`
}

func Parse(sources ...ConfigSource) *Config {
	conf := Config{}

	for _, item := range sources {
		if item.Type == "yaml" {
			yaml.Unmarshal(*item.Source, &conf)
		}
	}

	return &conf
}
