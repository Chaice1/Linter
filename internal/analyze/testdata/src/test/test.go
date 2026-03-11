package test

import (
	"log/slog"
)

func check() {
	password := "123"

	slog.Debug("DDDDDDDD") // want "First symbol should be low" "log message can contain only a-z 0-9 and basic punctuation symbols"

	slog.Info("m!")                        // want "log message can contain only a-z 0-9 and basic punctuation symbols"
	slog.Info("password", "val", password) // want "log message can't containt this keyword password" "logging variable with sensitive name 'password'"
	slog.Info("token")                     // want "log message can't containt this keyword token"
}
