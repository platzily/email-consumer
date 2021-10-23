package config

import (
	"github.com/platzily/dealer/utils/constants"
)

type RabbitMQConfig struct {
	URL string
}

func ReadRabbitMQConfig() *RabbitMQConfig {

	urlValue := getEnvVariableAsString(constants.EnvironmentVariables.RABBIT_URL)
	return &RabbitMQConfig{
		URL: urlValue,
	}
}
