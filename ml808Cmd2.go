package ml808

import (
	"log"
)

func (m *ML808) Version() (string, error) {
	if !m.connected {
		return "", ErrNotConnected
	}
	if err := CmdInit(m.s); err != nil {
		return "", err
	}
	m.s.Write(makeCmd([]byte("RM   ")))
	if _, dat, err := CmdEndWithData(m.s); err != nil {
		return "", err
	} else {
		return string(dat), nil
	}
}

func (m *ML808) GC(ch int) (p, t, d, f float64, err error) {
	if dat, e := m.chCommon(ch); err != nil {
		err = e
		return
	} else {
		log.Println(string(dat))
		p, t, d, f, err = ParseGC(dat)
		return
	}
}
