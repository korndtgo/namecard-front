package auth

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/hartsfield/gencrypt"
)

//DecryptPayload ...
func DecryptPayload(key, cypher string, out interface{}) error {
	byteKey := []byte(key)
	gcm, err := gencrypt.NewGCM(byteKey)

	if err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(cypher)

	if err != nil {
		return err
	}

	dec, err := gcm.AESDecrypt(decoded)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(dec, &out); err != nil {
		return err
	}

	return nil
}

//EncryptPayload ...
func EncryptPayload(key string, payload interface{}) (string, error) {
	byteKey := []byte(key)
	gcm, err := gencrypt.NewGCM(byteKey)

	if err != nil {
		return "", err
	}

	reqBodyBytes := new(bytes.Buffer)

	if err := json.NewEncoder(reqBodyBytes).Encode(&payload); err != nil {
		return "", err
	}

	enc, err := gcm.AESEncrypt(reqBodyBytes.Bytes())

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}

//GenerateJWTToken ...
func GenerateJWTToken(keyPath, tokenType, secret string, expire int64) (string, error) {
	key, err := ioutil.ReadFile(keyPath)

	if err != nil {
		return "", err
	}

	sign, err := jwt.ParseRSAPrivateKeyFromPEM(key)

	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodRS256)

	token.Claims = &JWTClaims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expire,
		},
		JWTPayload: JWTPayload{
			Type:   tokenType,
			Secret: secret,
		},
	}

	return token.SignedString(sign)
}

//EncryptSha1 ...
func EncryptSha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

//EncryptSha256 ...
func EncryptSha256(str string) string {
	data := []byte(str)
	hash := sha256.Sum256(data)

	return fmt.Sprintf("%x", hash[:])
}

//EncryptVerifiedToken ...
func EncryptVerifiedToken(key string, claim map[string]interface{}) (string, error) {
	byteKey := []byte(key)
	gcm, err := gencrypt.NewGCM(byteKey)

	if err != nil {
		return "", err
	}

	reqBodyBytes := new(bytes.Buffer)

	if err := json.NewEncoder(reqBodyBytes).Encode(&claim); err != nil {
		return "", err
	}

	enc, err := gcm.AESEncrypt(reqBodyBytes.Bytes())

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(enc), nil
}

//DecryptVerifiedToken ...
func DecryptVerifiedToken(key, cypher string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	byteKey := []byte(key)
	gcm, err := gencrypt.NewGCM(byteKey)

	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(cypher)

	if err != nil {
		return nil, err
	}

	dec, err := gcm.AESDecrypt(decoded)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(dec, &out); err != nil {
		return nil, err
	}

	return out, nil
}
