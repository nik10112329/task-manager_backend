package config

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"learning-project/pkg/logging"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env        string           `yaml:"env" env:"ENV" env-default:"local"`
	Listen     ListenConfig     `yaml:"listen"`
	PostgreSQL PostgreSQLConfig `yaml:"postgresql"`
}

type ListenConfig struct {
	Type            string        `yaml:"type" env-default:"tcp"`
	BindIP          string        `env:"BIND_IP" env-default:"127.0.0.1"`
	Port            string        `env:"PORT" env-default:"10000"`
	ReadTimeout     time.Duration `yaml:"read_timeout" env-default:"10s"`
	WriteTimeout    time.Duration `yaml:"write_timeout" env-default:"10s"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env-default:"15s"`
}

type PostgreSQLConfig struct {
	// Из окружения — секреты и то, что меняется между машинами
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
	Username string `env:"POSTGRES_USER" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Database string `env:"POSTGRES_DB" env-required:"true"`
	SSLMode  string `env:"POSTGRES_SSLMODE" env-default:"require"`

	// Из yaml — одинаково у всех
	MaxOpenConns    int           `yaml:"max_open_conns" env-default:"10"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env-default:"5"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env-default:"30m"`
	ConnMaxIdleTime time.Duration `yaml:"conn_max_idle_time" env-default:"2m"`
	ConnectTimeout  time.Duration `yaml:"connect_timeout" env-default:"15s"`
	QueryTimeout    time.Duration `yaml:"query_timeout" env-default:"5s"`
	MigrationsPath  string        `yaml:"migrations_path" env-default:"./migrations"`
}

// DSN собирает строку подключения. (Data Source Name)
// url.UserPassword экранирует спецсимволы — Neon генерирует
// пароли с #, /, @, и склейка через Sprintf ломает подключение.
func (c PostgreSQLConfig) DSN() string {
	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.Username, c.Password),
		Host:   fmt.Sprintf("%s:%s", c.Host, c.Port),
		Path:   c.Database,
	}
	q := u.Query()
	q.Set("sslmode", c.SSLMode)
	u.RawQuery = q.Encode()
	return u.String()
}

// Address — адрес для http.Server.
func (c ListenConfig) Address() string {
	return fmt.Sprintf("%s:%s", c.BindIP, c.Port)
}

func (c Config) IsDebug() bool {
	return c.Env == "local" || c.Env == "dev"
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read app config")

		// В проде .env нет — переменные приходят из окружения.
		// Отсутствие файла не ошибка.
		if err := godotenv.Load(); err != nil {
			logger.Debug("no .env file found, reading from environment")
		}

		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
