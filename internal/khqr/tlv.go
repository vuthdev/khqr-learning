package khqr

import (
	"fmt"
)

func tlv(tag KHQRTag, value string) (string) {
	// hint: tag is always 2 chars, length is always zero-padded to 2 digits
	return fmt.Sprintf("%.2s%02d%s", tag, len(value), value)
}

// func nestedTLV(tag string, children ...string) string {
// 	return fmt.Sprintf()
// }

func checksumPlaceholder() string {
	return "6304"
}