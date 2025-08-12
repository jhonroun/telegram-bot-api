#!/usr/bin/env bash
set -euo pipefail

DOC_FILE="doc.go"
MAIN_BRANCH="main"
REMOTE="origin"

die(){ echo "ERROR: $*" >&2; exit 1; }
need(){ command -v "$1" >/dev/null 2>&1 || die "missing '$1'"; }

need git; need go; need sed; need awk; need grep; need sort

if ! command -v gomarkdoc >/dev/null 2>&1; then
  die "gomarkdoc not found. install: go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest"
fi

eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519

LATEST_VER=$(
  find . -maxdepth 1 -type f -name 'api_coverage_*_test.go' \
  | sed -E 's|.*/api_coverage_([0-9]+(\.[0-9]+){1,2})_test\.go$|\1|' \
  | grep -E '^[0-9]+(\.[0-9]+){1,2}$' \
  | sort -V | tail -n 1
)
[[ -n "${LATEST_VER:-}" ]] || die "no files like api_coverage_X.Y(.Z)_test.go found"
echo "✅ latest Bot API version: ${LATEST_VER}"

[[ -f "$DOC_FILE" ]] || die "$DOC_FILE not found"
DOC_BAK="$(mktemp)"
cp "$DOC_FILE" "$DOC_BAK"
cleanup() { rm -f "$DOC_BAK"; }
restore_doc(){ cp "$DOC_BAK" "$DOC_FILE"; cleanup; }
trap 'restore_doc' ERR

awk -v ver="$LATEST_VER" '
  {
    if ($0 ~ /^[[:space:]]*\/\/[[:space:]]*Actual Bot API Version:/) {
      sub(/:.*/, ": " ver)
    }
    print
  }
' "$DOC_FILE" > "${DOC_FILE}.tmp"
mv "${DOC_FILE}.tmp" "$DOC_FILE"
echo "📝 doc.go updated → ${LATEST_VER}"

if [[ -f .env ]]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
  echo "🔧 .env loaded"
else
  echo "ℹ️  .env not found, skipping"
fi

MAJOR=$(echo "$LATEST_VER" | cut -d. -f1)
MINOR=$(echo "$LATEST_VER" | cut -d. -f2)
[[ -n "$MAJOR" && -n "$MINOR" ]] || die "cannot parse MAJOR/MINOR from ${LATEST_VER}"
TEST_REGEX="^Test${MAJOR}${MINOR}_.*"

echo "🧪 go test -v -run '${TEST_REGEX}' ./..."
go test -v -run "${TEST_REGEX}" ./...
echo "✅ tests passed for ${LATEST_VER}"

trap - ERR
cleanup

echo "📚 generating docs with gomarkdoc..."
gomarkdoc --output '{{.Dir}}/README.md' ./...
echo "✅ docs updated"

ROOT_README="README.md"
[[ -f "$ROOT_README" ]] || touch "$ROOT_README"

INTRO_START="<!-- BEGIN README INTRO -->"
INTRO_END="<!-- END README INTRO -->"

