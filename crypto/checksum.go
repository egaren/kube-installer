package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
)

// Sha256Sum - calcultes Sha256 sum
func Sha256Sum(filePath string) (hash string, err error) {
	sha256Sum := sha256.New()
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	sha256Sum.Write(f)
	hash = hex.EncodeToString(sha256Sum.Sum(nil))
	return hash, nil
}
