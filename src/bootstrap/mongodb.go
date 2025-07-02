package bootstrap

import (
	"os"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/mongodb"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

func NewRegisterMongoDB() mongodb.Adapter {
	db, err := mongodb.NewMongoClient(&mongodb.Config{
		Timeout:  util.StringToInt(os.Getenv("MONGO_DB_TIMEOUT")),
		Host:     os.Getenv("MONGO_DB_HOST"),
		Name:     os.Getenv("MONGO_DB_NAME"),
		Password: os.Getenv("MONGO_DB_PASSWORD"),
		Port:     os.Getenv("MONGO_DB_PORT"),
		Username: os.Getenv("MONGO_DB_USER"),
	})

	if err != nil {
		logger.Fatal(err,
			logger.EventName("db"),
			logger.SetField("host", os.Getenv("MONGO_DB_HOST")),
			logger.SetField("port", util.StringToInt(os.Getenv("MONGO_DB_PORT"))),
			logger.SetField("name", os.Getenv("MONGO_DB_NAME")),
			logger.SetField("user", os.Getenv("MONGO_DB_USER")),
		)
	}
	return db
}
