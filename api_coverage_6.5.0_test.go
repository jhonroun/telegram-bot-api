package tgbotapi

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func Test_65_SendKeyboard_RequestUser_FF_Integration(t *testing.T) {
	t.Skip("works fine")
	bot := getBot(t)

	last := 0
	for {
		ups, err := bot.GetUpdates(UpdateConfig{Offset: last + 1, Limit: 100, Timeout: 0})
		if err != nil {
			t.Fatal(err)
		}
		if len(ups) == 0 {
			break
		}
		last = ups[len(ups)-1].UpdateID
	}

	timeout := 15
	u := NewUpdate(last + 1)                                    // last+1
	u.Timeout = timeout                                         // long polling
	u.AllowedUpdates = []string{"message", "chat_join_request"} //  message!

	updates := bot.GetUpdatesChan(u)

	reqUser := NewButtonRequestUser(111, false, false) // request_id must be me unic. Else it will be cahshed.
	button := NewKeyboardButtonRequestUser("Share a user", reqUser)

	markup := ReplyKeyboardMarkup{
		Keyboard:        [][]KeyboardButton{{button}},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	msg := NewMessage(ChatID, "Test user request (not prem, not bot)")
	msg.ReplyMarkup = markup // answer passed as simple update.message

	if _, err := bot.Send(msg); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			t.Skipf("no user_shared/chat_shared received within %ds; tap the button and pick a user in the SAME chat", timeout)
			return
		case upd := <-updates:
			if upd.Message != nil {

				if upd.Message.UserShared != nil {
					t.Logf("received user_shared: request_id=%d user_id=%d",
						upd.Message.UserShared.RequestID, upd.Message.UserShared.UserID)
					PrintStruct(t, upd.Message.UserShared)
					return
				}
				if upd.Message.ChatShared != nil {
					t.Logf("received chat_shared: request_id=%d chat_id=%d",
						upd.Message.ChatShared.RequestID, upd.Message.ChatShared.ChatID)
					return
				}
			}
			if upd.ChatJoinRequest != nil {
				if upd.ChatJoinRequest.UserChatID == 0 {
					b, _ := json.Marshal(upd.ChatJoinRequest)
					t.Fatalf("expected user_chat_id > 0; got: %s", string(b))
				}
				t.Logf("received chat_join_request with user_chat_id=%d", upd.ChatJoinRequest.UserChatID)
				return
			}
		}
	}

}

func Test_65_SendKeyboard_RequestUser_TF_Integration(t *testing.T) {
	t.Skip("works fine")
	bot := getBot(t)

	req_user := NewButtonRequestUser(127, true, false)
	button := NewKeyboardButtonRequestUser("Share a user", req_user)

	markup := ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{
			{
				button,
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	msg := NewMessage(ChatID, "Test user request (only prem, not bot)")
	msg.ReplyMarkup = markup

	_, err := bot.Send(msg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_65_SendKeyboard_RequestUser_FT_Integration(t *testing.T) {
	t.Skip("works fine")
	bot := getBot(t)

	req_user := NewButtonRequestUser(103, false, true)
	button := NewKeyboardButtonRequestUser("Share a user", req_user)

	markup := ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{
			{
				button,
			},
		},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	msg := NewMessage(ChatID, "Test user request (not prem, only bot)")
	msg.ReplyMarkup = markup

	_, err := bot.Send(msg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_65_SendKeyboard_RequestChat_Integration(t *testing.T) {
	t.Skip("test passed")
	bot := getBot(t)

	last := 0
	for {
		ups, err := bot.GetUpdates(UpdateConfig{Offset: last + 1, Limit: 100, Timeout: 0})
		if err != nil {
			t.Fatal(err)
		}
		if len(ups) == 0 {
			break
		}
		last = ups[len(ups)-1].UpdateID
	}

	timeout := 15
	u := NewUpdate(last + 1)                                    // last+1
	u.Timeout = timeout                                         // long polling
	u.AllowedUpdates = []string{"message", "chat_join_request"} //  message!

	updates := bot.GetUpdatesChan(u)

	reqChat := NewButtonRequestChat(129, false, true, true) // request_id must be me unic. Else it will be cahshed.
	button := NewKeyboardButtonRequestChat("Share a chat", reqChat)

	markup := ReplyKeyboardMarkup{
		Keyboard:        [][]KeyboardButton{{button}},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}

	msg := NewMessage(ChatID, "Test chat request (is_channel false, is_forum true, bot_is_member true)")
	msg.ReplyMarkup = markup // answer passed as simple update.message

	if _, err := bot.Send(msg); err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			t.Skipf("no user_shared/chat_shared received within %ds; tap the button and pick a user in the SAME chat", timeout)
			return
		case upd := <-updates:
			if upd.Message != nil {

				if upd.Message.UserShared != nil {
					t.Logf("received user_shared: request_id=%d user_id=%d",
						upd.Message.UserShared.RequestID, upd.Message.UserShared.UserID)
					PrintStruct(t, upd.Message.UserShared)
					return
				}
				if upd.Message.ChatShared != nil {
					t.Logf("received chat_shared: request_id=%d chat_id=%d",
						upd.Message.ChatShared.RequestID, upd.Message.ChatShared.ChatID)
					return
				}
			}
			if upd.ChatJoinRequest != nil {
				if upd.ChatJoinRequest.UserChatID == 0 {
					b, _ := json.Marshal(upd.ChatJoinRequest)
					t.Fatalf("expected user_chat_id > 0; got: %s", string(b))
				}
				t.Logf("received chat_join_request with user_chat_id=%d", upd.ChatJoinRequest.UserChatID)
				return
			}
		}
	}

}
