package config

import (
	"fmt"
)

// DSN はdb接続のためのdsnを返す
func DSN() string {
	return fmt.Sprintf("user:password@tcp(db:3306)/peingo?charset=utf8")
}
