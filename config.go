package daytimer

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// Global configuration for the application
var config *Config

// LoadConfig returns the global configuration, instantiating it if needed.
// This function will also create the configuration stub file if the config
// is not at the specified location on disk.
func LoadConfig() (*Config, error) {
	if config == nil {
		c := new(Config)

		if err := c.checkConfig(); err != nil {
			return nil, err
		}

		if err := c.Load(); err != nil {
			return nil, err
		}

		config = c
	}

	return config, nil
}

// Config stored on disk and loaded from the configuration directory.
type Config struct {
	Editor string      `json:"editor"`
	Email  *SMTPConfig `json:"email"`
}

// Load the configuration from the internal path.
func (c *Config) Load() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, c)
}

// Dump the configuration to the internal path.
func (c *Config) Dump() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}

// Edit the configuration file.
func (c *Config) Edit() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	if err := EditFile(path); err != nil {
		return fmt.Errorf("%s: please edit %s directly", err, path)
	}

	return nil
}

// String returns the YAML representation of the configuration.
func (c *Config) String() string {
	data, err := yaml.Marshal(c)
	if err != nil {
		return "invalid configuration"
	}
	return string(data)
}

// Get the path to the configuration file
func (c *Config) configPath() (string, error) {
	confDir, err := configDirectory()
	if err != nil {
		return "", err
	}

	return filepath.Join(confDir, "config.yml"), nil
}

// Check if the config file exists, if it doesn't create the stub config from
// the assets template (e.g. with comments) for the user.
func (c *Config) checkConfig() error {
	path, err := c.configPath()
	if err != nil {
		return err
	}

	exists, err := pathExists(path)
	if err != nil {
		return err
	}

	if !exists {
		conf, err := Asset("templates/config.yml")
		if err != nil {
			return err
		}

		return ioutil.WriteFile(path, conf, 0644)
	}

	return nil
}
