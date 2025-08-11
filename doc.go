// #Golang bindings for the Telegram Bot API
//
// [![Go Reference](https://pkg.go.dev/badge/github.com/jhonroun/telegram-bot-api/v6.svg)](https://pkg.go.dev/github.com/jhonroun/telegram-bot-api/v6)
// [![Test](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml/badge.svg)](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml)
//
// Package tgbotapi — Go-client Telegram Bot API.
// Support: sending messages, inline-mode, webhooks, etc.
// Actual Bot API Version: 6.2.0
//
// The repo was created to study and check the relevance of the module for working with the Bot API (https://github.com/go-telegram-bot-api/telegram-bot-api), which is called step-by-step. Many thanks to the author for the awesome experience and idea.
// Initially, I wanted to create a tool for writing modern bots. But in the process of adding functionality, I thought
// that I was writing it for myself first and foremost. There are quite enough forms with an updated version of the Bot API on github.com.
//
// From now on, the abandonment of versioning like v0.*/v1.*
//
// Always the latest version on release:
//
//	`go get github.com/jhonroun/telegram-bot-api@latest`
//
// Fast start:
//
//	bot, _ := tgbotapi.NewBotAPI("TOKEN")
//	bot.Send(tgbotapi.NewMessage(chatID, "hi"))
package tgbotapi
