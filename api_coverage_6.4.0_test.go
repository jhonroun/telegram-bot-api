package tgbotapi

import "testing"

func Test_64_EditGeneralForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewEditGeneralForumTopicConfig(GroupWithTopicsChatID, "Test 2", "5420216386448270341")

	_, err := bot.EditGeneralForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_64_CloseGeneralForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewCloseGeneralForumTopicConfig(GroupWithTopicsChatID)

	_, err := bot.CloseGeneralForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_64_ReopenGeneralForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewReopenGeneralForumTopicConfig(GroupWithTopicsChatID)

	_, err := bot.ReopenGeneralForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_64_HideGeneralForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewHideGeneralForumTopicConfig(GroupWithTopicsChatID)

	_, err := bot.HideGeneralForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_64_UnhideGeneralForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewUnhideGeneralForumTopicConfig(GroupWithTopicsChatID)

	_, err := bot.UnhideGeneralForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}
