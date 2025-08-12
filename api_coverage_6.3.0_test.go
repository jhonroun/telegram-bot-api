package tgbotapi

import "testing"

func Test_63_createForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := CreateForumTopicConfig{
		ChatID:            GroupWithTopicsChatID,
		Name:              "Test",
		IconColor:         2,
		IconCustomEmojiID: "test",
	}

	_, err := bot.CreateForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_63_editForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := EditForumTopicConfig{
		ChatID:            GroupWithTopicsChatID,
		Name:              "Changed name of topic № 2 Test new",
		MessageThreadID:   2,
		IconCustomEmojiID: "5420216386448270341",
	}

	_, err := bot.EditForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_get_closeForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := CloseForumTopicConfig{
		ChatID:          GroupWithTopicsChatID,
		MessageThreadID: 2,
	}

	_, err := bot.CloseForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_get_reopenForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := ReopenForumTopicConfig{
		ChatID:          GroupWithTopicsChatID,
		MessageThreadID: 2,
	}

	_, err := bot.ReopenForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_get_deleteForumTopic(t *testing.T) {
	bot := getBot(t)

	cfg := DeleteForumTopicConfig{
		ChatID:          GroupWithTopicsChatID,
		MessageThreadID: 2,
	}

	_, err := bot.DeleteForumTopic(cfg)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_63_get_unpinAllForumTopicMessages(t *testing.T) {
	bot := getBot(t)

	cfg := UnpinAllForumTopicMessagesConfig{
		ChatID:          GroupWithTopicsChatID,
		MessageThreadID: 3,
	}

	_, err := bot.UnpinAllForumTopicMessages(cfg)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_63_getForumTopicIconStickers(t *testing.T) {
	bot := getBot(t)

	forumTopic, err := bot.GetForumTopicIconStickers(GetForumTopicIconStickersConfig{})
	if err != nil {
		t.Fatal(err)
	}

	PrintStruct(t, forumTopic)
}
