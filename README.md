# Golang bindings for the Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/jhonroun/telegram-bot-api/v6.svg)](https://pkg.go.dev/github.com/jhonroun/telegram-bot-api/v6)
[![Test](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml/badge.svg)](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml)

Рекомендуем закрепиться на релизе:

```
go get github.com/jhonroun/telegram-bot-api/v6@v5.7.0
```

## Changelog

### v5.7.0 — 2025-08-10
Поддержка Telegram **Bot API 5.7** (видеостикеры).

#### Added
- **Video stickers:**
  - Новые поля:
    - `Sticker.IsVideo`
    - `StickerSet.IsVideo`
  - Новые входные данные для стикер-методов:
    - `webm_sticker` в `NewStickerSetConfig` и `addStickerToSet`
    - В библиотеке: `NewStickerSetConfig.WebMSticker`, `AddStickerConfig.WebMSticker`
- Валидатор конфигов: требуется **ровно один** из `png_sticker` / `tgs_sticker` / `webm_sticker`.

#### Examples
```go
	// Создать новый набор с WEBM-стикером
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
_, _ = bot.Request(c)

### v5.6.0 — 2025-08-10
Поддержка Telegram **Bot API 5.5–5.6** и сопутствующие правки библиотеки.

#### Added
- **Поля (Bot API 5.5):**
  - `Message.IsAutomaticForward`
  - `Message.HasProtectedContent`
  - `Chat.HasProtectedContent`
  - `Chat.HasPrivateForwards`
- **Методы (Bot API 5.5):**
  - `banChatSenderChat` / `unbanChatSenderChat`  
    В библиотеке: `BanChatSenderChatConfig` / `UnbanChatSenderChatConfig`, конструкторы `NewBanChatSenderChat` и `NewUnbanChatSenderChat`.
- **Параметр (Bot API 5.6):**
  - `protect_content` для всех `send*` методов, а также `copyMessage` и `forwardMessage`.  
    В библиотеке: поле `ProtectContent bool` в соответствующих `*Config`.
- **Сущности (Bot API 5.6):**
  - `MessageEntity{Type: "spoiler"}` — спойлер-сущность поддерживается при (де)сериализации.

#### Examples
```go
// Protect content (нельзя переслать/сохранить)
msg := tgbotapi.NewMessage(chatID, "секрет")
msg.ProtectContent = true
_, _ = bot.Send(msg)

// Бан/разбан channel chat в супергруппе/канале
_, _ = bot.Request(tgbotapi.NewBanChatSenderChat(supergroupID, senderChannelID))
_, _ = bot.Request(tgbotapi.NewUnbanChatSenderChat(supergroupID, senderChannelID))

// Spoiler entity (через сущности)
entities := []tgbotapi.MessageEntity{{Type: "spoiler", Offset: 0, Length: 6}}
m := tgbotapi.NewMessage(chatID, "спойлер")
m.Entities = &entities
_, _ = bot.Send(m)
```


## Example

First, ensure the library is installed and up to date by running
`go get -u github.com/jhonroun/telegram-bot-api/v6`.

This is a very simple bot that just displays any gotten updates,
then replies it to that chat.

```go
package main

import (
	"log"

	tgbotapi "github.com/jhonroun/telegram-bot-api/v6"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
```

Tests
Добавлены unit-тесты без сети на поля (Message, Chat) и на параметры/методы Ban/UnbanChatSenderChat и ProtectContent.

Migration notes: Ломающих изменений нет.
