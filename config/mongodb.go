package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/platzily/email-consumer/utils/constants"
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
		URL:      urlValue,
		Database: databaseName,
	}
}

func getEnvVariableAsString(name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envVar := os.Getenv(name)

	if len(envVar) == 0 {
		log.Fatalf("Environment variable %s is not set", name)
	}

	return envVar
}
