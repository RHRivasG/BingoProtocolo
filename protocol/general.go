package protocol

import (
	port "serialport-protocol/protocol/ports"
)

//Converse .
func Converse(w port.Writer, l port.Listener, message []string) []string {
	w.Writing(message)
	return l.Listening()
}