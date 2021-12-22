package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

type MonitorOpts struct {
	Method     string `yaml:"method"`
	Url        string `yaml:"url"`
	Timeout    int    `yaml:"timeout"`
	CheckEvery int    `yaml:"check_every"`
}

type Config struct {
	Log struct {
		Level  string `yaml:"level" envconfig:"SITEMON_LOGLEVEL"`
		Format string `yaml:"format"`
	} `yaml:"log"`
	Sentry struct {
		Dsn string `yaml:"dsn" envconfig:"SITEMON_SENTRYDSN"`
	} `yaml:"sentry"`
	Telegram struct {
		Token string `yaml:"token" envconfig:"SITEMON_TGTOKEN"`
		Chat  string `yaml:"chat" envconfig:"SITEMON_TGCHAT"`
	} `yaml:"telegram"`
	Monitors struct {
		JTProg  MonitorOpts `yaml:"jtprog"`
		HTTPBin MonitorOpts `yaml:"httpbin"`
	} `yaml:"monitors"`
}

func BuildConfig(configPath string) (*Config, error) {
	// create config structure
	var c Config
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		return nil, err
	}
	err = envconfig.Process("", &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
