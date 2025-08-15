package tgbotapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// NewMessage creates a new Message.
//
// chatID is where to send it, text is the message text.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewMessage(chatID any, text string) MessageConfig {
	toID := getChatID(chatID)
	return MessageConfig{
		BaseChat: BaseChat{
			ChatID:           toID,
			ReplyToMessageID: 0,
		},
		Text:                  text,
		DisableWebPagePreview: false,
	}
}

// NewDeleteMessage creates a request to delete a message.
func NewDeleteMessage(chatID int64, messageID int) DeleteMessageConfig {
	return DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}
}

// NewMessageToChannel creates a new Message that is sent to a channel
// by username.
//
// username is the username of the channel, text is the message text,
// and the username should be in the form of `@username`.
// Or you can use bot.GetUserIDbyUsername("@username") and simple NewMessage
func NewMessageToChannel(username string, text string) MessageConfig {
	return MessageConfig{
		BaseChat: BaseChat{
			ChannelUsername: username,
		},
		Text: text,
	}
}

// NewForward creates a new forward.
//
// chatID is where to send it, fromChatID is the source chat,
// and messageID is the ID of the original message.
func NewForward(chatID int64, fromChatID int64, messageID int) ForwardConfig {
	return ForwardConfig{
		BaseChat:   BaseChat{ChatID: chatID},
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

// NewCopyMessage creates a new copy message.
//
// chatID is where to send it, fromChatID is the source chat,
// and messageID is the ID of the original message.
func NewCopyMessage(chatID int64, fromChatID int64, messageID int) CopyMessageConfig {
	return CopyMessageConfig{
		BaseChat:   BaseChat{ChatID: chatID},
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

// NewPhoto creates a new sendPhoto request.
//
// chatID is where to send it, file is a string path to the file,
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
//
// FileReader, or FileBytes.
//
// Note that you must send animated GIFs as a document.
func NewPhoto(chatID any, file RequestFileData) PhotoConfig {
	toID := getChatID(chatID)
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewPhotoToChannel creates a new photo uploader to send a photo to a channel.
//
// Note that you must send animated GIFs as a document.
// Or you can use bot.GetUserIDbyUsername("@username") and simple NewMessage
func NewPhotoToChannel(username string, file RequestFileData) PhotoConfig {
	return PhotoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{
				ChannelUsername: username,
			},
			File: file,
		},
	}
}

// NewAudio creates a new sendAudio request.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewAudio(chatID any, file RequestFileData) AudioConfig {
	toID := getChatID(chatID)
	return AudioConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewDocument creates a new sendDocument request.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewDocument(chatID any, file RequestFileData) DocumentConfig {
	toID := getChatID(chatID)
	return DocumentConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewSticker creates a new sendSticker request.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewSticker(chatID any, file RequestFileData) StickerConfig {
	toID := getChatID(chatID)
	return StickerConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewVideo creates a new sendVideo request.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewVideo(chatID any, file RequestFileData) VideoConfig {
	toID := getChatID(chatID)
	return VideoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewAnimation creates a new sendAnimation request.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewAnimation(chatID any, file RequestFileData) AnimationConfig {
	toID := getChatID(chatID)
	return AnimationConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewVideoNote creates a new sendVideoNote request.
//
// # Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
//
// chatID is where to send it, file is a string path to the file,
// FileReader, or FileBytes.
func NewVideoNote(chatID any, length int, file RequestFileData) VideoNoteConfig {
	toID := getChatID(chatID)
	return VideoNoteConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
		Length: length,
	}
}

// NewVoice creates a new sendVoice request.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewVoice(chatID any, file RequestFileData) VoiceConfig {
	toID := getChatID(chatID)
	return VoiceConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{ChatID: toID},
			File:     file,
		},
	}
}

// NewMediaGroup creates a new media group. Files should be an array of
// two to ten InputMediaPhoto or InputMediaVideo.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewMediaGroup(chatID any, files []interface{}) MediaGroupConfig {
	toID := getChatID(chatID)
	return MediaGroupConfig{
		ChatID: toID,
		Media:  files,
	}
}

// NewInputMediaPhoto creates a new InputMediaPhoto.
func NewInputMediaPhoto(media RequestFileData) InputMediaPhoto {
	return InputMediaPhoto{
		BaseInputMedia: BaseInputMedia{
			Type:  "photo",
			Media: media,
		},
		HasSpoiler: false,
	}
}

// NewInputMediaVideo creates a new InputMediaVideo.
func NewInputMediaVideo(media RequestFileData) InputMediaVideo {
	return InputMediaVideo{
		BaseInputMedia: BaseInputMedia{
			Type:  "video",
			Media: media,
		},
	}
}

// NewInputMediaAnimation creates a new InputMediaAnimation.
func NewInputMediaAnimation(media RequestFileData) InputMediaAnimation {
	return InputMediaAnimation{
		BaseInputMedia: BaseInputMedia{
			Type:  "animation",
			Media: media,
		},
	}
}

// NewInputMediaAudio creates a new InputMediaAudio.
func NewInputMediaAudio(media RequestFileData) InputMediaAudio {
	return InputMediaAudio{
		BaseInputMedia: BaseInputMedia{
			Type:  "audio",
			Media: media,
		},
	}
}

// NewInputMediaDocument creates a new InputMediaDocument.
func NewInputMediaDocument(media RequestFileData) InputMediaDocument {
	return InputMediaDocument{
		BaseInputMedia: BaseInputMedia{
			Type:  "document",
			Media: media,
		},
	}
}

// NewContact allows you to send a shared contact.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewContact(chatID any, phoneNumber, firstName string) ContactConfig {
	toID := getChatID(chatID)
	return ContactConfig{
		BaseChat: BaseChat{
			ChatID: toID,
		},
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// NewLocation shares your location.
//
// chatID is where to send it, latitude and longitude are coordinates.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewLocation(chatID any, latitude float64, longitude float64) LocationConfig {
	toID := getChatID(chatID)
	return LocationConfig{
		BaseChat: BaseChat{
			ChatID: toID,
		},
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewVenue allows you to send a venue and its location.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewVenue(chatID any, title, address string, latitude, longitude float64) VenueConfig {
	toID := getChatID(chatID)
	return VenueConfig{
		BaseChat: BaseChat{
			ChatID: toID,
		},
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewChatAction sets a chat action.
// Actions last for 5 seconds, or until your next action.
//
// chatID is where to send it, action should be set via ChatAction constants.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewChatAction(chatID any, action ChatAction) ChatActionConfig {
	toID := getChatID(chatID)
	return ChatActionConfig{
		BaseChat: BaseChat{ChatID: toID},
		Action:   action,
	}
}

// NewUserProfilePhotos gets user profile photos.
//
// userID is the ID of the user you wish to get profile photos from.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
func NewUserProfilePhotos(userID any) UserProfilePhotosConfig {
	toID := getChatID(userID)
	return UserProfilePhotosConfig{
		UserID: toID,
		Offset: 0,
		Limit:  0,
	}
}

// NewUpdate gets updates since the last Offset.
//
// offset is the last Update ID to include.
// You likely want to set this to the last Update ID plus 1.
func NewUpdate(offset int) UpdateConfig {
	return UpdateConfig{
		Offset:  offset,
		Limit:   0,
		Timeout: 0,
	}
}

// NewWebhook creates a new webhook.
//
// link is the url parsable link you wish to get the updates.
func NewWebhook(link string) (WebhookConfig, error) {
	u, err := url.Parse(link)

	if err != nil {
		return WebhookConfig{}, err
	}

	return WebhookConfig{
		URL: u,
	}, nil
}

// NewWebhookWithCert creates a new webhook with a certificate.
//
// link is the url you wish to get webhooks,
// file contains a string to a file, FileReader, or FileBytes.
func NewWebhookWithCert(link string, file RequestFileData) (WebhookConfig, error) {
	u, err := url.Parse(link)

	if err != nil {
		return WebhookConfig{}, err
	}

	return WebhookConfig{
		URL:         u,
		Certificate: file,
	}, nil
}

// NewInlineQueryResultArticle creates a new inline query article.
func NewInlineQueryResultArticle(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text: messageText,
		},
	}
}

// NewInlineQueryResultArticleMarkdown creates a new inline query article with Markdown parsing.
func NewInlineQueryResultArticleMarkdown(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text:      messageText,
			ParseMode: "Markdown",
		},
	}
}

// NewInlineQueryResultArticleMarkdownV2 creates a new inline query article with MarkdownV2 parsing.
func NewInlineQueryResultArticleMarkdownV2(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text:      messageText,
			ParseMode: "MarkdownV2",
		},
	}
}

