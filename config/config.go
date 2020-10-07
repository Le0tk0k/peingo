package config

import (
	"fmt"
	"os"
)

// DSN はdb接続のためのdsnを返す
func DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?chrset=utf8",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DATABASE"),
	)
}
