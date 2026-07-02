package khqr

import "fmt"

func ccrc16CCITT(data []byte) uint16 {
    var crc uint16 = 0xFFFF // starting value, always this for CCITT-FALSE

    for _, b := range data {
        crc ^= uint16(b) << 8 // mix this byte into the TOP 8 bits of crc

        for i := 0; i < 8; i++ { // do this 8 times, once per bit
            if crc&0x8000 != 0 { // is the top bit currently 1?
                crc = (crc << 1) ^ 0x1021
            } else {
                crc = crc << 1
            }
        }
    }

    return crc
}

func crc16Hex(data []byte) string {
	result := ccrc16CCITT(data)
	return fmt.Sprintf("%04X", result)
}
