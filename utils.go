package ml808

import (
	"errors"
	"fmt"
	"io"
	"log"

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
		csum := checkSum(data[1 : dlen+1+2])
		fmt.Println(fmt.Sprintf("%x", csum))
		fmt.Println(fmt.Sprintf("%s", data[dlen+1+1:dlen+2+1]))
		if fmt.Sprintf("%x", csum) != string(data[dlen+1+1:dlen+2+1]) {
			err = errors.New("Check sum error")
			return
		} else {
			cmd = string(data[3:5])
			dat = data[5 : 5+dlen-2]
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
			log.Printf("<<Receive data: %s\n", buf)
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
	}
	return
}
