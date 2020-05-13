package config

import cfg "github.com/Yalantis/go-config"

var config Config

func Get() *Config {
	return &config
}

func Load(fileName string) error {
	return cfg.Init(&config, fileName)
}

type (
	Config struct {
		AppName            string `json:"app_name" envconfig:"API_APP_NAME" default:"api"`
		LogPreset          string `json:"log_preset" envconfig:"API_LOG_PRESET" default:"development"`
		ListenURL          string `json:"listen_url" envconfig:"API_LISTEN_URL" default:":8080"`
		PaginationMaxLimit int64  `json:"pagination_max_limit" envconfig:"API_PAGINATION_MAX_LIMIT" default:"1000"`

		Postgres Postgres `json:"postgres"`
		RabbitMQ RabbitMQ `json:"rabbitmq"`
	}

	Postgres struct {
		Host         string       `json:"host"          envconfig:"POSTGRES_HOST"              default:"localhost"`
		Port         string       `json:"port"          envconfig:"API_POSTGRES_PORT"          default:"5432"`
		Database     string       `json:"database"      envconfig:"API_POSTGRES_DATABASE"      default:"stocks"`
		User         string       `json:"user"          envconfig:"API_POSTGRES_USER"          default:"postgres"`
		Password     string       `json:"password"      envconfig:"API_POSTGRES_PASSWORD"      default:"12345"`
		PoolSize     int          `json:"pool_size"     envconfig:"API_POSTGRES_POOL_SIZE"     default:"10"`
		MaxRetries   int          `json:"max_retries"   envconfig:"API_POSTGRES_MAX_RETRIES"   default:"5"`
		ReadTimeout  cfg.Duration `json:"read_timeout"  envconfig:"API_POSTGRES_READ_TIMEOUT"  default:"10s"`
		WriteTimeout cfg.Duration `json:"write_timeout" envconfig:"API_POSTGRES_WRITE_TIMEOUT" default:"10s"`
	}

	RabbitMQ struct {
		Host     string `json:"host" envconfig:"API_RABBITMQ_HOST" default:"localhost"`
		Port     string `json:"port" envconfig:"API_RABBITMQ_PORT" default:"5672"`
		User     string `json:"user" envconfig:"API_RABBITMQ_USER" default:"rabbit"`
		Password string `json:"password" envconfig:"API_RABBITMQ_PASSWORD" default:"12345"`
	}
)
