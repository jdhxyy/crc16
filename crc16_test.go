package crc16

import (
	"fmt"
	"testing"
)

func TestCheckSum(t *testing.T) {
	arr := []uint8{1, 2, 3}
	fmt.Printf("0x%04x", Checksum(arr))
}
