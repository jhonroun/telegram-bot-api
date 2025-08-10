//go:build integration

package tgbotapi

import (
	"os"
	"testing"
)

func getBot(t *testing.T) *BotAPI {
	t.Helper()
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		t.Skip("TELEGRAM_TOKEN not set; skipping integration tests")
	}
	bot, err := NewBotAPI(token)
	if err != nil {
		t.Fatalf("NewBotAPI: %v", err)
	}
	return bot
}

func TestGetUpdatesIntegration(t *testing.T) {
	bot := getBot(t)
	_, err := bot.GetUpdatesChan(UpdateConfig{Timeout: 1})
	if err != nil {
		t.Fatalf("GetUpdatesChan: %v", err)
	}
}
