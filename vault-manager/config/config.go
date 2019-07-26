package config

import (
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Brokers BrokersConfig
	Servers ServerConfig
	Chains map[string]ChainConfig
	Database DatabaseConfig
}

// BrokersConfig structure
type BrokersConfig struct {
	Consumers map[string]ConsumerConfig
	Producers map[string]ProducerConfig
}

// ServerConfig structure
type ServerConfig struct {
	JWTTokenSecret string `mapstructure:"jwt_token_secret"`
	Monitoring     MonitoringConfig
}

// MonitoringConfig structure
type MonitoringConfig struct {
	Enabled bool
	Host    string
	Port    string
}

// ChainConfig structure
type ChainConfig struct {
	Tx           string
	TxConfirmed  string `mapstructure:"tx_confirmed"`
	Command      string
	CommandReply string `mapstructure:"command_reply"`
}


// ConsumerConfig structure
type ConsumerConfig struct {
	Name   string
	Hosts  []string
	Topics []string
}

// ProducerConfig structure
type ProducerConfig struct {
	Hosts  []string
	Topics []string
}


// TopicConfig structure
type TopicConfig struct {
	Broker string
	Topic  string
}


// DatabaseConfig structure
type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	Name     string
	Port     string
}

// LoadConfig Load server configuration from the yaml file
func LoadConfig(viperConf *viper.Viper) Config {
	var config Config
	err := viperConf.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return config
}