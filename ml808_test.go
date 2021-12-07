package ml808

import (
	"log"
	"testing"
)

var ml = New("COM4")

func TestInit(t *testing.T) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	err := ml.Connect()
	if err != nil {
		t.Error(err)
		return
	}
}
func TestVersion(t *testing.T) {
	if version, err := ml.Version(); err != nil && version != "M8GX-01.5" {
		t.Error("Version() != 0.0.1")
	}
}

func TestGC(t *testing.T) {
	ml.GC(1)
}

func TestCH(t *testing.T) {
	ml.CH(209)
}
