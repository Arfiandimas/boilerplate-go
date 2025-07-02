package oauth

import "github.com/dgrijalva/jwt-go"

//go:generate easytags $GOFILE json, form

type Oauth interface {
	GetPublicKey() ([]byte, error)
	GetUserClaims(token string) (*UserTokenClaims, error)
	JWTParse(tokenHeaders string, pubkey []byte) (*jwt.Token, error)
	ClientGrant(expire uint64) (*ClientTokens, error)
	Validator(tokenHeaders, userType string) error
}
type UserTokenClaims struct {
	ID        uint64           `json:"id"`
	Name      string           `json:"name"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	SmcUserID uint64           `json:"smc_user_id"`
	IsActive  bool             `json:"is_active"`
	Scope     []ScopeStructure `json:"scope"`
	ClientID  uint64           `json:"client_id"`
	jwt.StandardClaims
}

type ScopeStructure struct {
	ID    int              `json:"id"`
	Name  string           `json:"name"`
	Roles []RolesStructure `json:"roles"`
}

type RolesStructure struct {
	ID        int      `json:"id"`
	RolesName string   `json:"roles_name"`
	Acls      []string `json:"acls"`
}

type OauthPublickeyResponse struct {
	Name    string                      `json:"name"`
	Message string                      `json:"message"`
	Data    *OauthPublickeyResponseData `json:"data"`
}

type OauthPublickeyResponseData struct {
	PublicCert string `json:"public_cert"`
}

type OauthPublickeyRequest struct {
	ApiKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

type GrantClientRequest struct {
	ApiKey    string `json:"api_key"`
	APISecret string `json:"api_secret"`
	ExpiresAt uint64 `json:"expired_at"`
}

type ClientClaims struct {
	ClientID uint64 `json:"client_id"`
	jwt.StandardClaims
}

type Response struct {
	Name    string        `json:"name"`
	Message interface{}   `json:"message,omitempty"`
	Errors  interface{}   `json:"errors,omitempty"`
	Data    *ClientTokens `json:"data,omitempty"`
	Lang    string        `json:"-"`
}

type ClientTokens struct {
	Type      string `json:"type"`
	ExpiresAt int64  `json:"expires_at"`
	Token     string `json:"access_token"`
}