// NewInlineQueryResultArticleHTML creates a new inline query article with HTML parsing.
func NewInlineQueryResultArticleHTML(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		Type:  "article",
		ID:    id,
		Title: title,
		InputMessageContent: InputTextMessageContent{
			Text:      messageText,
			ParseMode: "HTML",
		},
	}
}

// NewInlineQueryResultGIF creates a new inline query GIF.
func NewInlineQueryResultGIF(id, url string) InlineQueryResultGIF {
	return InlineQueryResultGIF{
		Type: "gif",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedGIF create a new inline query with cached photo.
func NewInlineQueryResultCachedGIF(id, gifID string) InlineQueryResultCachedGIF {
	return InlineQueryResultCachedGIF{
		Type:  "gif",
		ID:    id,
		GIFID: gifID,
	}
}

// NewInlineQueryResultMPEG4GIF creates a new inline query MPEG4 GIF.
func NewInlineQueryResultMPEG4GIF(id, url string) InlineQueryResultMPEG4GIF {
	return InlineQueryResultMPEG4GIF{
		Type: "mpeg4_gif",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedMPEG4GIF create a new inline query with cached MPEG4 GIF.
func NewInlineQueryResultCachedMPEG4GIF(id, MPEG4GIFID string) InlineQueryResultCachedMPEG4GIF {
	return InlineQueryResultCachedMPEG4GIF{
		Type:        "mpeg4_gif",
		ID:          id,
		MPEG4FileID: MPEG4GIFID,
	}
}

// NewInlineQueryResultPhoto creates a new inline query photo.
func NewInlineQueryResultPhoto(id, url string) InlineQueryResultPhoto {
	return InlineQueryResultPhoto{
		Type: "photo",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultPhotoWithThumb creates a new inline query photo.
func NewInlineQueryResultPhotoWithThumb(id, url, thumb string) InlineQueryResultPhoto {
	return InlineQueryResultPhoto{
		Type:     "photo",
		ID:       id,
		URL:      url,
		ThumbURL: thumb,
	}
}

// NewInlineQueryResultCachedPhoto create a new inline query with cached photo.
func NewInlineQueryResultCachedPhoto(id, photoID string) InlineQueryResultCachedPhoto {
	return InlineQueryResultCachedPhoto{
		Type:    "photo",
		ID:      id,
		PhotoID: photoID,
	}
}

// NewInlineQueryResultVideo creates a new inline query video.
func NewInlineQueryResultVideo(id, url string) InlineQueryResultVideo {
	return InlineQueryResultVideo{
		Type: "video",
		ID:   id,
		URL:  url,
	}
}

// NewInlineQueryResultCachedVideo create a new inline query with cached video.
func NewInlineQueryResultCachedVideo(id, videoID, title string) InlineQueryResultCachedVideo {
	return InlineQueryResultCachedVideo{
		Type:    "video",
		ID:      id,
		VideoID: videoID,
		Title:   title,
	}
}

// NewInlineQueryResultCachedSticker create a new inline query with cached sticker.
func NewInlineQueryResultCachedSticker(id, stickerID, title string) InlineQueryResultCachedSticker {
	return InlineQueryResultCachedSticker{
		Type:      "sticker",
		ID:        id,
		StickerID: stickerID,
		Title:     title,
	}
}

// NewInlineQueryResultAudio creates a new inline query audio.
func NewInlineQueryResultAudio(id, url, title string) InlineQueryResultAudio {
	return InlineQueryResultAudio{
		Type:  "audio",
		ID:    id,
		URL:   url,
		Title: title,
	}
}

// NewInlineQueryResultCachedAudio create a new inline query with cached photo.
func NewInlineQueryResultCachedAudio(id, audioID string) InlineQueryResultCachedAudio {
	return InlineQueryResultCachedAudio{
		Type:    "audio",
		ID:      id,
		AudioID: audioID,
	}
}

// NewInlineQueryResultVoice creates a new inline query voice.
func NewInlineQueryResultVoice(id, url, title string) InlineQueryResultVoice {
	return InlineQueryResultVoice{
		Type:  "voice",
		ID:    id,
		URL:   url,
		Title: title,
	}
}

// NewInlineQueryResultCachedVoice create a new inline query with cached photo.
func NewInlineQueryResultCachedVoice(id, voiceID, title string) InlineQueryResultCachedVoice {
	return InlineQueryResultCachedVoice{
		Type:    "voice",
		ID:      id,
		VoiceID: voiceID,
		Title:   title,
	}
}

// NewInlineQueryResultDocument creates a new inline query document.
func NewInlineQueryResultDocument(id, url, title, mimeType string) InlineQueryResultDocument {
	return InlineQueryResultDocument{
		Type:     "document",
		ID:       id,
		URL:      url,
		Title:    title,
		MimeType: mimeType,
	}
}

// NewInlineQueryResultCachedDocument create a new inline query with cached photo.
func NewInlineQueryResultCachedDocument(id, documentID, title string) InlineQueryResultCachedDocument {
	return InlineQueryResultCachedDocument{
		Type:       "document",
		ID:         id,
		DocumentID: documentID,
		Title:      title,
	}
}

// NewInlineQueryResultLocation creates a new inline query location.
func NewInlineQueryResultLocation(id, title string, latitude, longitude float64) InlineQueryResultLocation {
	return InlineQueryResultLocation{
		Type:      "location",
		ID:        id,
		Title:     title,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewInlineQueryResultVenue creates a new inline query venue.
func NewInlineQueryResultVenue(id, title, address string, latitude, longitude float64) InlineQueryResultVenue {
	return InlineQueryResultVenue{
		Type:      "venue",
		ID:        id,
		Title:     title,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// NewEditMessageText allows you to edit the text of a message.
func NewEditMessageText(chatID any, messageID int, text string) EditMessageTextConfig {
	toID := getChatID(chatID)
	return EditMessageTextConfig{
		BaseEdit: BaseEdit{
			ChatID:    toID,
			MessageID: messageID,
		},
		Text: text,
	}
}

// NewEditMessageTextAndMarkup allows you to edit the text and reply markup of a message.
func NewEditMessageTextAndMarkup(chatID any, messageID int, text string, replyMarkup InlineKeyboardMarkup) EditMessageTextConfig {
	toID := getChatID(chatID)
	return EditMessageTextConfig{
		BaseEdit: BaseEdit{
			ChatID:      toID,
			MessageID:   messageID,
			ReplyMarkup: &replyMarkup,
		},
		Text: text,
	}
}

// NewEditMessageCaption allows you to edit the caption of a message.
func NewEditMessageCaption(chatID any, messageID int, caption string) EditMessageCaptionConfig {
	toID := getChatID(chatID)
	return EditMessageCaptionConfig{
		BaseEdit: BaseEdit{
			ChatID:    toID,
			MessageID: messageID,
		},
		Caption: caption,
	}
}

// NewEditMessageReplyMarkup allows you to edit the inline
// keyboard markup.
func NewEditMessageReplyMarkup(chatID any, messageID int, replyMarkup InlineKeyboardMarkup) EditMessageReplyMarkupConfig {
	toID := getChatID(chatID)
	return EditMessageReplyMarkupConfig{
		BaseEdit: BaseEdit{
			ChatID:      toID,
			MessageID:   messageID,
			ReplyMarkup: &replyMarkup,
		},
	}
}

// NewRemoveKeyboard hides the keyboard, with the option for being selective
// or hiding for everyone.
func NewRemoveKeyboard(selective bool) ReplyKeyboardRemove {
	return ReplyKeyboardRemove{
		RemoveKeyboard: true,
		Selective:      selective,
	}
}

// NewKeyboardButton creates a regular keyboard button.
func NewKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

// NewButton creates a RequestUser keyboard button.
// request_id: The user id to be passed to the bot as message.Update. request_id must be me unic. Else it will be cahshed.
// only_prmium: Pass True to request only premium users, pass False to request only non-premium users or leave empty to request all users.
// only_bots: Pass True to request only bot users, pass False to request only regular users or leave empty to request all users.
func NewButtonRequestUser(request_id int, only_premium bool, only_bots bool) KeyboardButtonRequestUser {
	return KeyboardButtonRequestUser{
		RequestID:     request_id,
		UserIsPremium: only_premium,
		UserIsBot:     only_bots,
	}
}

// NewKeyboardButtonRequestUser creates a keyboard button with text
// and RequestUser for a callback.
func NewKeyboardButtonRequestUser(text string, button KeyboardButtonRequestUser) KeyboardButton {
	return KeyboardButton{
		Text:        text,
		RequestUser: &button,
	}
}

// NewButton creates a RequestChat keyboard button.
// request_id: The user id to be passed to the bot as message.Update. request_id must be me unic. Else it will be cahshed.

func NewButtonRequestChat(request_id int, is_channel bool, is_forum bool, bot_is_member bool) KeyboardButtonRequestChat {
	return KeyboardButtonRequestChat{
		RequestID: request_id,
		// request_id: The user id to be passed to the bot as message.Update. request_id must be me unic. Else it will be cahshed.
		ChatIsChannel: is_channel,
		ChatIsForum:   is_forum,
		BotIsMember:   bot_is_member,
	}
}

// NewKeyboardButtonRequestUser creates a keyboard button with text
// and RequestUser for a callback.
func NewKeyboardButtonRequestChat(text string, button KeyboardButtonRequestChat) KeyboardButton {
	return KeyboardButton{
		Text:        text,
		RequestChat: &button,
	}
}

// NewKeyboardButtonWebApp creates a keyboard button with text
// which goes to a WebApp.
func NewKeyboardButtonWebApp(text string, webapp WebAppInfo) KeyboardButton {
	return KeyboardButton{
		Text:   text,
		WebApp: &webapp,
	}
}

// NewKeyboardButtonContact creates a keyboard button that requests
// user contact information upon click.
func NewKeyboardButtonContact(text string) KeyboardButton {
	return KeyboardButton{
		Text:           text,
		RequestContact: true,
	}
}

// NewKeyboardButtonLocation creates a keyboard button that requests
// user location information upon click.
func NewKeyboardButtonLocation(text string) KeyboardButton {
	return KeyboardButton{
		Text:            text,
		RequestLocation: true,
	}
}

// NewKeyboardButtonRow creates a row of keyboard buttons.
func NewKeyboardButtonRow(buttons ...KeyboardButton) []KeyboardButton {
	var row []KeyboardButton

	row = append(row, buttons...)

	return row
}

// NewReplyKeyboard creates a new regular keyboard with sane defaults.
func NewReplyKeyboard(rows ...[]KeyboardButton) ReplyKeyboardMarkup {
	var keyboard [][]KeyboardButton

	keyboard = append(keyboard, rows...)

	return ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Keyboard:       keyboard,
	}
}

// NewOneTimeReplyKeyboard creates a new one time keyboard.
func NewOneTimeReplyKeyboard(rows ...[]KeyboardButton) ReplyKeyboardMarkup {
	markup := NewReplyKeyboard(rows...)
	markup.OneTimeKeyboard = true
	return markup
}

// NewInlineKeyboardButtonData creates an inline keyboard button with text
// and data for a callback.
func NewInlineKeyboardButtonData(text, data string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:         text,
		CallbackData: &data,
	}
}

// NewInlineKeyboardButtonWebApp creates an inline keyboard button with text
// which goes to a WebApp.
func NewInlineKeyboardButtonWebApp(text string, webapp WebAppInfo) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:   text,
		WebApp: &webapp,
	}
}

// NewInlineKeyboardButtonLoginURL creates an inline keyboard button with text
// which goes to a LoginURL.
func NewInlineKeyboardButtonLoginURL(text string, loginURL LoginURL) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:     text,
		LoginURL: &loginURL,
	}
}

// NewInlineKeyboardButtonURL creates an inline keyboard button with text
// which goes to a URL.
func NewInlineKeyboardButtonURL(text, url string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
		URL:  &url,
	}
}

// NewInlineKeyboardButtonSwitch creates an inline keyboard button with
// text which allows the user to switch to a chat or return to a chat.
func NewInlineKeyboardButtonSwitch(text, sw string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text:              text,
		SwitchInlineQuery: &sw,
	}
}

// NewInlineKeyboardRow creates an inline keyboard row with buttons.
func NewInlineKeyboardRow(buttons ...InlineKeyboardButton) []InlineKeyboardButton {
	var row []InlineKeyboardButton

	row = append(row, buttons...)

	return row
}

// NewInlineKeyboardMarkup creates a new inline keyboard.
func NewInlineKeyboardMarkup(rows ...[]InlineKeyboardButton) InlineKeyboardMarkup {
	var keyboard [][]InlineKeyboardButton

	keyboard = append(keyboard, rows...)

	return InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

// NewCallback creates a new callback message.
func NewCallback(id, text string) CallbackConfig {
	return CallbackConfig{
		CallbackQueryID: id,
		Text:            text,
		ShowAlert:       false,
	}
}

// NewCallbackWithAlert creates a new callback message that alerts
// the user.
func NewCallbackWithAlert(id, text string) CallbackConfig {
	return CallbackConfig{
		CallbackQueryID: id,
		Text:            text,
		ShowAlert:       true,
	}
}

// NewInvoice creates a new Invoice request to the user.
func NewInvoice(chatID any, title, description, payload, providerToken, startParameter string, currency Currency, prices []LabeledPrice) InvoiceConfig {
	toID := getChatID(chatID)
	return InvoiceConfig{
		BaseChat:       BaseChat{ChatID: toID},
		Title:          title,
		Description:    description,
		Payload:        payload,
		ProviderToken:  providerToken,
		StartParameter: startParameter,
		Currency:       currency,
		Prices:         prices}
}

// NewChatTitle allows you to update the title of a chat.
func NewChatTitle(chatID any, title string) SetChatTitleConfig {
	toID := getChatID(chatID)
	return SetChatTitleConfig{
		ChatID: toID,
		Title:  title,
	}
}

// NewChatDescription allows you to update the description of a chat.
func NewChatDescription(chatID any, description string) SetChatDescriptionConfig {
	toID := getChatID(chatID)
	return SetChatDescriptionConfig{
		ChatID:      toID,
		Description: description,
	}
}

// NewChatPhoto allows you to update the photo for a chat.
func NewChatPhoto(chatID any, photo RequestFileData) SetChatPhotoConfig {
	toID := getChatID(chatID)
	return SetChatPhotoConfig{
		BaseFile: BaseFile{
			BaseChat: BaseChat{
				ChatID: toID,
			},
			File: photo,
		},
	}
}

// NewDeleteChatPhoto allows you to delete the photo for a chat.
func NewDeleteChatPhoto(chatID any) DeleteChatPhotoConfig {
	toID := getChatID(chatID)
	return DeleteChatPhotoConfig{
		ChatID: toID,
	}
}

// NewPoll allows you to create a new poll.
func NewPoll(chatID any, question string, options ...string) SendPollConfig {
	toID := getChatID(chatID)
	return SendPollConfig{
		BaseChat: BaseChat{
			ChatID: toID,
		},
		Question:    question,
		Options:     options,
		IsAnonymous: true, // This is Telegram's default.
	}
}

// NewStopPoll allows you to stop a poll.
func NewStopPoll(chatID any, messageID int) StopPollConfig {
	toID := getChatID(chatID)
	return StopPollConfig{
		BaseEdit{
			ChatID:    toID,
			MessageID: messageID,
		},
	}
}

// NewDice allows you to send a random dice roll.
func NewDice(chatID any) DiceConfig {
	toID := getChatID(chatID)
	return DiceConfig{
		BaseChat: BaseChat{
			ChatID: toID,
		},
	}
}

// NewDiceWithEmoji allows you to send a random roll of one of many types.
//
// Emoji may be 🎲 (1-6), 🎯 (1-6), or 🏀 (1-5).
func NewDiceWithEmoji(chatID any, emoji string) DiceConfig {
	toID := getChatID(chatID)
	return DiceConfig{
		BaseChat: BaseChat{
			ChatID: toID,
		},
		Emoji: emoji,
	}
}

// NewBotCommandScopeDefault represents the default scope of bot commands.
func NewBotCommandScopeDefault() BotCommandScope {
	return BotCommandScope{Type: "default"}
}

// NewBotCommandScopeAllPrivateChats represents the scope of bot commands,
// covering all private chats.
func NewBotCommandScopeAllPrivateChats() BotCommandScope {
	return BotCommandScope{Type: "all_private_chats"}
}

// NewBotCommandScopeAllGroupChats represents the scope of bot commands,
// covering all group and supergroup chats.
func NewBotCommandScopeAllGroupChats() BotCommandScope {
	return BotCommandScope{Type: "all_group_chats"}
}

// NewBotCommandScopeAllChatAdministrators represents the scope of bot commands,
// covering all group and supergroup chat administrators.
func NewBotCommandScopeAllChatAdministrators() BotCommandScope {
	return BotCommandScope{Type: "all_chat_administrators"}
}

// NewBotCommandScopeChat represents the scope of bot commands, covering a
// specific chat.
func NewBotCommandScopeChat(chatID any) BotCommandScope {
	toID := getChatID(chatID)
	return BotCommandScope{
		Type:   "chat",
		ChatID: toID,
	}
}

// NewBotCommandScopeChatAdministrators represents the scope of bot commands,
// covering all administrators of a specific group or supergroup chat.
func NewBotCommandScopeChatAdministrators(chatID any) BotCommandScope {
	toID := getChatID(chatID)
	return BotCommandScope{
		Type:   "chat_administrators",
		ChatID: toID,
	}
}

// NewBotCommandScopeChatMember represents the scope of bot commands, covering a
// specific member of a group or supergroup chat.
func NewBotCommandScopeChatMember(chatID, userID int64) BotCommandScope {
	return BotCommandScope{
		Type:   "chat_member",
		ChatID: chatID,
		UserID: userID,
	}
}

// NewGetMyCommandsWithScope allows you to set the registered commands for a
// given scope.
func NewGetMyCommandsWithScope(scope BotCommandScope) GetMyCommandsConfig {
	return GetMyCommandsConfig{Scope: &scope}
}

// NewGetMyCommandsWithScopeAndLanguage allows you to set the registered
// commands for a given scope and language code.
func NewGetMyCommandsWithScopeAndLanguage(scope BotCommandScope, languageCode string) GetMyCommandsConfig {
	return GetMyCommandsConfig{Scope: &scope, LanguageCode: languageCode}
}

// NewSetMyCommands allows you to set the registered commands.
func NewSetMyCommands(commands ...BotCommand) SetMyCommandsConfig {
	return SetMyCommandsConfig{Commands: commands}
}

// NewSetMyCommandsWithScope allows you to set the registered commands for a given scope.
func NewSetMyCommandsWithScope(scope BotCommandScope, commands ...BotCommand) SetMyCommandsConfig {
	return SetMyCommandsConfig{Commands: commands, Scope: &scope}
}

// NewSetMyCommandsWithScopeAndLanguage allows you to set the registered commands for a given scope
// and language code.
func NewSetMyCommandsWithScopeAndLanguage(scope BotCommandScope, languageCode string, commands ...BotCommand) SetMyCommandsConfig {
	return SetMyCommandsConfig{Commands: commands, Scope: &scope, LanguageCode: languageCode}
}

// NewDeleteMyCommands allows you to delete the registered commands.
func NewDeleteMyCommands() DeleteMyCommandsConfig {
	return DeleteMyCommandsConfig{}
}

// NewDeleteMyCommandsWithScope allows you to delete the registered commands for a given
// scope.
func NewDeleteMyCommandsWithScope(scope BotCommandScope) DeleteMyCommandsConfig {
	return DeleteMyCommandsConfig{Scope: &scope}
}

// NewDeleteMyCommandsWithScopeAndLanguage allows you to delete the registered commands for a given
// scope and language code.
func NewDeleteMyCommandsWithScopeAndLanguage(scope BotCommandScope, languageCode string) DeleteMyCommandsConfig {
	return DeleteMyCommandsConfig{Scope: &scope, LanguageCode: languageCode}
}

// ValidateWebAppData validate data received via the Web App
// https://core.telegram.org/bots/webapps#validating-data-received-via-the-web-app
func ValidateWebAppData(token, telegramInitData string) (bool, error) {
	initData, err := url.ParseQuery(telegramInitData)
	if err != nil {
		return false, fmt.Errorf("error parsing data %w", err)
	}

	dataCheckString := make([]string, 0, len(initData))
	for k, v := range initData {
		if k == "hash" {
			continue
		}
		if len(v) > 0 {
			dataCheckString = append(dataCheckString, fmt.Sprintf("%s=%s", k, v[0]))
		}
	}

	sort.Strings(dataCheckString)

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(token))

	hHash := hmac.New(sha256.New, secret.Sum(nil))
	hHash.Write([]byte(strings.Join(dataCheckString, "\n")))

	hash := hex.EncodeToString(hHash.Sum(nil))

	if initData.Get("hash") != hash {
		return false, errors.New("hash not equal")
	}

	return true, nil
}

// NewCreateInvoiceLinkConfig creates a new CreateInvoiceLinkConfig with the given parameters.
//
// The ProviderToken must be obtained from the Telegram BotFather.
// The Currency must be a valid ISO 4217 3-letter currency code. Use Currency constant list. XTR for payments in Telegram Stars.
// The Prices must contain at least one LabeledPrice.
//
// The returned CreateInvoiceLinkConfig can be passed to the BotAPI.CreateInvoiceLink
// method to create an invoice link.
func NewCreateInvoiceLinkConfig(chatID any, title, description, payload, providerToken string, prices []LabeledPrice, currency Currency) CreateInvoiceLinkConfig {
	toID := getChatID(chatID)
	return CreateInvoiceLinkConfig{
		BaseChat:      BaseChat{ChatID: toID},
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

// NewGetCustomEmojiStickersConfig creates a new GetCustomEmojiStickersConfig with the specified custom emoji IDs.
// It accepts a variable number of string arguments, each representing a custom emoji ID.
func NewGetCustomEmojiStickersConfig(val ...string) GetCustomEmojiStickersConfig {
	return GetCustomEmojiStickersConfig{CustomEmojiIDs: val}
}

// ChatAction — enum actions for sendChatAction.
type ChatAction string

const (
	// typing text
	ChatActionTyping ChatAction = "typing"
	// uploading photo
	ChatActionUploadPhoto ChatAction = "upload_photo"
	// recording video
	ChatActionRecordVideo ChatAction = "record_video"
	// uploading video
	ChatActionUploadVideo ChatAction = "upload_video"
	// recording audio
	ChatActionRecordVoice ChatAction = "record_voice"
	// uploading audio
	ChatActionUploadVoice ChatAction = "upload_voice"
	// uploading document
	ChatActionUploadDocument ChatAction = "upload_document"
	// choosing sticker
	ChatActionChooseSticker ChatAction = "choose_sticker"
	// finding location
	ChatActionFindLocation ChatAction = "find_location"
	// recording video note
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	// uploading video note
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

// ToString возвращает строковое представление действия
func (a ChatAction) String() string {
	return string(a)
}

// NewChatActionConfig creates a new ChatActionConfig for sending a chat action.
// Now chatID can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User, Message
func NewChatActionConfig(to any, act ChatAction) ChatActionConfig {
	base := BaseChat{}
	var action ChatAction
	switch val := to.(type) {
	case int64:
		action = act
		base = BaseChat{ChatID: val}
	case BaseChat:
		action = act
		base = val
	case ChatConfig:
		action = act
		base = BaseChat{ChatID: val.ChatID}
	case ChatActionConfig:
		return val
	case Chat:
		action = act
		base = BaseChat{
			ChatID:          val.ID,
			ProtectContent:  val.HasProtectedContent,
			ChannelUsername: val.UserName,
		}
	case User:
		action = act
		base = BaseChat{ChatID: val.ID}
	case Message:
		action = act
		base = BaseChat{
			ChatID:           val.Chat.ID,
			ProtectContent:   val.HasProtectedContent,
			ChannelUsername:  val.Chat.UserName,
			ReplyToMessageID: val.ReplyToMessage.MessageID,
			ReplyMarkup:      val.ReplyMarkup,
			MessageThreadID:  val.Message_thread_id,
		}
	default:
		action = act
		base = BaseChat{ChatID: 0}
	}
	return ChatActionConfig{BaseChat: base,
		Action: action,
	}
}

// NewCreateForumTopicConfig creates a new CreateForumTopicConfig with the specified parameters.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// name is the name of the forum topic, which must be between 1 and 128 characters.
// iconColor is the RGB color of the topic icon and must be one of the allowed values.
// Optional iconCustomEmojiID is the ID of the custom emoji to use as the topic icon.
func NewCreateForumTopicConfig(chatID any, name string, iconColor int) CreateForumTopicConfig {
	toID := getChatID(chatID)
	return CreateForumTopicConfig{
		ChatID:    toID,
		Name:      "Test",
		IconColor: 2,
	}
}

// NewCreateForumTopicConfig creates a new CreateForumTopicConfig with the specified parameters.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// name is the name of the forum topic, which must be between 1 and 128 characters.
// iconColor is the RGB color of the topic icon and must be one of the allowed values.
// iconCustomEmojiID is the ID of the custom emoji to use as the topic icon.
func NewEditForumTopicConfig(chatID any, messageThreadID int, name string, iconCustomEmojiID string) EditForumTopicConfig {
	toID := getChatID(chatID)
	return EditForumTopicConfig{
		ChatID:            toID,
		MessageThreadID:   messageThreadID,
		Name:              name,
		IconCustomEmojiID: iconCustomEmojiID,
	}
}

// NewEditForumTopicConfigWithotIcon creates a new EditForumTopicConfig without specifying an icon.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// messageThreadID is the unique identifier of the forum topic thread.
// name is the name of the forum topic, which must be non-empty.
func NewEditForumTopicConfigWithotIcon(chatID any, messageThreadID int, name string) EditForumTopicConfig {
	toID := getChatID(chatID)
	return EditForumTopicConfig{
		ChatID:          toID,
		MessageThreadID: messageThreadID,
		Name:            name,
	}
}

// NewCloseForumTopicConfig creates a configuration to close a forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// messageThreadID is the unique identifier of the forum topic thread to be closed.
func NewCloseForumTopicConfig(chatID any, messageThreadID int) CloseForumTopicConfig {
	toID := getChatID(chatID)
	return CloseForumTopicConfig{
		ChatID:          toID,
		MessageThreadID: messageThreadID,
	}
}

// NewDeleteForumTopicConfig creates a configuration to delete a forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// messageThreadID is the unique identifier of the forum topic thread to be deleted.
func NewDeleteForumTopicConfig(chatID any, messageThreadID int) DeleteForumTopicConfig {
	toID := getChatID(chatID)
	return DeleteForumTopicConfig{
		ChatID:          toID,
		MessageThreadID: messageThreadID,
	}
}

// NewReopenForumTopicConfig creates a configuration to reopen a forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// messageThreadID is the unique identifier of the forum topic thread to be reopened.
func NewReopenForumTopicConfig(chatID any, messageThreadID int) ReopenForumTopicConfig {
	toID := getChatID(chatID)
	return ReopenForumTopicConfig{
		ChatID:          toID,
		MessageThreadID: messageThreadID,
	}
}

// NewUnpinAllForumTopicMessagesConfig creates a configuration to unpin all pinned messages in a forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// messageThreadID is the unique identifier of the forum topic thread to be unpinned.
func NewUnpinAllForumTopicMessagesConfig(chatID any, messageThreadID int) UnpinAllForumTopicMessagesConfig {
	toID := getChatID(chatID)
	return UnpinAllForumTopicMessagesConfig{
		ChatID:          toID,
		MessageThreadID: messageThreadID,
	}
}

// getChatID returns the chat ID of the given argument.
//
// The argument can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User.
// If the argument is not one of the above types, it returns 0.
func getChatID(chatID any) int64 {
	var toID int64
	switch val := chatID.(type) {
	case int64:
		toID = val
	case BaseChat:
		toID = val.ChatID
	case ChatConfig:
		toID = val.ChatID
	case ChatActionConfig:
		toID = val.ChatID
	case Chat:
		toID = val.ID
	case User:
		toID = val.ID
	case Message:
		toID = val.Chat.ID
	default:
		toID = 0
	}
	return toID
}

// NewEditGeneralForumTopicConfig creates a configuration to edit a general forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
// name is the new name of the general forum topic.
// iconCustomEmojiID is the ID of the custom emoji to use as the new icon of the general forum topic.
func NewEditGeneralForumTopicConfig(chatID any, name string, iconCustomEmojiID string) EditGeneralForumTopicConfig {
	toID := getChatID(chatID)
	return EditGeneralForumTopicConfig{
		ChatID:            toID,
		Name:              name,
		IconCustomEmojiID: iconCustomEmojiID,
	}
}

// NewCloseGeneralForumTopicConfig creates a configuration to close a general forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
func NewCloseGeneralForumTopicConfig(chatID any) CloseGeneralForumTopicConfig {
	toID := getChatID(chatID)
	return CloseGeneralForumTopicConfig{
		ChatID: toID,
	}
}

// NewReopenGeneralForumTopicConfig creates a configuration to reopen a general forum topic.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
func NewReopenGeneralForumTopicConfig(chatID any) ReopenGeneralForumTopicConfig {
	toID := getChatID(chatID)
	return ReopenGeneralForumTopicConfig{
		ChatID: toID,
	}
}

// NewHideGeneralForumTopicConfig creates a configuration to hide a general forum topic in a chat.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
func NewHideGeneralForumTopicConfig(chatID any) HideGeneralForumTopicConfig {
	toID := getChatID(chatID)
	return HideGeneralForumTopicConfig{
		ChatID: toID,
	}
}

// NewUnhideGeneralForumTopicConfig creates a configuration to unhide a general forum topic in a chat.
// chatID can be of any type that can be converted to a valid ChatID for the forum topic.
func NewUnhideGeneralForumTopicConfig(chatID any) UnhideGeneralForumTopicConfig {
	toID := getChatID(chatID)
	return UnhideGeneralForumTopicConfig{
		ChatID: toID,
	}
}
