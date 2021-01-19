package cryptoutils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// GetMd5 - to encrypt some string
func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))

	return hex.EncodeToString(hash.Sum(nil))
}

//GenerateHashFromStruct - receive a struct as parameter and return a hash
func GenerateHashFromStruct(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
