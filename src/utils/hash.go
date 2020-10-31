package utils

import (
	"crypto/sha1"
	"fmt"
)

// Hash 对字符串进行SHA1 Hash计算
func Hash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
