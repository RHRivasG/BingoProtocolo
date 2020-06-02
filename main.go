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
}

//Listener struct
type Listener struct {
	name string
}

//Writing to a port
func (w Writer) Writing(text string) {
	writer := &serial.Config{Name: w.name, Baud: 115200}
	sw, err := serial.OpenPort(writer)
	if err != nil {
		log.Fatal(err)
	}
	sw.Write([]byte(text))
	if err != nil {
		log.Fatal(err)
	}
}

//Listening a port
func (l Listener) Listening() []string {
	listener := &serial.Config{Name: l.name, Baud: 115200}
	sl, err := serial.OpenPort(listener)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	n, err := sl.Read(buf)
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
	if texts[1] == w.name {
		fmt.Println("Yo soy el arbitro")
	} else {
		fmt.Println(texts[1] + " es el arbitro")
		w.Writing(texts[1])
	}
	return texts
}

func main() {
	if os.Args[1:] != nil {
		l := os.Args[1]
		listener := Listener{l}
		w := os.Args[2]
		writer := Writer{w}
		// 	//data:=
		Meeting(listener, writer)
	}
	fmt.Print(strings.TrimSuffix(strings.TrimPrefix("B1V1Hello, GophersU75O75", "B1"), "O75"))
}
