package tgbotapi

import (
	"encoding/json"
	"testing"
)

const msgJSON = `{"is_automatic_forward":true,"has_protected_content":true}`
const chatJSON = `{"has_private_forwards":true,"has_protected_content":true}`

func Test_BotAPI55_MessageFields_JSON(t *testing.T) {
	var m Message
	if err := json.Unmarshal([]byte(msgJSON), &m); err != nil {
		t.Fatalf("unmarshal Message: %v", err)
	}
	if !m.IsAutomaticForward {
		t.Fatal("Message.IsAutomaticForward should be true after unmarshal")
	}
	if !m.HasProtectedContent {
		t.Fatal("Message.HasProtectedContent should be true after unmarshal")
	}
}

func Test_BotAPI55_ChatFields_JSON(t *testing.T) {
	var ch Chat
	if err := json.Unmarshal([]byte(chatJSON), &ch); err != nil {
		t.Fatalf("unmarshal Chat: %v", err)
	}
	if !ch.HasPrivateForwards {
		t.Fatal("Chat.HasPrivateForwards should be true after unmarshal")
	}
	if !ch.HasProtectedContent {
		t.Fatal("Chat.HasProtectedContent should be true after unmarshal")
	}
}
