package crypt

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

// Key RSA Config Struct
type keyRSAConfig struct {
}

// Initialize Function in Helper Cryptography
func NewCrypto() CryptContract {
	return &keyRSAConfig{}
}

// BytesToPrivateKey Function
func (c *keyRSAConfig) BytesToPrivateKey(bytePrivate []byte) (*rsa.PrivateKey, error) {
	var err error
	// Decode RSA Private Key Bytes to PEM Block Bytes
	pemBlock, _ := pem.Decode(bytePrivate)
	isEncrypted := x509.IsEncryptedPEMBlock(pemBlock)
	byteBlock := pemBlock.Bytes

	// Check If RSA Private Key Bytes Encrypted
	// If Encrypted Try to Decode it
	if isEncrypted {
		byteBlock, err = x509.DecryptPEMBlock(pemBlock, nil)
		if err != nil {
			return nil, err
		}
	}

	// Parse RSA Private Key Using PKCS1
	rsaKeyPrivate, err := x509.ParsePKCS1PrivateKey(byteBlock)
	if err != nil {
		return nil, err
	}

	// Return RSA Public Key
	return rsaKeyPrivate, nil
}

// BytesToPublicKey Function
func (c *keyRSAConfig) BytesToPublicKey(bytePublic []byte) (*rsa.PublicKey, error) {
	var err error

	// Decode RSA Public Key Bytes to PEM Block Bytes
	pemBlock, _ := pem.Decode(bytePublic)
	isEncrypted := x509.IsEncryptedPEMBlock(pemBlock)
	byteBlock := pemBlock.Bytes

	// Check If RSA Public Key Bytes Encrypted
	// If Encrypted Try to Decode it
	if isEncrypted {
		byteBlock, err = x509.DecryptPEMBlock(pemBlock, nil)
		if err != nil {
			return nil, err
		}
	}

	// Parse RSA Public Key Using PKCIX
	rsaKeyPublic, err := x509.ParsePKIXPublicKey(byteBlock)
	if err != nil {
		return nil, err
	}

	// Return RSA Public Key
	return rsaKeyPublic.(*rsa.PublicKey), nil
}

// EncryptWithRSA Function
func (c *keyRSAConfig) EncryptWithRSA(data string, keyPublic *rsa.PublicKey) (string, error) {
	// Generate New SHA512 Hash
	hash := sha512.New()

	// Encrypt Plain Text to Chiper Text Using RSA Encryption OAEP
	chiperText, err := rsa.EncryptOAEP(hash, rand.Reader, keyPublic, []byte(data), nil)
	if err != nil {
		return "", err
	}

	// Compress Chiper Text to Base64 Format
	compressText := base64.StdEncoding.EncodeToString(chiperText)

	// Return Compressed Text
	return compressText, nil
}

// DecryptWithRSA Function
func (c *keyRSAConfig) DecryptWithRSA(data string, keyPrivate *rsa.PrivateKey) (string, error) {
	// Generate New SHA512 Hash
	hash := sha512.New()

	// Decompress Chiper Text from Base64 Format
	decompressText, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	// Decrypt Chiper Text to Plain Text Using RSA Encryption OAEP
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, keyPrivate, []byte(decompressText), nil)
	if err != nil {
		return "", err
	}

	// Return Chiper Text
	return string(plainText), nil
}

// MD5Hash ...
func (c *keyRSAConfig) MD5Hash(value string) string {
	h := md5.New()
	io.WriteString(h, value)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GenerateKey ..
func (c *keyRSAConfig) GenerateKey(path, file string) (string, error) {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, BITSIZE)
	if err != nil {
		return "", err
	}
	publicKey := key.PublicKey
	files := path + "/" + file
	err = c.savePEMKey(files+".key", key)
	if err != nil {
		return "", err
	}
	err = c.savePublicPEMKey(files+".pub", publicKey)
	if err != nil {
		return "", err
	}
	return files, nil
}

func (c *keyRSAConfig) savePEMKey(fileName string, key *rsa.PrivateKey) error {
	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	return pem.Encode(outFile, privateKey)
}

func (c *keyRSAConfig) savePublicPEMKey(fileName string, pubkey rsa.PublicKey) error {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(&pubkey)
	if err != nil {
		return err
	}

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	pemfile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer pemfile.Close()

	return pem.Encode(pemfile, pemkey)
}
