package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	CloudEndpoint         string `mapstructure:"CLOUD_ENDPOINT"`
	CloudAccessKey        string `mapstructure:"CLOUD_ACCESS_KEY"`
	CloudSecretKey        string `mapstructure:"CLOUD_SECRET_KEY"`
	CloudBucketName       string `mapstructure:"CLOUD__BUCKET_NAME"`
	SignedUrlExpiryMinute int    `mapstructure:"SIGNED_URL_EXPIRY_MINUTE"`
	AppPort               string `mapstructure:"APP_PORT"`
}

func NewEnv() *Env {

	var env Env
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env :", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded")
	}

	return &env
}
