package bootstrap

import (
	"github.com/kiriminaja/kaj-rest-engine-go/src/pkg/crypt"
)

// Crypto initializer for best security
func Crypto() crypt.CryptContract {
	return crypt.NewCrypto()
}
