package config

import (
	"log"

	"github.com/caarlos0/env"

	_ "time/tzdata"
)

// Configuration ...
type Configuration struct {
	DB string `env:"DB" envDefault:"postgresql://alpha:GfsND9RR9zgLut@49.0.192.236/alpha_namecard_uat"`

	ENV             string `env:"ENV"`
	Port            string `env:"PORT"`
	PortRestful     string `env:"PORT_RESTFUL" envDefault:"8081"`
	Locale          string `env:"LOCALE" envDefault:"Asia/Bangkok"`
	LogFormat       string `env:"LOG_FORMAT" envDefault:""`
	LogLevel        string `env:"LOG_LEVEL" envDefault:"debug"`
	IsEnableProtoV1 bool   `env:"IS_ENABLE_PROTO_V1" envDefault:"true"`
	IsEnableProtoV2 bool   `env:"IS_ENABLE_PROTO_V2" envDefault:"true"`
	IsDebugDB       bool   `env:"IS_DEBUG_DB" envDefault:"true"`
	//JWT CONFIG
	AESSecretKey     string `env:"AES_SECRET_KEY"`
	SessionTimeout   int    `env:"SESSION_TIMEOUT"`
	PrivateKeyPath   string `env:"PRIVATE_KEY_PATH"`
	PublicKeyPath    string `env:"PUBLIC_KEY_PATH"`
	MessagingAMQPUri string `env:"MESSAGING_AMQP_URI"`

	AppDomain      string `env:"APP_DOMAIN"`
	AppVCardFolder string `env:"APP_VCARD_FOLDER"`
}

// NewConfiguration ...
func NewConfiguration() *Configuration {
	c := &Configuration{}

	if err := env.Parse(c); err != nil {
		log.Fatal(err)
	}

	return c
}
