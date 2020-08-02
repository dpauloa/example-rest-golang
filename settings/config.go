package settings

import (
	"os"

	"dpauloa/example-rest-golang/log"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DatabaseURL string
}

func LoadConfig() (Config, []error) {
	var errs []error

	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, append(errs, err)
	}

	cfg := Config{
		Port: envWithDefault("PORT", "8000"),
		DatabaseURL: envRequired("DATABASE_URL", &errs),
	}

	return cfg, errs
}

func LogErrsAndExit(logger log.Logger, errs []error) {
	LogErrs(logger, errs)
	os.Exit(1)
}

func LogErrs(logger log.Logger, errs []error) {
	logger.Critical("The following errors occurred while loading config:")
	for _, err := range errs {
		logger.Critical(err.Error())
	}
}
