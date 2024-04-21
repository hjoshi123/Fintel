package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	Spec *Configuration
)

type Configuration struct {
	Environment        string `mapstructure:"environment"`
	Version            string `mapstructure:"version"`
	Port               uint16 `mapstructure:"port"`
	DBPort             string `mapstructure:"db_port"`
	LogLevel           string `mapstructure:"log_level"`
	DBUser             string `mapstructure:"db_user"`
	DBHost             string `mapstructure:"db_host"`
	DBName             string `mapstructure:"db_name"`
	DBPassword         string `mapstructure:"db_password"`
	KafkaClientID      string `mapstructure:"kafka_clientID"`
	KafkaTxnID         string `mapstructure:"kafka_txnID"`
	KafkaBrokers       string `mapstructure:"kafka_brokers"`
	KafkaGroup         string `mapstructure:"kafka_group"`
	SaslEnable         bool   `mapstructure:"sasl_enable"`
	SaslMechanism      string `mapstructure:"sasl_mechanism"`
	KafkaUsername      string `mapstructure:"kafka_username"`
	KafkaPassword      string `mapstructure:"kafka_password"`
	AlphaVantageApiKey string `mapstructure:"alpha_vantage"`
}

func init() {
	Spec = new(Configuration)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Spec)
}

func IsDevelopment() bool {
	return Spec.Environment == "development"
}

func IsProduction() bool {
	return Spec.Environment == "production"
}
