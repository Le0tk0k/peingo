package db

import (
	"fmt"

	"github.com/Le0tk0k/peingo/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to Open mysql: %v", err)
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping: %v", err)
	}

	return db, nil
}
