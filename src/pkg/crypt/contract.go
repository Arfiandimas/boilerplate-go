package crypt

import "crypto/rsa"

type CryptContract interface {
	BytesToPrivateKey(bytePrivate []byte) (*rsa.PrivateKey, error)
	BytesToPublicKey(bytePublic []byte) (*rsa.PublicKey, error)
	EncryptWithRSA(data string, keyPublic *rsa.PublicKey) (string, error)
	DecryptWithRSA(data string, keyPrivate *rsa.PrivateKey) (string, error)
	GenerateKey(path, file string) (string, error)
}

const BITSIZE = 2048
