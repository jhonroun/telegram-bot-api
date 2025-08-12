package tgbotapi

import (
	"strings"
	"testing"
	"time"
)

// buildShowcaseAST build full showcase AST. Without set the mode.
// This showcase as AST rendered in all modes. AST causes double escaping and conflict nesting.
func buildShowcaseAST() Node {
	return Group(
		// Заголовок и пустая строка
		Text("__PLACEHOLDER_TITLE__"), Text("\n\n"),

		// Инлайновые стили
		Bold(Text("bold text")), Text("\n"),
		Italic(Text("italic text")), Text("\n"),
		Underline(Text("underline")), Text("\n"),
		Strike(Text("strikethrough")), Text("\n"),
		Spoiler(Text("spoiler")), Text("\n"),

		// Вложенная строка: *bold _italic bold ~... ||...|| ~ __...__ _ bold*
		Group(
			Bold(
				Italic(
					Text("italic bold "),
					Strike(Text("italic bold strikethrough ")),
					Spoiler(Text("italic bold strikethrough spoiler")),
					Text(" "),
					Underline(Text("underline italic bold")),
				),
			),
			Text(" "),
			Bold(Text("bold")),
		),
		Text("\n"),

		// Ссылки / упоминания / emoji
		Link(Text("inline URL"), "http://www.example.com/"), Text("\n"),
		Mention(Text("inline mention of a user"), ChatID), Text("\n"),
		EmojiID("5368324170671202286", "👍"), Text("\n"),

		// Код: inline + block + block с языком
		Code("inline fixed-width code"), Text("\n"),
		Pre("pre-formatted fixed-width code block", LangPython), Text("\n"),
		Pre("pre-formatted fixed-width code block written in the Python programming language", "python"), Text("\n"),

		// Цитаты
		Quote(
			Text("Block quotation started"),
			Text("Block quotation continued"),
			Text("The last line of the block quotation"),
		),
		Text("\n"),
		QuoteExpandable(
			Text("Expandable block quotation started"),
			Text("Expandable block quotation continued"),
			Text("Expandable block quotation continued"),
			Text("Hidden by default part of the block quotation started"),
			Text("Expandable block quotation continued"),
			Text("The last line of the block quotation"),
		),
	)
}

// renderShowcase for test varios markdown modes
func renderShowcase(mode string, t *testing.T) string {
	ast := buildShowcaseAST()
	raw := "__PLACEHOLDER_TITLE__"
	title := mode + " showcase"
	text := Render(ast, mode)

	// учтём, что после Render плейсхолдер уже экранирован для текущего режима
	escaped := escapeText(mode, raw)

	return strings.Replace(text, escaped, title, 1)
}

func Test_56_SendShowcase_Markdown(t *testing.T) {
	bot := getBot(t)

	text := renderShowcase(ModeMarkdown, t)
	if l := len(text); l > 4096 {
		t.Fatalf("message too long for Telegram: %d bytes", l)
	}

	msg := NewMessage(ChatID, text)
	msg.ParseMode = ModeMarkdown

	if _, err := bot.Send(msg); err != nil {
		t.Fatalf("Markdown send failed: %v\nTEXT:\n%s", err, text)
	}
	time.Sleep(200 * time.Millisecond)
}

func Test_56_SendShowcase_MarkdownV2(t *testing.T) {

	bot := getBot(t)

	text := renderShowcase(ModeMarkdownV2, t)
	if l := len(text); l > 4096 {
		t.Fatalf("message too long for Telegram: %d bytes", l)
	}

	msg := NewMessage(ChatID, text)
	msg.ParseMode = ModeMarkdownV2

	if _, err := bot.Send(msg); err != nil {
		t.Fatalf("MarkdownV2 send failed: %v\nTEXT:\n%s", err, text)
	}
	time.Sleep(200 * time.Millisecond)
}

func Test_56_SendShowcase_HTML(t *testing.T) {
	bot := getBot(t)

	text := renderShowcase(ModeHTML, t)
	if l := len(text); l > 4096 {
		t.Fatalf("message too long for Telegram: %d bytes", l)
	}

	msg := NewMessage(ChatID, text)
	msg.ParseMode = ModeHTML

	if _, err := bot.Send(msg); err != nil {
		t.Fatalf("HTML send failed: %v\nTEXT:\n%s", err, text)
	}
	time.Sleep(200 * time.Millisecond)
}

// Доп. диагностика: построчная отправка MarkdownV2 (если когда-то «съестся» текст без ошибки).
func Test_56_Debug_MarkdownV2_LineByLine(t *testing.T) {
	t.Skip("enable for troubleshooting only")
	bot := getBot(t)

	text := renderShowcase(ModeMarkdownV2, t)
	lines := strings.Split(text, "\n")
	for i, ln := range lines {
		if ln == "" {
			continue
		}
		msg := NewMessage(ChatID, ln)
		msg.ParseMode = ModeMarkdownV2
		if _, err := bot.Send(msg); err != nil {
			t.Fatalf("line %d failed: %v\nLINE:\n%s", i+1, err, ln)
		}
		time.Sleep(120 * time.Millisecond)
	}
}

func Test_56_ResolveLanguage_Aliases(t *testing.T) {
	cases := map[string]Language{
		"html":  LangMarkup,
		"xml":   LangMarkup,
		"js":    LangJavaScript,
		"ts":    LangTypeScript,
		"py":    LangPython,
		"go":    LangGo,
		"yml":   LangYAML,
		"scss":  LangSCSS,
		"objc":  LangObjectiveC,
		"razor": LangRazorCS,
	}
	for in, want := range cases {
		if got := ResolveLanguage(in); got != want {
			t.Fatalf("ResolveLanguage(%q) = %q, want %q", in, got, want)
		}
	}
}

func Test_56_CodeBlockTest(t *testing.T) {
	bot := getBot(t)

	ast := Pre("package main\nimport (\n\"os\"\ntgbotapi \"github.com/jhonroun/telegram-bot-api\"\n)\n\nfunc main() {\nbot, err := tgbotapi.NewBotAPI(os.Getenv(\"TELEGRAM_APITOKEN\"))\nif err != nil {\npanic(err)\n}\n\nbot.Debug = true\n}", LangGo)

	text := Render(ast, ModeMarkdownV2)

	msg := NewMessage(ChatID, text)
	msg.ParseMode = ModeMarkdownV2

	if _, err := bot.Send(msg); err != nil {
		t.Fatalf("MarkdownV2 send failed: %v\nTEXT:\n%s", err, text)
	}
	time.Sleep(200 * time.Millisecond)
}

func Test_56_what_happend_when_markdown_doesnt_set(t *testing.T) {
	bot := getBot(t)

	text := renderShowcase(ModeMarkdownV2, t)

	msg := NewMessage(ChatID, text)
	if _, err := bot.Send(msg); err != nil {
		t.Fatalf("failed: %v\nTEXT:\n%s", err, text)
	}
	time.Sleep(200 * time.Millisecond)
}
