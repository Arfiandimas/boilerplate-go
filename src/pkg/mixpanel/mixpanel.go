package mixpanel

import (
	"os"
	"time"

	mxpanel "github.com/dukex/mixpanel"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/logger"
)

type mixPanel struct {
	client mxpanel.Mixpanel
}

func NewMixPanel(token, secret string) MixpanelClient {
	return &mixPanel{
		client: mxpanel.NewWithSecret(token, secret, ""),
	}
}

func (m *mixPanel) Track(distincID, event string, property map[string]interface{}) error {
	logger.Info(logger.SetMessageFormat("Send to mixpanel %s", event))
	if os.Getenv("MIXPANEL_TOKEN") != "" || os.Getenv("MIXPANEL_SECRET") != "" {
		now := time.Now()
		err := m.client.Track(distincID, event, &mxpanel.Event{
			IP:         "",
			Timestamp:  &now,
			Properties: property,
		})
		if err != nil {
			logger.Info(logger.SetMessageFormat("Error send to mixpanel: %s", err.Error()))
			return err
		}
		return nil
	}
	return nil
}
