package protocol

import (
	port "serialport-protocol/protocol/ports"
)

type ProtocolI interface {
	Converse(message []string) []string
	EndConvesation(lastMessage []string)
	GetWriterName() string
}

type Protocol struct {
	l port.Listener
	w port.Writer
}

func NewProtocol(listener string, writer string) Protocol {
	return Protocol{port.NewListener(listener), port.NewWriter(writer)}
}

//Converse .
func (p *Protocol) Converse(message []string) []string {
	p.w.Writing(message)
	return p.l.Listening()
}

func (p *Protocol) EndConversation(lastMessage []string) {
	p.w.Writing(lastMessage)
}

func (p *Protocol) GetWriterName() string {
	return p.w.GetName()
}
