package settings

import (
	"fmt"
	"os"
	"strconv"
)

func env(key string) string {
	return os.Getenv(key)
}

func envRequired(key string, errs *[]error) string {
	value := env(key)

	if value == "" {
		*errs = append(*errs, fmt.Errorf("missing value for '%s'", key))
	}

	return value
}

func envWithDefault(key string, defaultValue string) string {
	value := env(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func envInt(key string, errs *[]error) int {
	value := env(key)
	return parseInt(key, value, errs)
}

func envIntWithDefault(key string, defaultValue int, errs *[]error) int {
	value := env(key)

	if value == "" {
		return defaultValue
	}

	return parseInt(key, value, errs)
}

func parseInt(key, value string, errs *[]error) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		*errs = append(*errs, fmt.Errorf("invalid int value '%s' for %s", value, key))
	}

	return i
}
