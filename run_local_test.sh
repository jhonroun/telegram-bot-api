# Подхватим .env если есть
if [ -f .env ]; then set -a; source .env; set +a; fi

# Проверка
go vet ./...

# Тесты с/без race в зависимости от архитектуры
ARCH=$(go env GOARCH)
if [ "$ARCH" = "arm64" ] || [ "$ARCH" = "arm" ]; then
  go test ./... -count=1
else
  go test ./... -race -count=1
fi