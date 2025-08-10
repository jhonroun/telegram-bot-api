package tgbotapi

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test_56_MessageConfig_ProtectContent(t *testing.T) {
	cfg := MessageConfig{BaseChat: BaseChat{ChatID: 1, ProtectContent: true}, Text: "x"}
	p, err := cfg.params()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.method() != "sendMessage" {
		t.Fatal("method mismatch")
	}
	if p["protect_content"] != "true" {
		t.Fatalf("protect_content=%q", p["protect_content"])
	}
}

func Test_56_Copy_And_Forward_ProtectContent(t *testing.T) {
	copyCfg := CopyMessageConfig{BaseChat: BaseChat{ChatID: 1, ProtectContent: true}, FromChatID: 2, MessageID: 3}
	p, err := copyCfg.params()
	if err != nil {
		t.Fatal(err)
	}
	if copyCfg.method() != "copyMessage" {
		t.Fatal("method mismatch")
	}
	if p["protect_content"] != "true" {
		t.Fatal("copy: protect_content missing")
	}

	fwdCfg := ForwardConfig{BaseChat: BaseChat{ChatID: 1, ProtectContent: true}, FromChatID: 2, MessageID: 3}
	p, err = fwdCfg.params()
	if err != nil {
		t.Fatal(err)
	}
	if fwdCfg.method() != "forwardMessage" {
		t.Fatal("method mismatch")
	}
	if p["protect_content"] != "true" {
		t.Fatal("forward: protect_content missing")
	}
}

func Test_56_MessageEntity_Spoiler_JSON(t *testing.T) {
	const js = `{"type":"spoiler","offset":0,"length":5}`
	var e MessageEntity
	if err := json.Unmarshal([]byte(js), &e); err != nil {
		t.Fatal(err)
	}
	if e.Type != "spoiler" {
		t.Fatalf("type=%q", e.Type)
	}

	b, err := json.Marshal(e)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), `"spoiler"`) {
		t.Fatalf("marshal=%s", b)
	}
}

func Test_56_ProtectContent_MessageConfig(t *testing.T) {
	m := MessageConfig{BaseChat: BaseChat{ChatID: 1, ProtectContent: true}, Text: "x"}
	p, _ := m.params()
	if p["protect_content"] != "true" {
		t.Fatal("protect_content missing in sendMessage")
	}
}

func Test_56_ProtectContent_ForwardCopy(t *testing.T) {
	f := ForwardConfig{BaseChat: BaseChat{ChatID: 1, ProtectContent: true}, FromChatID: 2, MessageID: 3}
	pf, _ := f.params()
	if pf["protect_content"] != "true" {
		t.Fatal("protect_content missing in forwardMessage")
	}

	c := CopyMessageConfig{BaseChat: BaseChat{ChatID: 1, ProtectContent: true}, FromChatID: 2, MessageID: 3}
	pc, _ := c.params()
	if pc["protect_content"] != "true" {
		t.Fatal("protect_content missing in copyMessage")
	}
}

func Test_56_SendProtectedMessage(t *testing.T) {
	bot := getBot(t)

	msg := NewMessage(ChatID, "A secret message")
	msg.ProtectContent = true
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}

func Test_56_SendProtectedPhoto(t *testing.T) {
	bot := getBot(t)

	msg := NewPhoto(ChatID, FilePath("tests/image.jpg"))
	msg.Caption = "Test"
	msg.ProtectContent = true
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}

// Отправка сообщения со спойлером через entities (Bot API 5.6)
func Test_56_SendWithSpoilerEntity(t *testing.T) {
	bot := getBot(t)

	text := "secret with spoiler"
	ents := []MessageEntity{
		{Type: "spoiler", Offset: 0, Length: 6}, // "secret"
	}

	msg := NewMessage(ChatID, text)
	msg.Entities = ents

	_, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
}

// Отправка сообщения со спойлером через MarkdownV2: ||spoiler||
func Test_56_SendWithSpoilerMarkdownV2(t *testing.T) {
	bot := getBot(t)

	// MarkdownV2 для спойлера — оборачиваем текст в двойные пайпы
	msg := NewMessage(ChatID, "||secret|| with markdown spoiler")
	msg.ParseMode = "MarkdownV2"

	_, err := bot.Send(msg)
	if err != nil {
		t.Error(err)
	}
}
