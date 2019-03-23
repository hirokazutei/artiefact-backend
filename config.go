package artiefact

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	c "github.com/hirokazu/artiefact-backend/constants"
	"github.com/pkg/errors"
)

// AppConfig application config
type AppConfig struct {
	Environment     string `toml:"environment"`
	ServerPort      uint16 `toml:"server_port"`
	DatabaseName    string
	DatabaseHost    string
	DatabasePass    string
	DatabaseUser    string
	DatabasePort    uint16 `toml:"database_port"`
	DatabaseSSLMode string `toml:"database_ssl_mode"`
	TokenSecret     string
	PasswordPepper  string
}

// NewAppConfig create application config
func NewAppConfig(configPath string) (*AppConfig, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(c.ErrorOpening, "file"))
	}
	defer file.Close()

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(c.ErrorReading, "file"))
	}

	var config AppConfig
	if err := toml.Unmarshal(buffer, &config); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf(c.ErrorCreatingFrom, "AppConfig", "file"))
	}

	env := os.Getenv("ARTIEFACT_ENV")
	if env != "" {
		config.DatabaseHost = os.Getenv("ARTIEFACT_DATABASE_URL")
		config.DatabaseName = os.Getenv("ARTIEFACT_DATABASE_NAME")
		config.DatabasePass = os.Getenv("ARTIEFACT_DATABASE_PASS")
		config.DatabaseUser = os.Getenv("ARTIEFACT_DATABASE_USER")
		config.PasswordPepper = os.Getenv("ARTIEFACT_PASSWORD_PEPPER")
	}

	return &config, nil
}
