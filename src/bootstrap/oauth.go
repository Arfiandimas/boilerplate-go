package bootstrap

import (
	"os"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/logger"
	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/oauth"
)

func RegistryOauthService() oauth.Oauth {
	apiKey := os.Getenv("OAUTH_API_KEY")
	apiSecret := os.Getenv("OAUTH_API_SECRET")
	uri := os.Getenv("OAUTH_URL")
	keyData := os.Getenv("OAUTH_KEY_PATH")
	client := oauth.NewOauthService(apiKey, apiSecret, uri, keyData)
	_, err := client.GetPublicKey()
	if err != nil {
		logger.Fatal(`oauth service cannot connect, please check your config or network`,
			logger.SetField("key", apiKey),
			logger.SetField("secret", apiSecret),
			logger.SetField("url", uri),
			logger.SetField("error", err.Error()),
		)
	}
	return client
}
