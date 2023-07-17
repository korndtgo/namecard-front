package byte2hash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
)

func HashSha256Bytes(data []byte) (string, error) {
	byteHashed := sha256.Sum256(data)
	fileHashed := base64.StdEncoding.EncodeToString(byteHashed[:])
	return fileHashed, nil
}

func HashMD5OnBytes(data []byte) (string, error) {
	byteHashed := md5.Sum(data)
	fileHashed := base64.StdEncoding.EncodeToString(byteHashed[:])
	return fileHashed, nil
}
