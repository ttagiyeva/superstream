package config

import (
	"strings"

	"github.com/spf13/viper"
)

type queueConfig struct {
	SubscriberTopic string
	SubscriberName  string
	PublishTopic    string
	PoisonTopic     string
}

type kafkaConfig struct {
	Brokers []string
}

type loggerConfig struct {
	Level string
	Env   string
}

// Config holds the configuration data
type Config struct {
	Queue  queueConfig
	Kafka  kafkaConfig
	Logger loggerConfig
}

func New() *Config {
	config := viper.New()

	config.AutomaticEnv()
	config.SetEnvPrefix("superstream")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	conf := &Config{
		Queue: queueConfig{
			SubscriberTopic: config.GetString("queue.subscriber.topic"),
			SubscriberName:  config.GetString("queue.subscriber.group"),
			PublishTopic:    config.GetString("queue.publish.topic"),
			PoisonTopic:     config.GetString("queue.poison.topic"),
		},
		Kafka: kafkaConfig{
			Brokers: config.GetStringSlice("kafka.brokers"),
		},
		Logger: loggerConfig{
			Level: config.GetString("log.level"),
			Env:   config.GetString("log.env"),
		},
	}

	return conf
}
