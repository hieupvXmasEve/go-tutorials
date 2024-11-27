package slices

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line
// Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.

func filterMessages(messages []Message, filterType string) []Message {
	var filteredMessages []Message
	for _, message := range messages {
		if message.Type() == filterType {
			filteredMessages = append(filteredMessages, message)
		}
	}
	return filteredMessages
}

func messageFilter() {
	messages := []Message{
		TextMessage{"Alice", "Hello, world!"},
		MediaMessage{Sender: "Bob", MediaType: "image", Content: "https://example.com/image.jpg"},
		LinkMessage{"Charlie", "https://example.com", "Check out this link!"},
	}

	filteredMessages := filterMessages(messages, "media")
	fmt.Println(filteredMessages)
}
