package main

import (
	"fmt"
	"strings"
)

//Separate a text
func Separate(t string) []string {
	return strings.SplitAfter(t, "O75")
}

//CutIndexAndSuffix of a text
func CutIndexAndSuffix(t string) string {
	return strings.TrimSuffix(strings.TrimPrefix(t, "B1"), "O75")
}

//CutLimiter of a message.
func CutLimiter(m []string) []string {
	var messages []string
	for _, message := range m {
		messages = append(messages, CutIndexAndSuffix(message))
	}
	return messages
}

//PutLimiter .
func PutLimiter(m []string) []string {
	var messages []string
	for _, message := range m {
		messages = append(messages, "B1"+message+"O75")
	}
	return messages
}

//Unite .
func Unite(m []string) string {
	var messages string
	for _, message := range m {
		messages += message
	}
	return messages
}

func main() {
	// if os.Args[1:] != nil {
	// 	listener := NewListener(os.Args[1])
	// 	writer := NewWriter(os.Args[2])
	// 	// 	//data:=
	// 	Meeting(listener, writer)
	// }
	var saludo []string
	saludo = append(saludo, "COM1sayV1")
	saludo = append(saludo, "COM1sayU75")
	writeSaludo := Unite(PutLimiter(saludo))
	fmt.Println(writeSaludo)
	listenSaludo := CutLimiter(Separate(writeSaludo))
	fmt.Println(listenSaludo)
}
