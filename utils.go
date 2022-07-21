package ml808

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/tarm/serial"
)

func checkSum(data []byte) byte {
	var sum byte
	for _, b := range data {
		sum -= b
	}
	return sum
}

func CmdInit(s *serial.Port) error {
	s.Write([]byte(ENQ))
	log.Println(">>Send ENQ")
	b := make([]byte, 1)
	if s.Read(b); string(b) != ACK {
		log.Printf("<<Receive Error response: %x\n", b)
		return InvalidRsponse
	} else {
		log.Printf("<<Receive device ACK")
		return nil
	}
}

func readUntil(s *serial.Port, end string) ([]byte, error) {
	buff := []byte{}
	b := make([]byte, 1)
	for {
		if n, e := s.Read(b); e != nil {
			return buff, e
		} else {
			buff = append(buff, b[:n]...)
			if string(b) == end {
				return buff, nil
			}
		}
	}
}

func extractData(data []byte) (cmd string, dat []byte, err error) {
	if len(data) < 8 {
		err = errors.New("Invalid data length")
		return
	} else {
		var dlen int
		fmt.Sscanf(string(data[1:3]), "%x", &dlen)
		if len(data) < dlen+3 {
			err = errors.New("Invalid data")
			return
		}
		csum := checkSum(data[1 : len(data)-3])
		log.Println(strings.ToUpper(fmt.Sprintf("%02x", csum)))
		log.Println(string(data[len(data)-3 : len(data)-1]))
		if strings.ToUpper(fmt.Sprintf("%02x", csum)) != string(data[len(data)-3:len(data)-1]) {
			err = errors.New("Check sum error")
			return
		} else {
			cmd = string(data[3:5])
			dat = data[5 : len(data)-3]
			return
		}
	}
}

func CmdEndWithData(s *serial.Port) (cmd string, dat []byte, err error) {
	b := make([]byte, len(A0))
	_, err = io.ReadFull(s, b)
	if err != nil {
		return
	}
	switch string(b) {
	case A0:
		log.Printf("<<Receive device A0")
		s.Write([]byte(ACK))
		log.Println(">>Send ACK")
		if buf, e := readUntil(s, ETX); e != nil {
			log.Println(e)
			err = e
			return
		} else {
			log.Printf("<<Receive data: %x\n", buf)
			cmd, dat, err = extractData(buf)
			if err != nil {
				return
			}
			_, err = s.Write([]byte(EOT))
			if err != nil {
				return
			}
			log.Println(">>Send EOT")
		}
	case A2:
		log.Printf("<<Receive device A0")
		_, err = s.Write([]byte(CAN))
		if err != nil {
			return
		}
	}
	return
}

func CmdEnd(s *serial.Port) (err error) {
	b := make([]byte, len(A0))
	_, err = io.ReadFull(s, b)
	if err != nil {
		return
	}
	switch string(b) {
	case A0:
		log.Printf("<<Receive device A0")
		_, err = s.Write([]byte(EOT))
		if err != nil {
			return
		}
		log.Println(">>Send EOT")
	case A2:
		log.Printf("<<Receive device A2")
		_, err = s.Write([]byte(CAN))
		if err != nil {
			return
		}
	}
	return nil
}

func makeCmd(cmd []byte) []byte {
	var data []byte
	dlen := len(cmd)
	data = append(data, []byte(STX)...)
	data = append(data, []byte(fmt.Sprintf("%02X", dlen))...)
	data = append(data, cmd...)
	data = append(data, []byte(strings.ToUpper(fmt.Sprintf("%x", checkSum(data[1:]))))...)
	data = append(data, []byte(ETX)...)
	return data
}

func checkChan(ch int) error {
	if ch < 0 || ch > 399 {
		log.Fatalf("Invalid channel: %d\n", ch)
		return InvalidChannel
	}
	return nil
}

func checkChPressure(p float64) error {
	if p < 0.0 || p > 800.0 {
		log.Fatalf("Invalid channel pressure: %f\n", p)
		return InvalidChannelPressure
	}
	return nil
}

func checkChTime(t float64) error {
	if t < 0.010 || t > 9.999 {
		log.Fatalf("Invalid channel time: %f\n", t)
		return InvalidChannelTime
	}
	return nil
}
func checkChTime2(t float64) error {
	if t < 0.000 || t > 9.9999 {
		log.Fatalf("Invalid channel time2: %f\n", t)
		return InvalidChannelTime
	}
	return nil
}
