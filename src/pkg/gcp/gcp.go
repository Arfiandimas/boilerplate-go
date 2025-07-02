package gcp

import (
	"encoding/json"
	"os"

	"google.golang.org/api/option"
)

type googlePKG struct {
}

func New() Contract {
	return &googlePKG{}
}

func (g *googlePKG) Option() []option.ClientOption {
	return optionCredential(configEnvironment())
}

func configEnvironment() []byte {
	confiCredential := ConfigServer{
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
	data, _ := json.Marshal(confiCredential)

	return data
}

func optionCredential(cfg []byte) []option.ClientOption {
	return []option.ClientOption{
		option.WithCredentialsJSON(cfg),
	}
}
