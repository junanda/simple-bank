package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct{}

func (config *configImpl) Get(key string) string {
	val, err := os.LookupEnv(key)
	if !err {
		log.Fatal("Env " + key + " Not Found")
	}
	return val
	// return os.Getenv(key)
}

func New(filenames ...string) Config {
	_ = godotenv.Load(filenames...)
	return &configImpl{}
}
