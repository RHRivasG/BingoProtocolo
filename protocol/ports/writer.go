package ports

import (
	"log"
	message "serialport-protocol/protocol"

	"github.com/tarm/serial"
)

type Writer struct {
	name string
	sw   *serial.Port
}

func NewWriter(name string) Writer {
	w := &serial.Config{Name: name, Baud: 115200}
	sw, err := serial.OpenPort(w)
	if err != nil {
		log.Fatal(err)
	}
	return Writer{name, sw}
}

func (w *Writer) GetName() string {
	return w.name
}
func (w *Writer) GetSerialPort() *serial.Port {
	return w.sw
}

func (w *Writer) Writing(messages []string) {
	(w.sw).Write([]byte(message.Unite(message.PutLimiter(messages))))
}
