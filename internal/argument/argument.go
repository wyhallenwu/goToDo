package argument

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Fileconfig File `yaml:"fileconfig"`
}

type File struct {
	Dir      string `yaml:"dir"`
	TodoFile string `yaml:"todofile"`
	DoneFile string `yaml:"donefile"`
}

func ReadConfig(path string) (*Config, error) {
	conf := &Config{}
	if file, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(file).Decode(conf)
	}
	return conf, nil
}
