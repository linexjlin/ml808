package ml808

import (
	"fmt"
	"log"
	"math"
)

func (m *ML808) DI() (err error) {
	return m.pureCmd1("DI")
}

func (m *ML808) SE() (err error) {
	return m.pureCmd1("SE")
}

func (m *ML808) pureCmd1(cmd string) (err error) {
	if !m.connected {
		return ErrNotConnected
	}
	if err = CmdInit(m.s); err != nil {
		return err
	}
	_, err = m.s.Write(makeCmd([]byte(fmt.Sprintf("%s  ", cmd))))
	if err != nil {
		return err
	}
	if err = CmdEnd(m.s); err != nil {
		return
	}
	return nil
}

func (m *ML808) SS(startCh, endCh int) (err error) {
	if !m.connected {
		return ErrNotConnected
	}
	if err = CmdInit(m.s); err != nil {
		return err
	}
	_, err = m.s.Write(makeCmd([]byte(fmt.Sprintf("%s  S%03dE%03d", "SS", startCh, endCh))))
	if err != nil {
		return err
	}
	if err = CmdEnd(m.s); err != nil {
		return
	}
	return nil
}

func (m *ML808) CH(ch int) (err error) {
	if !m.connected {
		return ErrNotConnected
	}
	if err = CmdInit(m.s); err != nil {
		return err
	}
	_, err = m.s.Write(makeCmd([]byte(fmt.Sprintf("%s  %03d", "CH", ch))))
	if err != nil {
		return err
	}
	if err = CmdEnd(m.s); err != nil {
		return
	}
	return nil
}

func (m *ML808) TT() (err error) {
	return m.pureCmd1("TT")
}

func (m *ML808) MT() (err error) {
	return m.pureCmd1("MT")
}

func (m *ML808) PH(ch int, p float64) (err error) {
	if err = checkChan(ch); err != nil {
		return err
	}
	if err = checkChPressure(p); err != nil {
		return err
	}
	p = p * 10
	return m.setChanValueCmd("PH", ch, fmt.Sprintf("P%04d", int(math.Floor(p+0.5))))
}

func (m *ML808) DH(ch int, t float64) (err error) {
	if err = checkChan(ch); err != nil {
		return err
	}
	if err = checkChTime(t); err != nil {
		return err
	}
	t = t * 1000
	return m.setChanValueCmd("DH", ch, fmt.Sprintf("T%04d", int(math.Floor(t+0.5))))
}

func (m *ML808) setChanValueCmd(cmd string, ch int, v string) (err error) {
	if !m.connected {
		return ErrNotConnected
	}
	if err = CmdInit(m.s); err != nil {
		return err
	}
	_data := makeCmd([]byte(fmt.Sprintf("%s  CH%03d%s", cmd, ch, v)))
	log.Printf("%x \n %s\n", _data, _data)
	_, err = m.s.Write(makeCmd([]byte(fmt.Sprintf("%s  CH%03d%s", cmd, ch, v))))
	if err != nil {
		return err
	}
	if err = CmdEnd(m.s); err != nil {
		return
	}
	return nil
}

func (m *ML808) SC(ch int, p, t, d, f float64) (err error) {
	if err = checkChan(ch); err != nil {
		return err
	}
	if err = checkChPressure(p); err != nil {
		return err
	}
	if err = checkChTime(t); err != nil {
		return err
	}
	if err = checkChTime2(d); err != nil {
		return err
	}
	if err = checkChTime2(f); err != nil {
		return err
	}
	p = p * 10
	t = t * 1000
	d = d * 10000
	f = f * 10000
	return m.setChanValueCmd("SC", ch, fmt.Sprintf("P%04dT%04dOD%05dOF%05d", int(math.Floor(p+0.5)), int(math.Floor(t+0.5)), int(math.Floor(d+0.5)), int(math.Floor(f+0.5))))
}

func (m *ML808) CD() (err error) {
	return m.pureCmd1("CD")
}
