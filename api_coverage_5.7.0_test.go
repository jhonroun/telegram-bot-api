package tgbotapi

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"
)

// --- Типы ---

func Test_57_Sticker_IsVideo_JSON(t *testing.T) {
	const js = `{"is_video":true}`
	var s Sticker
	if err := json.Unmarshal([]byte(js), &s); err != nil {
		t.Fatal(err)
	}
	if !s.IsVideo {
		t.Fatal("Sticker.IsVideo expected true")
	}

	b, _ := json.Marshal(s)
	if !strings.Contains(string(b), `"is_video"`) {
		t.Fatalf("marshal missing is_video: %s", b)
	}
}

func Test_57_StickerSet_IsVideo_JSON(t *testing.T) {
	const js = `{"is_video":true}`
	var ss StickerSet
	if err := json.Unmarshal([]byte(js), &ss); err != nil {
		t.Fatal(err)
	}
	if !ss.IsVideo {
		t.Fatal("StickerSet.IsVideo expected true")
	}
}

// --- Конфиги ---

func Test_57_CreateNewStickerSet_WebM_FileKey(t *testing.T) {
	c := NewStickerSetConfig{
		UserID:      1,
		Name:        "pack_by_bot",
		Title:       "Pack",
		Emojis:      "😀",
		WebMSticker: FilePath("tests/1347045309.webm"),
	}
	files := c.files()
	if len(files) != 1 {
		t.Fatalf("files len=%d, want 1", len(files))
	}
	if files[0].Name != "webm_sticker" {
		t.Fatalf("file key=%q, want webm_sticker", files[0].Name)
	}
}

func Test_57_CreateNewStickerSet_TGS_FileKey(t *testing.T) {
	c := NewStickerSetConfig{
		UserID:     1,
		Name:       "pack_by_bot",
		Title:      "Pack",
		Emojis:     "😀",
		TGSSticker: FilePath("tests/1083673963.tgs"), // <-- TGSSticker, не WebMSticker
	}
	files := c.files()
	if len(files) != 1 {
		t.Fatalf("files len=%d, want 1", len(files))
	}
	if files[0].Name != "tgs_sticker" {
		t.Fatalf("file key=%q, want tgs_sticker", files[0].Name)
	}
}

func Test_57_AddSticker_WebM_FileKey(t *testing.T) {
	c := AddStickerConfig{
		UserID:      1,
		Name:        "pack_by_bot",
		Emojis:      "😀",
		WebMSticker: FilePath("tests/1347045309.webm"),
	}
	files := c.files()
	if len(files) != 1 || files[0].Name != "webm_sticker" {
		t.Fatalf("bad files: %#v", files)
	}
}

// Создаёт новый набор c WEBM-стикером, получает set обратно и шлёт первый стикер в чат.
func Test_57_CreateWebMStickerAndSend(t *testing.T) {
	bot := getBot(t)

	// Имя набора должно оканчиваться на _by_<botusername>, только [a-z0-9_]
	username := strings.ToLower(bot.Self.UserName)
	name := fmt.Sprintf("api57_%d_by_%s", time.Now().Unix(), username)
	title := fmt.Sprintf("API57 %d", time.Now().Unix())

	// 1) createNewStickerSet (webm_sticker)
	create := NewStickerSetConfig{
		UserID:      ChatID,
		Name:        name,
		Title:       title,
		Emojis:      "😀",
		WebMSticker: FilePath("tests/1347045309.webm"),
	}
	if _, err := bot.Request(create); err != nil {
		t.Fatalf("createNewStickerSet: %v", err)
	}

	// 2) getStickerSet → берём file_id первого стикера
	set, err := bot.GetStickerSet(GetStickerSetConfig{Name: name})
	if err != nil {
		t.Fatalf("getStickerSet: %v", err)
	}
	if len(set.Stickers) == 0 {
		t.Fatalf("sticker set %s is empty", name)
	}
	fileID := set.Stickers[0].FileID
	if fileID == "" {
		t.Fatalf("empty file_id in sticker")
	}

	// 3) sendSticker в тестовый чат
	st := NewSticker(ChatID, FileID(fileID))
	if _, err := bot.Send(st); err != nil {
		t.Fatalf("sendSticker: %v", err)
	}
}
