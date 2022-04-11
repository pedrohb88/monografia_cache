package cache

import (
	"crypto/md5"
	"fmt"
)

func CalculateDigest(i interface{}) string {
	data := []byte(fmt.Sprintf("%#v", i))
	return fmt.Sprintf("%x", md5.Sum(data))[:7]
}
