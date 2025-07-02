package bootstrap

import (
	"os"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/database"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

// RegistryMariaMasterSlave initialize maria or mysql session
func RegistryMariaMasterSlave() database.Adapter {
	db, err := database.NewMariaMasterSlave(
		&database.Config{
			Driver:       os.Getenv("DB_DRIVER"),
			Host:         os.Getenv("DB_HOST"),
			Name:         os.Getenv("DB_NAME"),
			Password:     os.Getenv("DB_PASSWORD"),
			Port:         util.StringToInt(os.Getenv("DB_PORT")),
			User:         os.Getenv("DB_USER"),
			Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT"))) * time.Second,
			MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN")),
			MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN")),
			MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME"))) * time.Millisecond,
			Charset:      os.Getenv("DB_CHARSET"),
			Debug:        util.StringToBool(os.Getenv("DB_DEBUG_MODE")),
		},

		&database.Config{
			Driver:       os.Getenv("DB_DRIVER"),
			Host:         os.Getenv("DB_HOST_READ"),
			Name:         os.Getenv("DB_NAME_READ"),
			Password:     os.Getenv("DB_PASSWORD_READ"),
			Port:         util.StringToInt(os.Getenv("DB_PORT_READ")),
			User:         os.Getenv("DB_USER_READ"),
			Timeout:      time.Duration(util.StringToInt(os.Getenv("DB_TIMEOUT_READ"))) * time.Second,
			MaxOpenConns: util.StringToInt(os.Getenv("DB_MAX_OPEN_CONN_READ")),
			MaxIdleConns: util.StringToInt(os.Getenv("DB_MAX_IDLE_CONN_READ")),
			MaxLifetime:  time.Duration(util.StringToInt(os.Getenv("DB_MAX_LIFETIME_READ"))) * time.Millisecond,
			Charset:      os.Getenv("DB_CHARSET_READ"),
			Debug:        util.StringToBool(os.Getenv("DB_DEBUG_MODE_READ")),
		},
	)

	if err != nil {
		logger.Fatal(err,
			logger.EventName("db"),
			logger.SetField("host_read", os.Getenv("DB_HOST_READ")),
			logger.SetField("port_read", util.StringToInt(os.Getenv("DB_PORT_READ"))),
			logger.SetField("host_write", os.Getenv("DB_HOST")),
			logger.SetField("port_write", util.StringToInt(os.Getenv("DB_PORT"))),
		)
	}

	return db
}
