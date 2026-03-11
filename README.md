# Linter

## Описание
Линтер для проверки правил логирования в Go (slog, zap, log).

## Правила
- Регистр: только строчные буквы в начале.
- Язык: только английский.
- Символы: запрещены эмодзи и спецсимволы.
- Security: запрещено логировать пароли и токены.

## Команды
- Клонирование репозитория:
```bash
git clone https://github.com/Chaice1/Linter.git
cd Linter
```
- Сборка:
```bash
go build -o loglinter.exe ./cmd/linter/main.go
```
- Тесты:
```bash
go test -v ./internal/analyze/...
```
- Запуск:
```bash
.\loglinter.exe ./internal/analyze/testdata/src/test
 ```
