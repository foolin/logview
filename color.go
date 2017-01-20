package logview

import (
	"fmt"
	"runtime"
)

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

var (
	dateTimeFmt = "2006-01-02 15:04:05"
)

func Black(str string) string {
	return TextColor(TextBlack, str)
}

func Red(str string) string {
	return TextColor(TextRed, str)
}

func Green(str string) string {
	return TextColor(TextGreen, str)
}

func Yellow(str string) string {
	return TextColor(TextYellow, str)
}

func Blue(str string) string {
	return TextColor(TextBlue, str)
}

func Magenta(str string) string {
	return TextColor(TextMagenta, str)
}

func Cyan(str string) string {
	return TextColor(TextCyan, str)
}

func White(str string) string {
	return TextColor(TextWhite, str)
}

func TextColor(color int, str string) string {
	if IsWindows() {
		return str
	}

	switch color {
	case TextBlack:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlack, str)
	case TextRed:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextRed, str)
	case TextGreen:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextGreen, str)
	case TextYellow:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextYellow, str)
	case TextBlue:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextBlue, str)
	case TextMagenta:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextMagenta, str)
	case TextCyan:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextCyan, str)
	case TextWhite:
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", TextWhite, str)
	default:
		return str
	}
}

func IsWindows() bool {
	if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}