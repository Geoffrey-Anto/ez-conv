package config

import "os"

func IsDebug() bool {
	env := os.Getenv("DEBUG")
	isdebug := env == "true"

	return isdebug
}
