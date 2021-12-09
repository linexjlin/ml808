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
	if _, _, _, _, err := ml.GC(1); err != nil {
		t.Fail()
	}
}

func TestCH(t *testing.T) {
	if err := ml.CH(208); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestGU(t *testing.T) {
	if _, _, _, _, _, _, _, _, _, err := ml.GU(209); err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestCU(t *testing.T) {
	if cnt, err := ml.CU(); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		log.Println(cnt)
	}
}

func TestAR(t *testing.T) {
	if cnt, err := ml.AR(); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		log.Println(cnt)
	}
}
