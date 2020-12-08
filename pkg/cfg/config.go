package cfg

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
)

const FilePath = "./config.yaml"

type (
	// Cfg defines the properties of the application configuration
	Cfg struct {
		APIServer     Server        `yaml:"api-server"`
		MessageBroker MessageBroker `yaml:"message-broker"`
		Logger        Logger        `yaml:"logger"`
		Storage       Storage       `yaml:"storage"`
	}

	// Broker defines Broker section of the API server configuration
	MessageBroker struct {
		Type           string   `yaml:"type"`
		ConnectionAddr string   `yaml:"connection-address"`
		ServiceName    string   `yaml:"service-name"`
		MaxRetries     int      `yaml:"max-retries"`
		RetryDelayMs   int      `yaml:"retry-delay-ms"`
		RequestTimeout int      `yaml:"request-timeout"`
		Name           string   `yaml:"name"`
		Pass           string   `yaml:"pass"`
		Subjects       Subjects `yaml:"subjects"`
	}

	Subjects struct {
		Requests map[string]string
	}

	// Server defines API server configuration
	Server struct {
		HTTP HTTP `yaml:"http"`
	}

	// HTTP defines HTTP section of the API server configuration
	HTTP struct {
		AuthRequired bool   `yaml:"auth-required"`
		ListenAddr   string `yaml:"listen-address"`
	}

	// Logger defines logger section of the API server configuration
	Logger struct {
		OutputFilePath        string `yaml:"output-file-path"`
		DebugLevel            string `yaml:"debug-level"`
		LogFormat             string `yaml:"log-format"`
		IncludeCallerMethod   bool   `yaml:"include-caller-method"`
		RequestOutputFilePath string `yaml:"requests-log-output-file-path"`
	}

	// Storage defines database engines configuration
	Storage struct {
		Postgres Postgres `yaml:"postgres"`
	}

	// Postgres defines PostgreSQL specific configuration
	Postgres struct {
		ConnectionString   string `yaml:"connection-string"`
		Driver             string `yaml:"driver"`
		MaxRetries         int    `yaml:"max-retries"`
		RetryDelay         int    `yaml:"retry-delay"`
		QueryTimeout       int    `yaml:"query-timeout"`
		AutoMigrate        bool   `yaml:"auto-migrate"`
		MigrationDirectory string `yaml:"migration-directory"`
		MigrationDirection string `yaml:"migration-direction"`
	}
)

// Init loads and validates all configuration data
func Init(filename string) (*Cfg, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "read config file error")
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open cfg file")
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read cfg file content")
	}

	var cfg Cfg
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cfg")
	}

	return &cfg, nil
}
