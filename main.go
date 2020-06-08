package main

import (
	"fmt"
	"os"
	"serialport-protocol/protocol"
)

func main() {
	if os.Args[1:] != nil {
		protocol := protocol.NewProtocol(os.Args[1], os.Args[2])
		var arbitro []string
		arbitro = append(arbitro, os.Args[2])
		messages := protocol.Converse(arbitro)
		if messages[0] == protocol.GetWriterName() {
			fmt.Println("Yo soy el arbitro")
		} else {
			fmt.Println(messages[0] + " es el arbitro")
			protocol.EndConversation(messages)
		}

	} else {
		fmt.Println("No hay argumentos")
	}
}
