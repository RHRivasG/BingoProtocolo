package ports

import (
	"log"
	message "serialport-protocol/protocol"

	"github.com/tarm/serial"
)

//Writer struct
type Writer struct {
	name string
	sw   *serial.Port
}

//NewWriter .
func NewWriter(name string) Writer {
	w := &serial.Config{Name: name, Baud: 115200}
	sw, err := serial.OpenPort(w)
	if err != nil {
		log.Fatal(err)
	}
	return Writer{name, sw}
}

//Writing to a port
func (w *Writer) Writing(messages []string) {
	(w.sw).Write([]byte(message.Unite(message.PutPutLimiter(messages))))
}
