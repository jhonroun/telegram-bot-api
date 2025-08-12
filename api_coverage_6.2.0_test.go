package tgbotapi

import (
	"encoding/json"
	"strings"
	"testing"
)

func Test62_MessageEntity_CustomEmoji_Unmarshal(t *testing.T) {
	raw := `{"type":"custom_emoji","offset":0,"length":2,"custom_emoji_id":"54321"}`
	var e MessageEntity
	if err := json.Unmarshal([]byte(raw), &e); err != nil {
		t.Fatal(err)
	}
	if e.Type != "custom_emoji" || e.CustomEmojiID != "54321" {
		t.Fatalf("bad entity: %+v", e)
	}
}

func Test62_GetCustomEmojiStickers_Params(t *testing.T) {
	cfg := GetCustomEmojiStickersConfig{CustomEmojiIDs: []string{"A", "B"}}
	p, err := cfg.params()
	if err != nil {
		t.Fatal(err)
	}
	if got := cfg.method(); got != "getCustomEmojiStickers" {
		t.Fatal("method mismatch")
	}
	v, ok := p["custom_emoji_ids"]
	if !ok {
		t.Fatal("missing ids")
	}
	var ids []string
	if err := json.Unmarshal([]byte(v), &ids); err != nil || len(ids) != 2 {
		t.Fatal("ids json bad")
	}
}

func Test62_StickerSet_StickerType_Unmarshal(t *testing.T) {
	raw := []byte(`{"title":"A","name":"a_by_bot","sticker_type":"custom_emoji"}`)
	var ss StickerSet
	if err := json.Unmarshal(raw, &ss); err != nil {
		t.Fatal(err)
	}
	if ss.StickerType != StickerTypeCustomEmoji {
		t.Fatalf("sticker_type mismatch: %q", ss.StickerType)
	}
}

func Test62_StickerSet_LegacyContainsMasks_Fallback(t *testing.T) {
	raw := []byte(`{"title":"B","name":"b_by_bot","contains_masks":true}`)
	var ss StickerSet
	if err := json.Unmarshal(raw, &ss); err != nil {
		t.Fatal(err)
	}
}

func Test62_StickerSet_DefaultRegular(t *testing.T) {
	raw := []byte(`{"title":"C","name":"c_by_bot"}`)
	var ss StickerSet
	if err := json.Unmarshal(raw, &ss); err != nil {
		t.Fatal(err)
	}
}

func Test62_Chat_HasRestrictedVoiceAndVideoMessages_Unmarshal(t *testing.T) {
	raw := []byte(`{"id":123,"type":"supergroup","has_restricted_voice_and_video_messages":true}`)
	var c Chat
	if err := json.Unmarshal(raw, &c); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if !c.HasRestrictedVoiceAndVideoMessages {
		t.Fatalf("expected true, got false")
	}
}

func Test62_Chat_HasRestrictedVoiceAndVideoMessages_DefaultFalse(t *testing.T) {
	// Поле отсутствует в JSON -> по умолчанию false
	raw := []byte(`{"id":123,"type":"supergroup"}`)
	var c Chat
	if err := json.Unmarshal(raw, &c); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if c.HasRestrictedVoiceAndVideoMessages {
		t.Fatalf("expected false by default")
	}
}

func Test62_Chat_HasRestrictedVoiceAndVideoMessages_MarshalOmitEmpty(t *testing.T) {
	c := Chat{ID: 123, Type: "supergroup"} // zero-value: false
	b, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	if got := string(b); strings.Contains(got, "has_restricted_voice_and_video_messages") {
		t.Fatalf("omitempty failed, got: %s", got)
	}
}

