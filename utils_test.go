package ml808

import (
	"testing"
)

func TestCheckSum(t *testing.T) {
	//0x04DI
	if checkSum([]byte("\x00\x30\x34\x44\x49\x20\x20")) != 0xCF {
		t.Error("CheckSum error")
	}
}
