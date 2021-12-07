package ml808

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

const (
	STX = "\x02"
	ETX = "\x03"
	ENQ = "\x05"
	ACK = "\x06"
	A0  = "\x02\x30\x32\x41\x30\x32\x44\x03"
	A2  = "\x02\x30\x32\x41\x32\x32\x42\x03"
	EOT = "\x04"
	CAN = "\x0204cancan6C\x03"
)

var (
	ErrNotConnected = errors.New("Not connected")
	InvalidRsponse  = errors.New("Invalid response")
	VersionWrong    = errors.New("Unable to get right version M8GX-01.5")
)

//Commands G
type ML808 struct {
	port      string
	close     chan bool
	connected bool
	s         *serial.Port
}

func New(port string) *ML808 {
	return &ML808{port: port}
}

func (m *ML808) Connect() error {
	c := &serial.Config{Name: m.port, Baud: 19200, ReadTimeout: time.Second * 3}
	p, e := serial.OpenPort(c)
	if e != nil {
		return e
	}
	m.s = p
	m.connected = true
	if version, err := m.Version(); err != nil && version != "M8GX-01.5" {
		log.Println("Wrong firmware version:", version)
		return VersionWrong
	} else {
		log.Println("firmware version:", version)
	}
	go func() {
		<-m.close
		p.Close()
	}()
	return nil
}

func (m *ML808) Disconnect() error {
	m.connected = false
	m.close <- true
	return nil
}

func (m *ML808) IsConnected() bool {
	return m.connected
}

func (m *ML808) chCommon(ch int) ([]byte, error) {
	if !m.connected {
		return []byte{}, ErrNotConnected
	}
	if err := CmdInit(m.s); err != nil {
		return []byte{}, err
	}
	_, e := m.s.Write(makeCmd([]byte(fmt.Sprintf("%s%03d", "GC", ch))))
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
