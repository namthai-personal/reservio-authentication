package hasher

import (
	"crypto/sha256"
	"encoding/base64"
)

func HashSha256(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha
}
