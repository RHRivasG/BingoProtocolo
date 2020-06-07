package protocol

import (
	port "serialport-protocol/protocol/ports"
)

func Converse(w port.Writer, l port.Listener, message []string) []string {
	w.Writing(message)
	return l.Listening()
}

func Finalize(w port.Writer, lastMessage string) {
	w.Writer(lastMessage)
}
