package khqr

import (
	"crypto/md5"
	"fmt"
)

func HashMD5(qrString string) string {
	sum := md5.Sum([]byte(qrString))
	return fmt.Sprintf("%x", sum)
}
