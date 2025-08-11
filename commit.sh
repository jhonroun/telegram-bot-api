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

# Запуск ssh-agent и добавление ключа
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa

# --- 1) найти актуальную версию по файлам api_coverage_X.Y(.Z)_test.go ---
LATEST_VER=$(
  find . -maxdepth 1 -type f -name 'api_coverage_*_test.go' \
  | sed -E 's|.*/api_coverage_([0-9]+(\.[0-9]+){1,2})_test\.go$|\1|' \
  | grep -E '^[0-9]+(\.[0-9]+){1,2}$' \
  | sort -V | tail -n 1
)
[[ -n "${LATEST_VER:-}" ]] || die "no files like api_coverage_X.Y(.Z)_test.go found"
echo "✅ latest Bot API version: ${LATEST_VER}"

# --- сохранить оригинал doc.go для отката при ошибке ---
[[ -f "$DOC_FILE" ]] || die "$DOC_FILE not found"
DOC_BAK="$(mktemp)"
cp "$DOC_FILE" "$DOC_BAK"
cleanup() { rm -f "$DOC_BAK"; }
restore_doc(){ cp "$DOC_BAK" "$DOC_FILE"; cleanup; }
trap 'restore_doc' ERR

# --- 1a) обновить строку 'Actual Bot API Version:' в doc.go ---
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

# --- 2) source .env ---
if [[ -f .env ]]; then
  set -a
  # shellcheck disable=SC1091
  source .env
  set +a
  echo "🔧 .env loaded"
else
  echo "ℹ️  .env not found, skipping"
fi

# --- 3) запуск тестов '^Test<MAJOR><MINOR>_.*' ---
MAJOR=$(echo "$LATEST_VER" | cut -d. -f1)
MINOR=$(echo "$LATEST_VER" | cut -d. -f2)
[[ -n "$MAJOR" && -n "$MINOR" ]] || die "cannot parse MAJOR/MINOR from ${LATEST_VER}"
TEST_REGEX="^Test${MAJOR}${MINOR}_.*"

echo "🧪 go test -v -run '${TEST_REGEX}' ./..."
go test -v -run "${TEST_REGEX}" ./...
echo "✅ tests passed for ${LATEST_VER}"

# дальше откат doc.go не нужен
trap - ERR
cleanup

# --- 4) перегенерация документации ---
echo "📚 generating docs with gomarkdoc..."
gomarkdoc --output '{{.Dir}}/README.md' ./...
echo "✅ docs updated"

# --- 4a) вставить/обновить фиксированный блок в начале README.md ---
ROOT_README="README.md"
[[ -f "$ROOT_README" ]] || touch "$ROOT_README"

INTRO_START="<!-- BEGIN README INTRO -->"
INTRO_END="<!-- END README INTRO -->"

INTRO_BLOCK=$(cat <<'EOF'
<!-- BEGIN README INTRO -->
# Golang bindings for the Telegram Bot API

[![Go Reference](https://pkg.go.dev/badge/github.com/jhonroun/telegram-bot-api.svg)](https://pkg.go.dev/github.com/jhonroun/telegram-bot-api) 
[![Test](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml/badge.svg)](https://github.com/jhonroun/telegram-bot-api/actions/workflows/ci.yml)

The repo was created to study and check the relevance of the module for working with the Bot API \(https://github.com/go-telegram-bot-api/telegram-bot-api\), which is called step-by-step. Many thanks to the author for the awesome experience and idea. Initially, I wanted to create a tool for writing modern bots. But in the process of adding functionality, I thought that I was writing it for myself first and foremost. There are quite enough forms with an updated version of the Bot API on github.com.

From now on, the abandonment of versioning like v0.\*/v1.\*
<!-- END README INTRO -->
EOF
)

# удаляем старый блок, если есть
TMP_README="$(mktemp)"
awk -v s="$INTRO_START" -v e="$INTRO_END" '
  BEGIN{skip=0}
  index($0,s){skip=1; next}
  index($0,e){skip=0; next}
  !skip{print}
' "$ROOT_README" > "$TMP_README"

# добавляем новый блок в начало
{
  printf "%s\n\n" "$INTRO_BLOCK"
  cat "$TMP_README"
} > "$ROOT_README"

rm -f "$TMP_README"
echo "🧷 README intro block updated"

# --- 5) коммитим дифф и пушим ---
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

Diff vs ${REMOTE}/${MAIN_BRANCH}:

${DIFFSTAT}

Files:
${NAMEDIFF}
EOF
)

git commit -m "$COMMIT_TITLE" -m "$COMMIT_BODY"
git push "$REMOTE" HEAD:"$MAIN_BRANCH"
echo "🚀 pushed to ${REMOTE}/${MAIN_BRANCH}"
