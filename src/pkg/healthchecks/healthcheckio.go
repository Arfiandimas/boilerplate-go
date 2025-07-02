package healthchecks

import (
	"net/http"
	"time"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
)

type healthchecks struct {
	url      string
	logField []logger.Field
}

func NewHeatlhChecksIO() HealthchecksIO {
	return &healthchecks{
		url: "https://hc-ping.com/",
		logField: []logger.Field{
			logger.EventName("HealthChecksIO"),
		},
	}
}

func (h *healthchecks) Ping(uuid string) error {
	var client = &http.Client{
		Timeout: 10 * time.Second,
	}
	url := h.url + uuid
	h.logField = append(h.logField, logger.Any("url", url))
	_, err := client.Head(url)
	if err != nil {
		logger.Info(logger.SetMessageFormat("Healthchecks send error: %s", err.Error()), h.logField...)
		return err
	}

	h.logField = append(h.logField, logger.Any("status", "success"))
	logger.Info(logger.SetMessageFormat("Healthchecks send success"), h.logField...)
	return nil
}
