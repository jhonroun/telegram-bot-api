package tgbotapi

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test60_Message_VideoChat_JSON(t *testing.T) {
	const jsNew = `{"video_chat_scheduled":{"start_date":1},"video_chat_started":{},"video_chat_ended":{"duration":2},"video_chat_participants_invited":{"users":[]}}`
	var m Message
	if err := json.Unmarshal([]byte(jsNew), &m); err != nil {
		t.Fatal(err)
	}
	if m.VideoChatScheduled == nil || m.VideoChatStarted == nil || m.VideoChatEnded == nil || m.VideoChatParticipantsInvited == nil {
		t.Fatal("video_chat_* not filled from JSON")
	}
}

func Test60_PromoteChatMember_VideoChats_Param(t *testing.T) {
	c := PromoteChatMemberConfig{ChatMemberConfig: ChatMemberConfig{ChatID: ChatID, UserID: 42},
		CanManageVideoChats: true,
	}
	p, err := c.params()
	if err != nil {
		t.Fatal(err)
	}
	if c.method() != "promoteChatMember" {
		t.Fatal("method mismatch")
	}
	if p["can_manage_video_chats"] != "true" {
		t.Fatalf("want can_manage_video_chats=true, got: %#v", p["can_manage_video_chats"])
	}
	if _, ok := p["can_manage_voice_chats"]; ok {
		t.Fatal("must not send legacy can_manage_voice_chats when video flag is set")
	}
}

func Test60_ChatMemberAdmin_VideoChats_JSON(t *testing.T) {
	const js = `{"status":"administrator","user":{"id":1},"can_manage_video_chats":true}`
	var m ChatAdministratorRights
	if err := json.Unmarshal([]byte(js), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !m.CanManageVideoChats {
		t.Fatal("CanManageVideoChats expected true")
	}
}

func Test60_WebhookInfo_LastSyncError_JSON(t *testing.T) {
	const js = `{"last_synchronization_error_date":1660000000}`
	var w WebhookInfo
	if err := json.Unmarshal([]byte(js), &w); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if w.LastSynchronizationErrorDate != 1660000000 {
		t.Fatalf("want 1660000000, got %d", w.LastSynchronizationErrorDate)
	}
}

func Test60_ChatAdministratorRights_JSON(t *testing.T) {
	const js = `{"can_manage_chat":true,"can_manage_video_chats":true,"can_invite_users":true}`
	var r ChatAdministratorRights
	if err := json.Unmarshal([]byte(js), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !r.CanManageChat || !r.CanManageVideoChats || !r.CanInviteUsers {
		t.Fatal("rights fields not populated from JSON")
	}
}

func Test60_SetMyDefaultAdministratorRights_Params(t *testing.T) {
	cfg := SetMyDefaultAdministratorRightsConfig{
		Rights: ChatAdministratorRights{
			CanManageChat:       true,
			CanManageVideoChats: true,
			CanInviteUsers:      true,
		},
		ForChannels: true,
	}
	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params: %v", err)
	}
	if cfg.method() != "setMyDefaultAdministratorRights" {
		t.Fatalf("method=%q", cfg.method())
	}
	// "rights" должен присутствовать и содержать JSON с can_manage_video_chats
	rights, ok := p["rights"]
	if !ok || !strings.Contains(rights, "can_manage_video_chats") {
		t.Fatalf("rights missing or invalid: %v", rights)
	}
	if p["for_channels"] != "true" {
		t.Fatalf("for_channels=%q want true", p["for_channels"])
	}
}

func Test60_GetMyDefaultAdministratorRights_Params(t *testing.T) {
	cfg := GetMyDefaultAdministratorRightsConfig{ForChannels: true}
	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params: %v", err)
	}
	if cfg.method() != "getMyDefaultAdministratorRights" {
		t.Fatalf("method=%q", cfg.method())
	}
	if p["for_channels"] != "true" {
		t.Fatalf("for_channels=%q want true", p["for_channels"])
	}
}

