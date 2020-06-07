package ports

import (
	"log"
	message "serialport-protocol/protocol"

	"github.com/tarm/serial"
)

//Listener struct
type Listener struct {
	name string
	sl   *serial.Port
}

//NewListener .
func NewListener(name string) Listener {
	l := &serial.Config{Name: name, Baud: 115200}
	sl, err := serial.OpenPort(l)
	if err != nil {
		log.Fatal(err)
	}
	return Listener{name, sl}
}

//Listening a port
func (l *Listener) Listening() []string {
	buf := make([]byte, 128)
	n, err := (l.sl).Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	messages := message.CutLimiter(message.Separate(string(buf[:n])))
	return messages
}
