package protocol

import "strings"

//FOR WRITER//

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

//FOR LISTENER//

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

//Separate a text
func Separate(t string) []string {
	return strings.SplitAfter(t, "O75")
}
