package tgbotapi

import (
	"fmt"
	"html"
	"strconv"
	"strings"
)

//
// -----------------------------
//            AST
// -----------------------------
// Additional tools for build Markuped Text

// Node — node of tree of markup (AST). Rendered to string according to target mode.
type Node interface {
	// render(mode string) rendering node to string according to target mode.
	// Inside render, only text leaves are escaped, formatting nodes add markers/tags around already rendered children.
	// formatting nodes add markers/tags around already rendered children
	render(mode string) string
}

// Text creates a text leaf AST. Escaping is performed according to the target mode at render time.
func Text(s string) Node { return textNode{s: s} }

type textNode struct{ s string }

func (n textNode) render(mode string) string { return escapeText(mode, n.s) }

// Code creates inline code (monospace). In Markdown/MarkdownV2, contents are NOT escaped. P.S. escaping only ` and \
func Code(s string) Node { return codeNode{s: s} }

type codeNode struct{ s string }

func (n codeNode) render(mode string) string {
	switch mode {
	case ModeHTML:
		return "<code>" + escapeText(mode, n.s) + "</code>"
	case ModeMarkdown, ModeMarkdownV2:
		s := strings.ReplaceAll(n.s, `\`, `\\`)
		s = strings.ReplaceAll(s, "`", "\\`")
		return "`" + s + "`"
	default:
		return n.s
	}
}

// Pre builds a multi-line code block.
// In Markdown/MarkdownV2, block contents are NOT escaped.
// In MarkdownV2 and HTML, an optional language can be provided to enable syntax highlighting
// in Telegram clients using libprisma (Prism) grammars.
// The language must be one of the supported Prism/libprisma languages.
// Use ResolveLanguage to validate or normalize from an alias (e.g., "html" → "markup").
func Pre(code string, lang Language) Node { return preNode{s: code, language: lang} }

// preNode builds a multi-line code block.
type preNode struct {
	s        string
	language Language
}

func (n preNode) render(mode string) string {
	switch mode {
	case ModeHTML:
		// HTML requires <pre><code class="language-xxx">...</code></pre>
		if n.language != "" {
			return `<pre><code class="language-` + html.EscapeString(string(n.language)) + `">` +
				escapeText(mode, n.s) + `</code></pre>`
		}
		return "<pre>" + escapeText(mode, n.s) + "</pre>"

	case ModeMarkdown:
		// Legacy Markdown has no language tag; do not escape the block content.
		return "```\n" + n.s + "\n```"

	case ModeMarkdownV2:
		// MarkdownV2 supports an optional language tag on the opening fence.
		if n.language != "" {
			return "```" + string(n.language) + "\n" + n.s + "\n```"
		}
		return "```\n" + n.s + "\n```"

	default:
		return n.s
	}
}

// Group concatenates multiple nodes without additional markers.
func Group(children ...Node) Node { return groupNode{children: children} }

type groupNode struct{ children []Node }

func (n groupNode) render(mode string) string {
	var b strings.Builder
	for _, c := range n.children {
		b.WriteString(c.render(mode))
	}
	return b.String()
}

// Bold make child node as bold: <b>...</b> (HTML) or *...* (MD/MDV2).
func Bold(children ...Node) Node {
	return wrapNode{left: openBold, right: closeBold, children: children}
}

// Italic makes child nodes italic: <i>...</i> (HTML), _..._ (MDV2).
func Italic(children ...Node) Node {
	return wrapNode{left: openItalic, right: closeItalic, children: children}
}

// Underline underlines: <u>...</u> (HTML), __...__ (MDV2). In Markdown (legacy), it is a soft degradation (no underlining).
func Underline(children ...Node) Node {
	return wrapNode{left: openUnderline, right: closeUnderline, children: children}
}

// Strike strikes: <s>...</s> (HTML), ~...~ (MDV2). In Markdown (legacy), it is a soft degradation (no strikethrough).
func Strike(children ...Node) Node {
	return wrapNode{left: openStrike, right: closeStrike, children: children}
}

// Spolier hides text under spoiler: <span class="tg-spoiler">...</span> (HTML), ||...|| (MDV2). In Markdown (legacy), it is a soft degradation (no underlining).
func Spoiler(children ...Node) Node {
	return wrapNode{left: openSpoiler, right: closeSpoiler, children: children}
}

type wrapNode struct {
	left, right func(string) string
	children    []Node
}

func (n wrapNode) render(mode string) string {
	var b strings.Builder
	for _, c := range n.children {
		b.WriteString(c.render(mode))
	}
	return n.left(mode) + b.String() + n.right(mode)
}

func openBold(m string) string {
	if m == ModeHTML {
		return "<b>"
	} else {
		return "*"
	}
}
func closeBold(m string) string {
	if m == ModeHTML {
		return "</b>"
	} else {
		return "*"
	}
}
func openItalic(m string) string {
	if m == ModeHTML {
		return "<i>"
	} else {
		return "_"
	}
}
func closeItalic(m string) string {
	if m == ModeHTML {
		return "</i>"
	} else {
		return "_"
	}
}

func openUnderline(m string) string {
	switch m {
	case ModeHTML:
		return "<u>"
	case ModeMarkdownV2:
		return "__"
	default:
		return ""
	}
}
func closeUnderline(m string) string {
	switch m {
	case ModeHTML:
		return "</u>"
	case ModeMarkdownV2:
		return "__"
	default:
		return ""
	}
}
func openStrike(m string) string {
	if m == ModeHTML {
		return "<s>"
	}
	if m == ModeMarkdownV2 {
		return "~"
	}
	return ""
}
func closeStrike(m string) string {
	if m == ModeHTML {
		return "</s>"
	}
	if m == ModeMarkdownV2 {
		return "~"
	}
	return ""
}
func openSpoiler(m string) string {
	if m == ModeHTML {
		return `<span class="tg-spoiler">`
	}
	if m == ModeMarkdownV2 {
		return "||"
	}
	return ""
}
func closeSpoiler(m string) string {
	if m == ModeHTML {
		return `</span>`
	}
	if m == ModeMarkdownV2 {
		return "||"
	}
	return ""
}

// Link creates a clickable link: <a href="...">label</a> (HTML), [label](url) (MD/MDV2).
// URL is escaped according to different rules for HTML/Markdown/MarkdownV2.
func Link(label Node, url string) Node { return linkNode{label: label, url: url} }

type linkNode struct {
	label Node
	url   string
}

func (n linkNode) render(mode string) string {
	switch mode {
	case ModeHTML:
		return `<a href="` + escapeURL(mode, n.url) + `">` + n.label.render(mode) + `</a>`
	case ModeMarkdown, ModeMarkdownV2:
		return "[" + n.label.render(mode) + "](" + escapeURL(mode, n.url) + ")"
	default:
		return n.label.render(mode) + " (" + n.url + ")"
	}
}

// Mention creates a mention by userID: tg://user?id=...
func Mention(label Node, userID int64) Node {
	return Link(label, "tg://user?id="+strconv.FormatInt(userID, 10))
}

// EmojiID creates a tg-emoji node (HTML-only). In other modes, returns fallback.
func EmojiID(id, fallback string) Node { return tgEmojiNode{id: id, fallback: fallback} }

type tgEmojiNode struct {
	id       string
	fallback string
}

func (n tgEmojiNode) render(mode string) string {
	if mode == ModeHTML {
		return `<tg-emoji emoji-id="` + html.EscapeString(n.id) + `">` +
			html.EscapeString(n.fallback) + `</tg-emoji>`
	}
	return escapeText(mode, n.fallback)
}

// Quote create a block quote: <blockquote>...</blockquote> (HTML), >... (MD/MDV2).
// In MD/MDV2, each line is prefixed with '>'.
func Quote(lines ...Node) Node { return quoteNode{lines: lines} }

// QuoteExpandable creates an expandable block quote: <blockquote expandable>...</blockquote> (HTML), >... In MD/MDV2 returns a regular quote.
func QuoteExpandable(lines ...Node) Node { return quoteNode{lines: lines, expandable: true} }

type quoteNode struct {
	lines      []Node
	expandable bool
}

func (n quoteNode) render(mode string) string {
	switch mode {
	case ModeHTML:
		var b strings.Builder
		for i, ln := range n.lines {
			if i > 0 {
				b.WriteByte('\n')
			}
			b.WriteString(ln.render(mode))
		}
		if n.expandable {
			return "<blockquote expandable>" + b.String() + "</blockquote>"
		}
		return "<blockquote>" + b.String() + "</blockquote>"
	case ModeMarkdown, ModeMarkdownV2:
		var b strings.Builder
		for i, ln := range n.lines {
			if i > 0 {
				b.WriteByte('\n')
			}
			b.WriteString(">")
			b.WriteString(ln.render(mode))
		}
		return b.String()
	default:
		var b strings.Builder
		for i, ln := range n.lines {
			if i > 0 {
				b.WriteByte('\n')
			}
			b.WriteString(ln.render(mode))
		}
		return b.String()
	}
}

// Render various AST to string according to target mode. Only text leaves are escaped, formatting nodes add markers/tags around already rendered children. This in only one published entry point. For return string, see RenderShowcase (markups_test.go).
func Render(n Node, mode string) string { return n.render(mode) }

//
// -----------------------------
//        Escaping helpers
// -----------------------------

// escapeText escaapes text according to target mode (only for pieces of Text).
func escapeText(mode, s string) string {
	switch mode {
	case ModeHTML:
		return html.EscapeString(s)
	case ModeMarkdown:
		return escapeMarkdown(s)
	case ModeMarkdownV2:
		return escapeMarkdownV2(s)
	default:
		return s
	}
}

// escapeMarkdown escapes a string for Markdown (legacy) mode. It escapes symbols which have special meaning in Markdown:
// _, *, `, and [.
func escapeMarkdown(text string) string {
	return strings.NewReplacer("_", "\\_", "*", "\\*", "`", "\\`", "[", "\\[").Replace(text)
}

// escapeMarkdownV2 escapes a string for MarkdownV2 mode. It escapes all special symbols
// used in MarkdownV2 syntax including: _, *, [, ], (, ), ~, `, >, #, +, -, =, |, {, }, ., and !.
// This ensures that the text is displayed as-is without being interpreted as MarkdownV2 formatting.
func escapeMarkdownV2(text string) string {
	return strings.NewReplacer(
		"_", "\\_", "*", "\\*", "[", "\\[", "]", "\\]", "(", "\\(",
		")", "\\)", "~", "\\~", "`", "\\`", ">", "\\>", "#", "\\#",
		"+", "\\+", "-", "\\-", "=", "\\=", "|", "\\|", "{", "\\{",
		"}", "\\}", ".", "\\.", "!", "\\!",
	).Replace(text)
}

// escapeURL escapes a URL string according to the specified formatting mode.
// In Markdown and MarkdownV2 modes, it escapes backslashes and parentheses
// to prevent them from being interpreted as part of the markup syntax.
// In HTML mode, it uses HTML escaping to ensure the URL is safe for inclusion
// in HTML content. The function returns the escaped URL string.
func escapeURL(mode, u string) string {
	switch mode {
	case ModeMarkdown:
		u = strings.ReplaceAll(u, "\\", "\\\\")
		u = strings.ReplaceAll(u, ")", "\\)")
	case ModeMarkdownV2:
		u = strings.ReplaceAll(u, "\\", "\\\\")
		u = strings.ReplaceAll(u, ")", "\\)")
		u = strings.ReplaceAll(u, "(", "\\(")
	case ModeHTML:
		u = html.EscapeString(u)
	}
	return u
}

// DebugDump returns a string with title, length and a 60 character preview of input string s.
// Useful for logging/debugging.
func DebugDump(title string, s string) string {
	return fmt.Sprintf("[%s] len=%d head=%.60q", title, len(s), s)
}

// EscapeText takes an input text and escape Telegram markup symbols.
// In this way we can send a text without being afraid of having to escape the characters manually.
// Note that you don't have to include the formatting style in the input text, or it will be escaped too.
// If there is an error, an empty string will be returned.
//
// parseMode is the text formatting mode (ModeMarkdown, ModeMarkdownV2 or ModeHTML)
// text is the input string that will be escaped
func EscapeText(parseMode string, text string) string {
	var replacer *strings.Replacer

	switch parseMode {
	case ModeHTML:
		replacer = strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")
	case ModeMarkdown:
		replacer = strings.NewReplacer("_", "\\_", "*", "\\*", "", "\\", "[", "\\[")
	case ModeMarkdownV2:
		replacer = strings.NewReplacer(
			"_", "\\_", "*", "\\*", "[", "\\[", "]", "\\]", "(", "\\(",
			")", "\\)", "~", "\\~", "", "\\", ">", "\\>",
			"#", "\\#", "+", "\\+", "-", "\\-", "=", "\\=", "|",
			"\\|", "{", "\\{", "}", "\\}", ".", "\\.", "!", "\\!",
		)
	default:
		return ""
	}

	return replacer.Replace(text)
}
