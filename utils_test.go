package ml808

import (
	"fmt"
	"testing"
)

func TestCheckSum(t *testing.T) {
	//0x04DI
	if checkSum([]byte("\x00\x30\x34\x44\x49\x20\x20")) != 0xCF {
		t.Error("CheckSum error")
	}
}

func TestMakeCmd(t *testing.T) {
	cmd := makeCmd([]byte(fmt.Sprintf("%s  %03d", "CH", 3)))
	fmt.Printf("hex:%x str:%s", cmd, cmd)
}
