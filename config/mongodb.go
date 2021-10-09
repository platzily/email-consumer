package config

import (
	"github.com/platzily/consumer/utils/constants"
	log "github.com/sirupsen/logrus"
)

type MongoDBConfig struct {
	URL      string
	Database string
}

func ReadMongoDBConfig() *MongoDBConfig {

	urlValue := getEnvVariableAsString(constants.EnvironmentVariables.MONGO_URL)
	databaseName := getEnvVariableAsString(constants.EnvironmentVariables.MONGO_DATABASE)
	return &MongoDBConfig{
		URL: urlValue,
		Database: databaseName,
	}
}

func getEnvVariableAsString(name string) string {
	envVar := os.Getenv(name)

	if len(envVar) == {
		log.Fatalf("Environment variable %s is not set", name)
	}

	return envVar
}