package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// It should loop through each message and set the tags to the result of the passed in function.
	for i := range messages {
		messages[i].tags = tagger(messages[i])
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	// Convert content to lowercase for case-insensitive checking
	lowercaseContent := strings.ToLower(msg.content)

	// Check for "urgent" first (case-insensitive)
	if strings.Contains(lowercaseContent, "urgent") {
		tags = append(tags, "Urgent")
	}

	// Check for "sale" second (case-insensitive)
	if strings.Contains(lowercaseContent, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

func messagesTager() {
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see sale!"},
		{id: "002", content: "Big sale on all items!"},
		// Additional messages...
	}
	taggedMessages := tagMessages(messages, tagger)
	fmt.Println(taggedMessages)
}
