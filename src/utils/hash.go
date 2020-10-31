package utils

import (
	"crypto/sha1"
	"fmt"
)

func Hash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
