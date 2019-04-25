package secret_utils

import (
	sha "crypto/sha256"
)

func ToSha256(text *string) {
	h := sha.New()
	h.Write([]byte(*text))
	bs := string(h.Sum(nil))
	text = &bs
}
