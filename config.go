package artiefact

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// AppConfig application config
type AppConfig struct {
	Environment     string `toml:"environment"`
	ServerPort      uint16 `toml:"server_port"`
	DatabaseName    string `toml:"database_name"`
	DatabaseHost    string `toml:"database_host"`
	DatabasePass    string `toml:"database_pass"`
	DatabaseUser    string `toml:"database_user"`
	DatabasePort    uint16 `toml:"database_port"`
	DatabaseSSLMode string `toml:"database_ssl_mode"`
}

// NewAppConfig create application config
func NewAppConfig(configPath string) (*AppConfig, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	var config AppConfig
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, errors.Wrap(err, "failed to create AppConfig from file")
	}

	env := os.Getenv("ARTIEFACT_ENV")
	if env != "" {
		config.DatabaseHost = os.Getenv("ARTIEFACT_DATABASE_URL")
		config.DatabaseName = os.Getenv("ARTIEFACT_DATABASE_NAME")
		config.DatabasePass = os.Getenv("ARTIEFACT_DATABASE_PASS")
		config.DatabaseUser = os.Getenv("ARTIEFACT_DATABASE_USER")
	}

	return &config, nil
}
