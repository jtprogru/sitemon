package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

// MonitorOpts stores the settings of the request for the requested resource
type MonitorOpts struct {
	Method     string `yaml:"method"`
	Url        string `yaml:"url"`
	Timeout    int    `yaml:"timeout"`
	CheckEvery int    `yaml:"check_every"`
}

// Config contains application settings
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

// BuildConfig creates a configuration structure using a config file
// in .yml format and also overrides base settings with environment variables
func BuildConfig(configPath string) (*Config, error) {
	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// create config structure
	var c Config

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	// If any parameter is passed through environment variables,
	// then the passed value must override the value from the configuration file.
	err = envconfig.Process("", &c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
