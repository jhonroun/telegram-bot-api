package tgbotapi

import "testing"

func Test_63_createForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewCreateForumTopicConfig(GroupWithTopicsChatID, "Test", 2)

	_, err := bot.CreateForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_63_editForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewEditForumTopicConfig(GroupWithTopicsChatID, 2, "Changed name of topic № 2 Test new", "5420216386448270341")

	_, err := bot.EditForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_get_closeForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewCloseForumTopicConfig(GroupWithTopicsChatID, 2)

	_, err := bot.CloseForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_get_reopenForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewReopenForumTopicConfig(GroupWithTopicsChatID, 2)

	_, err := bot.ReopenForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_get_deleteForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := NewDeleteForumTopicConfig(GroupWithTopicsChatID, 2)

	_, err := bot.DeleteForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_63_get_unpinAllForumTopicMessages(t *testing.T) {
	bot := getBot(t)

	cfg := NewUnpinAllForumTopicMessagesConfig(GroupWithTopicsChatID, 3)

	_, err := bot.UnpinAllForumTopicMessages(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_getForumTopicIconStickers(t *testing.T) {
	bot := getBot(t)

	forumTopic, err := bot.GetForumTopicIconStickers()
	if err != nil {
		t.Fatal(err)
	}

	PrintStruct(t, forumTopic)
}
