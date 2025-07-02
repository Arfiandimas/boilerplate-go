// Package mariadb
package database

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

/// CreateSession create new session maria db
func CreateSession(cfg *Config) (*sqlx.DB, error) {
	if len(strings.Trim(cfg.Charset, "")) == 0 {
		cfg.Charset = "UTF8"
	}

	connStr := databaseDriver(cfg)

	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), connStr)
	if err != nil {
		return db, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	return db, nil
}

func databaseDriver(cfg *Config) string {
	const connStringTemplate = "%s:%s@tcp(%s:%d)/%s?timeout=%v&charset=%s&parseTime=true&loc=Local"
	const connStringPG = "host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable"
	switch cfg.Driver {
	case "mysql":
		return fmt.Sprintf(connStringTemplate,
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
			cfg.Name,
			cfg.Timeout,
			cfg.Charset,
		)
	case "postgres":
		return fmt.Sprintf(connStringPG,
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.Password,
			cfg.Name,
		)
	}
	return ""
}
