package oauth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/requester"
)

type oauth struct {
	APIKey    string
	APISecret string
	KeyPath   string
	OauthURL  string

	request requester.Contract
}

func NewOauthService(oauthKey, oauthSecret, oauthURL, keyPath string) Oauth {
	return &oauth{
		KeyPath:   keyPath + "/keydata/",
		request:   requester.New(),
		APIKey:    oauthKey,
		APISecret: oauthSecret,
		OauthURL:  oauthURL,
	}
}

func (o *oauth) ClientGrant(expire uint64) (*ClientTokens, error) {
	uri := o.OauthURL + "/ex/v1/grant/client"
	payload, _ := json.Marshal(&GrantClientRequest{
		ApiKey:    o.APIKey,
		APISecret: o.APISecret,
		ExpiresAt: expire,
	})

	header := map[string]string{
		"Content-Type": "application/json",
	}
	dataRes, err := o.request.POST(uri, header, payload)
	if err != nil {
		return nil, err
	}
	response := &Response{}
	err = json.Unmarshal(dataRes, response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (o *oauth) GetUserClaims(token string) (*UserTokenClaims, error) {
	tokenPayloads := strings.Split(token, ".")
	payloadIssueBody, err := base64.RawStdEncoding.DecodeString(tokenPayloads[1])
	if err != nil {
		return nil, err
	}
	tokenPayloadData := &UserTokenClaims{}
	err = json.Unmarshal(payloadIssueBody, tokenPayloadData)
	if err != nil {
		return nil, err
	}
	return tokenPayloadData, nil
}

func (o *oauth) GetPublicKey() ([]byte, error) {
	response := &OauthPublickeyResponse{}
	publicKeyPath := o.KeyPath + o.APIKey + ".pub"
	if _, err := os.Stat(publicKeyPath); err != nil {
		header := map[string]string{
			"Content-Type": "application/json",
		}
		uri := o.OauthURL + "/ex/v1/keys/public"
		payload, _ := json.Marshal(&OauthPublickeyRequest{
			ApiKey:    o.APIKey,
			APISecret: o.APISecret,
		})

		dataRes, err := o.request.POST(uri, header, payload)
		if err != nil {
			return nil, err
		}

		json.Unmarshal(dataRes, response)
		_, err = os.Stat(o.KeyPath)
		if os.IsNotExist(err) {
			os.MkdirAll(o.KeyPath, os.ModePerm)
		}

		err = os.WriteFile(publicKeyPath, []byte(response.Data.PublicCert), os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return os.ReadFile(publicKeyPath)
}

func (o *oauth) JWTParse(tokenHeaders string, pubkey []byte) (*jwt.Token, error) {
	verifKey, err := jwt.ParseRSAPublicKeyFromPEM(pubkey)
	if err != nil {

		return nil, err
	}
	claims := &ClientClaims{}
	token, err := jwt.ParseWithClaims(tokenHeaders, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return verifKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (o *oauth) Validator(tokenHeaders, userType string) error {
	uri := o.OauthURL + "/ex/v1/validator/" + userType
	header := map[string]string{
		"Content-Type": "application/json",
	}
	_, err := o.request.GET(uri, header)
	if err != nil {
		return err
	}
	return nil
}
