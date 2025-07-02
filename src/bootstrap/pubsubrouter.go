package bootstrap

import (
	"os"
	"strings"

	pubsubrouter "github.com/sofyan48/pubsub-router"
)

func PubsubRouterCfg() *pubsubrouter.Config {
	return &pubsubrouter.Config{
		Type:                    "service_account",
		ProjectID:               os.Getenv("GOOGLE_PROJECT_ID"),
		PrivateKeyID:            os.Getenv("GOOGLE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("GOOGLE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("GOOGLE_CLIENT_EMAIL"),
		ClientID:                os.Getenv("GOOGLE_CLIENT_ID"),
		AuthURI:                 os.Getenv("GOOGLE_AUTH_URI"),
		TokenURI:                os.Getenv("GOOGLE_TOKEN_URI"),
		AuthProviderX509CertURL: os.Getenv("GOOGLE_AUTH_PROVIDER"),
		ClientX509CertURL:       os.Getenv("GOOGLE_CLIENT_CERT_URL"),
	}
}

func PubSubRouter() *pubsubrouter.Router {
	subscription := strings.Split(os.Getenv("SUBSCRIBER_TOPIC"), ",")
	if len(subscription) > 0 && os.Getenv("SUBSCRIBER_TOPIC") != "" {
		return pubsubrouter.NewRouter()
	}
	return nil
}
