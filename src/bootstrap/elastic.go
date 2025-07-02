package bootstrap

import (
	"os"
	"strings"

	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/elastic"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

// NewSessionElastic ...
func NewSessionElastic() elastic.Client {
	addr := strings.Split(os.Getenv("ELASTICSEARCH_URL"), ",")
	config := &elastic.Configuration{
		Address:  addr,
		Username: os.Getenv("ELASTICSEARCH_USERNAME"),
		Password: os.Getenv("ELASTICSEARCH_PASSWORD"),
	}
	cfgClient := elastic.Config(config)
	client, err := elastic.NewClient(cfgClient)
	if err != nil {
		logger.Fatal(err,
			logger.EventName("elastic"),
			logger.SetField("addres", addr),
			logger.SetField("Username", os.Getenv("ELASTICSEARCH_USERNAME")),
		)
	}
	return client
}