func Test62_Sticker_CustomEmoji_Unmarshal(t *testing.T) {
	raw := []byte(`{"type":"custom_emoji","custom_emoji_id":"CE_123"}`)
	var s Sticker
	if err := json.Unmarshal(raw, &s); err != nil {
		t.Fatal(err)
	}
	if s.Type != StickerTypeCustomEmoji {
		t.Fatalf("type mismatch: %q", s.Type)
	}
	if s.CustomEmojiID != "CE_123" {
		t.Fatalf("custom_emoji_id mismatch: %q", s.CustomEmojiID)
	}
	if !s.IsCustomEmoji() {
		t.Fatalf("IsCustomEmoji expected true")
	}
}

func Test62_Sticker_Regular_NoCustomID(t *testing.T) {
	raw := []byte(`{"type":"regular"}`)
	var s Sticker
	if err := json.Unmarshal(raw, &s); err != nil {
		t.Fatal(err)
	}
	if s.Type != StickerTypeRegular {
		t.Fatalf("type mismatch: %q", s.Type)
	}
	if s.CustomEmojiID != "" {
		t.Fatalf("custom_emoji_id must be empty, got %q", s.CustomEmojiID)
	}
}

func Test62_Sticker_Marshal_OmitEmpty(t *testing.T) {
	s := Sticker{Type: StickerTypeRegular}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	js := string(b)
	if len(js) == 0 || js == "{}" {
		// ok: может не быть других полей
	}
	// Проверим, что пустого custom_emoji_id нет в JSON
	if strings.Contains(js, `"custom_emoji_id"`) {
		t.Fatalf("omitempty failed, got: %s", js)
	}
}

func Test62_GetCustomEmojiStickers_Params_OK(t *testing.T) {
	cfg := NewGetCustomEmojiStickersConfig("ID1", "ID2")

	if got, want := cfg.method(), "getCustomEmojiStickers"; got != want {
		t.Fatalf("method mismatch: got %q want %q", got, want)
	}

	p, err := cfg.params()
	if err != nil {
		t.Fatalf("params error: %v", err)
	}

	raw, ok := p["custom_emoji_ids"]
	if !ok {
		t.Fatal("custom_emoji_ids missing")
	}

	var ids []string
	if err := json.Unmarshal([]byte(raw), &ids); err != nil {
		t.Fatalf("ids not valid JSON: %v", err)
	}
	if len(ids) != 2 || ids[0] != "ID1" || ids[1] != "ID2" {
		t.Fatalf("ids content mismatch: %+v", ids)
	}
}

func Test62_GetCustomEmojiStickers_Params_Empty(t *testing.T) {
	cfg := GetCustomEmojiStickersConfig{}
	if _, err := cfg.params(); err == nil {
		t.Fatal("expected error on empty custom_emoji_ids")
	}
}

func Test62_GetCustomEmojiStickers_Params_TooMany(t *testing.T) {
	var ids []string
	for i := 0; i < 201; i++ {
		ids = append(ids, "X")
	}
	cfg := GetCustomEmojiStickersConfig{CustomEmojiIDs: ids}
	if _, err := cfg.params(); err == nil {
		t.Fatal("expected error when >200 ids")
	}
}

func Test62_GetCustomEmojiStickers_Response_Unmarshal(t *testing.T) {
	raw := []byte(`{
		"ok": true,
		"result": [
			{"type":"custom_emoji","custom_emoji_id":"CE_1"},
			{"type":"regular"}
		]
	}`)
	var resp struct {
		Ok     bool      `json:"ok"`
		Result []Sticker `json:"result"`
	}
	if err := json.Unmarshal(raw, &resp); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if !resp.Ok || len(resp.Result) != 2 {
		t.Fatalf("bad payload: %+v", resp)
	}
	if resp.Result[0].CustomEmojiID != "CE_1" {
		t.Fatalf("first sticker mismatch: %+v", resp.Result[0])
	}
	if resp.Result[0].Type != StickerTypeCustomEmoji || resp.Result[0].CustomEmojiID != "CE_1" {
		t.Fatalf("first sticker mismatch: %+v", resp.Result[0])
	}
	if resp.Result[1].Type != StickerTypeRegular {
		t.Fatalf("second sticker type mismatch: %+v", resp.Result[1])
	}
}
