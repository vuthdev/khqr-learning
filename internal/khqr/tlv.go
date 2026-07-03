package khqr

import (
	"fmt"
)

func tlv(tag KHQRTag, value string) (string) {
	return fmt.Sprintf("%.2s%02d%s", tag, len(value), value)
}

// func nestedTLV(tag string, children ...string) string {
// 	return fmt.Sprintf()
// }

func checksumPlaceholder() string {
	return "6304"
}