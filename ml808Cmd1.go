package ml808

import "fmt"

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
