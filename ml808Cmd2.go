package ml808

import (
	"fmt"
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

func (m *ML808) cmd2WithCh(ch int, cmd string) ([]byte, error) {
	if !m.connected {
		return []byte{}, ErrNotConnected
	}
	if err := CmdInit(m.s); err != nil {
		return []byte{}, err
	}
	_, e := m.s.Write(makeCmd([]byte(fmt.Sprintf("%s%03d", cmd, ch))))
	if e != nil {
		return []byte{}, e
	}
	if cmd, dat, err := CmdEndWithData(m.s); err != nil {
		log.Println(err)
		return []byte{}, err
	} else {
		log.Println(string(dat), cmd)
		return dat, nil
	}
}

func (m *ML808) cmd2(cmd string) ([]byte, error) {
	if !m.connected {
		return []byte{}, ErrNotConnected
	}
	if err := CmdInit(m.s); err != nil {
		return []byte{}, err
	}
	_, e := m.s.Write(makeCmd([]byte(fmt.Sprintf("%s   ", cmd))))
	if e != nil {
		return []byte{}, e
	}
	if cmd, dat, err := CmdEndWithData(m.s); err != nil {
		log.Println(err)
		return []byte{}, err
	} else {
		log.Println(string(dat), cmd)
		return dat, nil
	}
}

func (m *ML808) GC(ch int) (p, t, d, f float64, err error) {
	if dat, e := m.cmd2WithCh(ch, "GC"); err != nil {
		err = e
		return
	} else {
		log.Println(string(dat))
		var tp, tt, td, tf int
		_, err = fmt.Sscanf(string(dat), "P%04dT%04dOD%05dOF%05d", &tp, &tt, &td, &tf)
		return float64(tp / 10.0), float64(tt / 1000), float64(td / 10000), float64(tf / 10000), err
	}
}

func (m *ML808) GU(ch int) (startCh, endCh, aiON, aiMode, aiCon, aiTime, aiCount, aiPotlife, outMode int, err error) {
	if dat, e := m.cmd2WithCh(ch, "GU"); err != nil {
		err = e
		return
	} else {
		log.Println(string(dat))
		_, err = fmt.Sscanf(string(dat), "SC%03dEC%03dAI%dM%dT%dAT%04dAC%04dAL%04dTM%d", &startCh, &endCh, &aiON, &aiMode, &aiCon, &aiTime, &aiCount, &aiPotlife, &outMode)
		log.Println(startCh, endCh, aiON, aiMode, aiCon, aiTime, aiCount, aiPotlife, outMode)
		return
	}
}

func (m *ML808) CU() (cnt int, err error) {
	if dat, e := m.cmd2("CU"); err != nil {
		err = e
		return
	} else {
		log.Println(string(dat))
		_, err = fmt.Sscanf(string(dat), "%09d", &cnt)
		return cnt, err
	}
}

func (m *ML808) AR() (cnt int, err error) {
	if dat, e := m.cmd2("AR"); err != nil {
		err = e
		return
	} else {
		log.Println(string(dat))
		_, err = fmt.Sscanf(string(dat), "%09d", &cnt)
		return cnt, err
	}
}
