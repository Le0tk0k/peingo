package config

import (
	"fmt"
	"os"
)

// DSN はdb接続のためのdsnを返す
func DSN() string {
	return fmt.Sprintf("user:password@tcp(db:3306)/peingo?charset=utf8")
}

// WebHook はslack通知のWebHook URL を返す
func WebHook() string {
	webHook := os.Getenv("WEBHOOK")
	return fmt.Sprintf("https://hooks.slack.com/services/%s", webHook)
}
