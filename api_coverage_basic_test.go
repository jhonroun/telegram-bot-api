package tgbotapi

import "testing"

func Test_getMe(t *testing.T) {
	bot := getBot(t)

	u, err := bot.GetMe()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(u)
}

func Test_username2UserID(t *testing.T) {
	bot := getBot(t)

	ch, err := bot.GetUserIDbyUsername("@tggobotapitest")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch)
}

func Test_getChatInfo(t *testing.T) {
	bot := getBot(t)

	ch, err := bot.GetChat(ChatInfoConfig{
		ChatConfig: ChatConfig{
			ChatID: GroupWithTopicsChatID,
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Log(ch)
}

func Test_chatActions(t *testing.T) {
	bot := getBot(t)

	_, err := bot.Typing(ChatID)
	if err != nil {
		t.Fatal(err)
	}
}
