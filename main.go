package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

//Writer struct
type Writer struct {
	name string
	sw   *serial.Port
}

//Listener struct
type Listener struct {
	name string
	sl   *serial.Port
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

//NewListener .
func NewListener(name string) Listener {
	l := &serial.Config{Name: name, Baud: 115200}
	sl, err := serial.OpenPort(l)
	if err != nil {
		log.Fatal(err)
	}
	return Listener{name, sl}
}

//Writing to a port
func (w *Writer) Writing(text string) {
	(w.sw).Write([]byte(text))
}

//Listening a port
func (l *Listener) Listening() []string {
	buf := make([]byte, 128)
	n, err := (l.sl).Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	texts := Separate(string(buf[:n]))
	return texts
}

//Separate a text
func Separate(t string) []string {
	return strings.SplitAfter(t, "O75")
}

//Meeting .
func Meeting(l Listener, w Writer) []string {

	w.Writing(w.name)
	fmt.Println("Esperando...")
	texts := l.Listening()
	if texts[0] == w.name {
		fmt.Println("Yo soy el arbitro")
	} else {
		fmt.Println(texts[0] + " es el arbitro")
		w.Writing(texts[0])
	}

	return nil
}

func main() {
	if os.Args[1:] != nil {
		listener := NewListener(os.Args[1])
		writer := NewWriter(os.Args[2])
		// 	//data:=
		Meeting(listener, writer)
	}
	fmt.Print(strings.TrimSuffix(strings.TrimPrefix("B1V1Hello, GophersU75O75", "B1"), "O75"))
}
