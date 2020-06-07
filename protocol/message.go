package protocol

func PutLimiter(m []string) []string {
	var messages []string
	for _, message := range m {
		messages = append(messages, "B1"+message+"O75")
	}
	return messages
}
func Unite(m []string) string {
	var messages string
	for _, message := range m {
		messages += message
	}
	return messages
}
