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
		AppName   string `json:"app_name" envconfig:"FETCHER_APP_NAME" default:"fetcher"`
		LogPreset string `json:"log_preset" envconfig:"FETCHER_LOG_PRESET" default:"development"`
		ListenURL string `json:"listen_url" envconfig:"FETCHER_LISTEN_URL" default:":8080"`

		HTTPClient HTTPClient `json:"http_client"`

		RabbitMQ RabbitMQ `json:"rabbitmq"`

		Stocks Stocks `json:"stocks"`
	}

	HTTPClient struct {
		Timeout cfg.Duration `json:"timeout" envconfig:"FETCHER_HTTP_CLIENT_TIMEOUT" default:"10s"`
	}

	Stocks struct {
		Adapter string  `json:"adapter" envconfig:"FETCHER_STOCKS_ADAPTER" default:"finnhub"`
		Finnhub Finnhub `json:"finnhub"`
		Symbols Symbols `json:"symbols"`
	}

	Finnhub struct {
		BaseURL string `json:"base_url" envconfig:"FETCHER_FINNHUB_BASE_URL" default:"https://finnhub.io/api/v1"`
		Token   string `json:"token" envconfig:"FETCHER_FINNHUB_TOKEN"`
	}

	Symbols struct {
		RefreshFreq cfg.Duration `json:"refresh_freq" envconfig:"FETCHER_REFRESH_FREQ" default:"1h"`
	}

	RabbitMQ struct {
		Host     string `json:"host" envconfig:"FETCHER_RABBITMQ_HOST" default:"localhost"`
		Port     string `json:"port" envconfig:"FETCHER_RABBITMQ_PORT" default:"5672"`
		User     string `json:"user" envconfig:"FETCHER_RABBITMQ_USER" default:"rabbit"`
		Password string `json:"password" envconfig:"FETCHER_RABBITMQ_PASSWORD" default:"12345"`
	}
)
