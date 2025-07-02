package mixpanel

type MixpanelClient interface {
	Track(distincID, event string, property map[string]interface{}) error
}