INTRO_BLOCK=$(cat <<'EOF'
<!-- BEGIN README INTRO -->
# Golang bindings for the Telegram Bot API v__MAJOR__.__MINOR__

[![Go Reference](https://pkg.go.dev/badge/github.com/jhonroun/telegram-bot-api.svg)](https://pkg.go.dev/github.com/jhonroun/telegram-bot-api) 
[![Test](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml/badge.svg)](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml)

The repo was created to study and check the relevance of the module for working with the Bot API [original module](https://github.com/go-telegram-bot-api/telegram-bot-api), which is called step-by-step. Many thanks to the author for the awesome experience and idea. Initially, I wanted to create a tool for writing modern bots. But in the process of adding functionality, I thought that I was writing it for myself first and foremost. There are quite enough forms with an updated version of the Bot API on github.com.

## Some intresting things..

 - From now the abandonment of versioning like v0/v1.
 - Added AST for render makdown messages. Example see at [murkups_test](https://github.com/jhonroun/telegram-bot-api/blob/main/markups_test.go).
  For render programming language code block added const list of supported languages. Supported list see at [programming languages supported](https://github.com/jhonroun/telegram-bot-api/blob/main/programm_laguage.go).
 - Added supports for create Invoice and supported currencies. See at [test api coverage](https://github.com/jhonroun/telegram-bot-api/blob/main/api_coverage_6.1.0_test.go). [more info in Telegram Docs](https://core.telegram.org/bots/payments). Example usage at [test api coverage](https://github.com/jhonroun/telegram-bot-api/blob/main/api_coverage_6.1.0_test.go).
 - Added supports of protected content. At Bot API v5.7.
 - Added supports of WebM sticker. See at [test api coverage](https://github.com/jhonroun/telegram-bot-api/blob/main/api_coverage_5.7.0_test.go)
 - Added new method bot.GetUserIDbyUsername(username). This method obtain ChatID by username. 
 Example:
 `bot.GetUserIDbyUsername("@tggobotapitest")`
 - Now chatID in construct New... (helpers) can be int64, BaseChat, ChatConfig, ChatActionConfig, Chat, User
 - Added universal methods sendChatAction and short form. Example usage:
 `_, _ = bot.Typing(chatID)`

## NOTE: It's not obvious:

 - message_thread_id as part of BaseChat type because sendMessage, sendPhoto, sendVideo, sendAnimation, sendAudio, sendDocument, sendSticker, sendVideoNote, sendVoice, sendLocation, sendVenue, sendContact, sendDice, sendInvoice, sendGame, copyMessage, forwardMessage - used it.
And addition in type MediaGroupConfig.

# Getting Started

This library is designed as a simple wrapper around the Telegram Bot API.
It's encouraged to read [Telegram's docs][telegram-docs] first to get an
understanding of what Bots are capable of doing. They also provide some good
approaches to solve common problems.

[telegram-docs]: https://core.telegram.org/bots

## Installing

```bash
go get -u github.com/jhonroun/telegram-bot-api
```

## A Simple Bot

To walk through the basics, let's create a simple echo bot that replies to your
messages repeating what you said. Make sure you get an API token from
[@Botfather][botfather] before continuing.

Let's start by constructing a new [BotAPI].

[botfather]: https://t.me/Botfather

```go
package main

import (
	"os"

	tgbotapi "github.com/jhonroun/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
```

Instead of typing the API token directly into the file, we're using
environment variables. This makes it easy to configure our Bot to use the right
account and prevents us from leaking our real token into the world. Anyone with
your token can send and receive messages from your Bot!

We've also set `bot.Debug = true` in order to get more information about the
requests being sent to Telegram. If you run the example above, you'll see
information about a request to the [`getMe`][get-me] endpoint. The library
automatically calls this to ensure your token is working as expected. It also
fills in the `Self` field in your `BotAPI` struct with information about the
Bot.

Now that we've connected to Telegram, let's start getting updates and doing
things. We can add this code in right after the line enabling debug mode.

[get-me]: https://core.telegram.org/bots/api#getme

```go
	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		// Now that we know we've gotten a new message, we can construct a
		// reply! We'll take the Chat ID and Text from the incoming message
		// and use it to create a new message.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// We'll also say that this message is a reply to the previous message.
		// For any other specifications than Chat ID or Text, you'll need to
		// set fields on the `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
	}
```

Congratulations! You've made your very own bot!

Now that you've got some of the basics down, we can start talking about how the
library is structured and more advanced features.

# Library Structure

This library is generally broken into three components you need to understand.

## Configs

Configs are collections of fields related to a single request. For example, if
one wanted to use the `sendMessage` endpoint, you could use the `MessageConfig`
struct to configure the request. There is a one-to-one relationship between
Telegram endpoints and configs. They generally have the naming pattern of
removing the `send` prefix and they all end with the `Config` suffix. They
generally implement the `Chattable` interface. If they can send files, they
implement the `Fileable` interface.

## Helpers

Helpers are easier ways of constructing common Configs. Instead of having to
create a `MessageConfig` struct and remember to set the `ChatID` and `Text`,
you can use the `NewMessage` helper method. It takes the two required parameters
for the request to succeed. You can then set fields on the resulting
`MessageConfig` after it's creation. They are generally named the same as
method names except with `send` replaced with `New`.

## Methods

Methods are used to send Configs after they are constructed. Generally,
`Request` is the lowest level method you'll have to call. It accepts a
`Chattable` parameter and knows how to upload files if needed. It returns an
`APIResponse`, the most general return type from the Bot API. This method is
called for any endpoint that doesn't have a more specific return type. For
example, `setWebhook` only returns `true` or an error. Other methods may have
more specific return types. The `getFile` endpoint returns a `File`. Almost
every other method returns a `Message`, which you can use `Send` to obtain.

There's lower level methods such as `MakeRequest` which require an endpoint and
parameters instead of accepting configs. These are primarily used internally.
If you find yourself having to use them, please open an issue.

# Important Notes

The Telegram Bot API has a few potentially unanticipated behaviors. Here are a
few of them. If any behavior was surprising to you, please feel free to open a
pull request!

## Callback Queries

- Every callback query must be answered, even if there is nothing to display to
  the user. Failure to do so will show a loading icon on the keyboard until the
  operation times out.

## ChatMemberUpdated

- In order to receive `ChatMember` updates, you must explicitly add
  `UpdateTypeChatMember` to your `AllowedUpdates` when getting updates or
  setting your webhook.

## Entities use UTF16

- When extracting text entities using offsets and lengths, characters can appear
  to be in incorrect positions. This is because Telegram uses UTF16 lengths
  while Golang uses UTF8. It's possible to convert between the two.

## GetUpdatesChan

- This method is very basic and likely unsuitable for production use. Consider
  creating your own implementation instead, as it's very simple to replicate.
- This method only allows your bot to process one update at a time. You can
  spawn goroutines to handle updates concurrently or switch to webhooks instead.
  Webhooks are suggested for high traffic bots.

## Nil Updates

- At most one of the fields in an `Update` will be set to a non-nil value. When
  evaluating updates, you must make sure you check that the field is not nil
  before trying to access any of it's fields.

## Privacy Mode

- By default, bots only get updates directly addressed to them. If you need to
  get all messages, you must disable privacy mode with Botfather. Bots already
  added to groups will need to be removed and re-added for the changes to take
  effect. You can read more on the [Telegram Bot API docs][api-docs].

[api-docs]: https://core.telegram.org/bots/faq#what-messages-will-my-bot-get

## User and Chat ID size

- These types require up to 52 significant bits to store correctly, making a
  64-bit integer type required in most languages. They are already `int64` types
  in this library, but make sure you use correct types when saving them to a
  database or passing them to another language.

# Files

Telegram supports specifying files in many different formats. In order to
accommodate them all, there are multiple structs and type aliases required.

All of these types implement the `RequestFileData` interface.

| Type         | Description                                                               |
| ------------ | ------------------------------------------------------------------------- |
| `FilePath`   | A local path to a file                                                    |
| `FileID`     | Existing file ID on Telegram's servers                                    |
| `FileURL`    | URL to file, must be served with expected MIME type                       |
| `FileReader` | Use an `io.Reader` to provide a file. Lazily read to save memory.         |
| `FileBytes`  | `[]byte` containing file data. Prefer to use `FileReader` to save memory. |

## `FilePath`

A path to a local file.

```go
file := tgbotapi.FilePath("tests/image.jpg")
```

## `FileID`

An ID previously uploaded to Telegram. IDs may only be reused by the same bot
that received them. Additionally, thumbnail IDs cannot be reused.

```go
file := tgbotapi.FileID("AgACAgIAAxkDAALesF8dCjAAAa_…")
```

## `FileURL`

A URL to an existing resource. It must be served with a correct MIME type to
work as expected.

```go
file := tgbotapi.FileURL("https://i.imgur.com/unQLJIb.jpg")
```

## `FileReader`

Use an `io.Reader` to provide file contents as needed. Requires a filename for
the virtual file.

```go
var reader io.Reader

file := tgbotapi.FileReader{
    Name: "image.jpg",
    Reader: reader,
}
```

## `FileBytes`

Use a `[]byte` to provide file contents. Generally try to avoid this as it
results in high memory usage. Also requires a filename for the virtual file.

```go
var data []byte

file := tgbotapi.FileBytes{
    Name: "image.jpg",
    Bytes: data,
}
```

Another Examples see into api_coverage tests.

<!-- END README INTRO -->
EOF
)

INTRO_BLOCK=$(printf '%s' "$INTRO_BLOCK" \
  | sed -e "s/__MAJOR__/${MAJOR}/g" \
        -e "s/__MINOR__/${MINOR}/g")

TMP_README="$(mktemp)"
awk -v s="$INTRO_START" -v e="$INTRO_END" '
  BEGIN{skip=0}
  index($0,s){skip=1; next}
  index($0,e){skip=0; next}
  !skip{print}
' "$ROOT_README" > "$TMP_README"

{
  printf "%s\n\n" "$INTRO_BLOCK"
  cat "$TMP_README"
} > "$ROOT_README"

rm -f "$TMP_README"
echo "🧷 README intro block updated"

CI_FILE=".github/workflows/ci.yml"
if [[ -f "$CI_FILE" ]]; then
  CI_MASK="^Test_${MAJOR}${MINOR}_"

  sed -E -i.bak "s@(-run[[:space:]]*['\"])\\^Test_[0-9]+_@\1${CI_MASK}@g" "$CI_FILE" && rm -f "${CI_FILE}.bak"

  echo "🛠  CI test mask set to: ${CI_MASK}"
else
  echo "ℹ️  ${CI_FILE} not found, skipping CI mask update"
fi

git fetch "$REMOTE" "$MAIN_BRANCH" || true

if git diff --quiet && git diff --cached --quiet; then
  echo "ℹ️ nothing to commit."
  exit 0
fi

git add -A

DIFFSTAT=$(git diff --stat "${REMOTE}/${MAIN_BRANCH}...HEAD" || true)
NAMEDIFF=$(git diff --name-status "${REMOTE}/${MAIN_BRANCH}...HEAD" || true)

COMMIT_TITLE="Docs & tests: Bot API v${LATEST_VER} (autoupdate)"
COMMIT_BODY=$(
  cat <<EOF
- Update doc.go: Actual Bot API Version → ${LATEST_VER}
- Tests: go test -run '${TEST_REGEX}' ./...
- Regenerate docs with gomarkdoc
- Prepend README intro block
- Update CI test mask to '^Test_${MAJOR}${MINOR}_'

Diff vs ${REMOTE}/${MAIN_BRANCH}:

${DIFFSTAT}

Files:
${NAMEDIFF}
EOF
)

git commit -m "$COMMIT_TITLE" -m "$COMMIT_BODY"
git push "$REMOTE" HEAD:"$MAIN_BRANCH"
echo "🚀 pushed to ${REMOTE}/${MAIN_BRANCH}"
