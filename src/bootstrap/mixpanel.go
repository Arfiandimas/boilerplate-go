package bootstrap

import (
	"os"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/mixpanel"
)

// RegistryMixPanel for pkg mixpanel
func RegistryMixPanel() mixpanel.MixpanelClient {
	return mixpanel.NewMixPanel(os.Getenv("MIXPANEL_TOKEN"), os.Getenv("MIXPANEL_SECRET"))
}
