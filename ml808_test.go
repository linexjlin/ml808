package ml808

import (
	"testing"
)

func TestVersion(t *testing.T) {
	ml := New("COM4")
	err := ml.Connect()
	if err != nil {
		t.Error(err)
		return
	}
	if version, err := ml.Version(); err != nil && version != "M8GX-01.5" {
		t.Error("Version() != 0.0.1")
	}
}
