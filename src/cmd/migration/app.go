// Package migration
package migration

import (
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/database"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/util"
)

func migrateDatabase() {
	database.DatabaseMigration(&database.Config{
		Driver:       os.Getenv("DB_DRIVER"),
		Host:         os.Getenv("DB_HOST_MIGRATION"),
		Name:         os.Getenv("DB_NAME_MIGRATION"),
		Password:     os.Getenv("DB_PASSWORD_MIGRATION"),
		Port:         util.StringToInt(os.Getenv("DB_PORT_MIGRATION")),
		User:         os.Getenv("DB_USER_MIGRATION"),
		Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT_MIGRATION"))) * time.Second,
		MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN_MIGRATION")),
		MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN_MIGRATION")),
		MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME_MIGRATION"))) * time.Millisecond,
		Charset:      os.Getenv("DB_CHARSET_MIGRATION"),
	})
}

func Start() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "db:migrate",
		Short: "database migration",
		Run: func(c *cobra.Command, args []string) {
			migrateDatabase()
		},
	}

	migrateCmd.Flags().BoolP("version", "", false, "print version")
	migrateCmd.Flags().StringP("dir", "", "database/migration/", "directory with migration files")
	migrateCmd.Flags().StringP("table", "", "db", "migrations table name")
	migrateCmd.Flags().BoolP("verbose", "", false, "enable verbose mode")
	migrateCmd.Flags().BoolP("guide", "", false, "print help")
	return migrateCmd
}
