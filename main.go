package main

import (
	"os"
	"serialport-protocol/protocol"
	port "serialport-protocol/protocol/ports"
)

func main() {
	if os.Args[1:] != nil {
		// 	listener := NewListener(os.Args[1])
		// 	writer := NewWriter(os.Args[2])
		// 	// 	//data:=
		// 	Meeting(listener, writer)
		listener := port.NewListener(os.Args[1])
		writer := port.NewWriter(os.Args[2])
		var arbitro []string
		arbitro = append(arbitro, os.Args[2])
		protocol.Converse(writer, listener, arbitro)
	}

}
