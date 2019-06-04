package secret_utils

import (
	sha "crypto/sha256"
	"encoding/base64"
)

func ToSha256(text *string) {
	h := sha.New()
	h.Write([]byte(*text))
	bs:=base64.StdEncoding.EncodeToString(h.Sum(nil))
	*text = bs
}
