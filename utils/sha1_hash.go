package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