func Test60_MenuButton_WebApp_JSON(t *testing.T) {
	btn := NewMenuButtonWebApp("Open", "https://example.com")
	b, err := json.Marshal(btn)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"type":"web_app"`) || !strings.Contains(s, `"web_app"`) {
		t.Fatalf("bad json: %s", s)
	}
}

func Test60_SetChatMenuButton_Params_WebApp(t *testing.T) {
	cfg := SetChatMenuButtonConfig{
		ChatID:     0, // глобально (для приватных чатов)
		MenuButton: NewMenuButtonWebApp("Open", "https://example.com"),
	}
	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params: %v", err)
	}
	if cfg.method() != "setChatMenuButton" {
		t.Fatalf("method=%q", cfg.method())
	}
	v, ok := p["menu_button"]
	if !ok {
		t.Fatal("menu_button missing")
	}
	if !strings.Contains(v, `"type":"web_app"`) || !strings.Contains(v, `"url":"https://example.com"`) {
		t.Fatalf("menu_button json bad: %s", v)
	}
	if _, has := p["chat_id"]; has {
		t.Fatalf("chat_id must be omitted when zero, got %q", p["chat_id"])
	}
}

func Test60_GetChatMenuButton_Params_WithChat(t *testing.T) {
	cfg := GetChatMenuButtonConfig{ChatID: 123}
	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params: %v", err)
	}
	if cfg.method() != "getChatMenuButton" {
		t.Fatalf("method=%q", cfg.method())
	}
	if p["chat_id"] != "123" {
		t.Fatalf("chat_id=%q", p["chat_id"])
	}
}

func Test60_Message_WebAppData_JSON_Unmarshal(t *testing.T) {
	const js = `{"web_app_data":{"data":"PAYLOAD","button_text":"Open"}}`
	var m Message
	if err := json.Unmarshal([]byte(js), &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if m.WebAppData == nil {
		t.Fatal("WebAppData is nil")
	}
	if m.WebAppData.Data != "PAYLOAD" || m.WebAppData.ButtonText != "Open" {
		t.Fatalf("unexpected WebAppData: %+v", *m.WebAppData)
	}
}

func Test60_Message_WebAppData_JSON_Marshal(t *testing.T) {
	m := Message{WebAppData: &WebAppData{Data: "X", ButtonText: "Open"}}
	b, err := json.Marshal(m)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"web_app_data"`) ||
		!strings.Contains(s, `"data":"X"`) ||
		!strings.Contains(s, `"button_text":"Open"`) {
		t.Fatalf("bad json: %s", s)
	}
}

func Test60_AnswerWebAppQuery_Params(t *testing.T) {
	res := NewInlineQueryResultArticle("1", "ok", "hi from web app")
	res.InputMessageContent = InputTextMessageContent{
		Text: "hi from web app",
	}

	cfg := AnswerWebAppQueryConfig{
		WebAppQueryID: "abc123",
		Result:        res,
	}

	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params(): %v", err)
	}
	if got, want := cfg.method(), "answerWebAppQuery"; got != want {
		t.Fatalf("method=%q want %q", got, want)
	}
	if p["web_app_query_id"] != "abc123" {
		t.Fatalf("web_app_query_id=%q", p["web_app_query_id"])
	}
	// result должен быть JSON-строкой с сериализованным InlineQueryResult
	if v, ok := p["result"]; !ok || !strings.Contains(v, `"type":"article"`) || !strings.Contains(v, `"message_text"`) {
		t.Fatalf("bad result: %v", p["result"])
	}
}

// 2) Проверяем (де)сериализацию ответа SentWebAppMessage
func Test60_SentWebAppMessage_JSON(t *testing.T) {
	const js = `{"inline_message_id":"AAQCDEF"}`
	var s SentWebAppMessage
	if err := json.Unmarshal([]byte(js), &s); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if s.InlineMessageID != "AAQCDEF" {
		t.Fatalf("inline_message_id=%q", s.InlineMessageID)
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(string(b), `"inline_message_id"`) {
		t.Fatalf("marshal missing field: %s", string(b))
	}
}

func Test60_InlineButton_WebApp_JSON(t *testing.T) {
	btn := NewInlineKeyboardButtonWebApp("Open", WebAppInfo{URL: "https://example.com/app"})
	b, err := json.Marshal(btn)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"type":"web_app"`) && !strings.Contains(s, `"web_app"`) {
		t.Fatalf("web_app missing: %s", s)
	}
	if !strings.Contains(s, `"url":"https://example.com/app"`) {
		t.Fatalf("url missing: %s", s)
	}
}

func Test60_KeyboardButton_WebApp_JSON(t *testing.T) {
	btn := NewKeyboardButtonWebApp("Open", WebAppInfo{URL: "https://example.com/app"})
	b, err := json.Marshal(btn)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	s := string(b)
	if !strings.Contains(s, `"web_app"`) || !strings.Contains(s, `"url":"https://example.com/app"`) {
		t.Fatalf("bad json: %s", s)
	}
}
