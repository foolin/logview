package logview

import (
	"strings"
	"os"
	"fmt"
)

var logColors = []int{
	TextRed,
	TextGreen,
	TextYellow,
	TextBlue,
	TextMagenta,
	TextCyan,
	TextWhite,
}

type LogWriter struct {
	idx int
	len int
}

func NewWriter() *LogWriter {
	return &LogWriter{
		idx: 0,
		len: len(logColors),
	}
}

func (l *LogWriter) Write(p []byte) (n int, err error) {
	if l.len <= 0 {
		l.len = len(logColors)
	}
	contents := string(p)
	lines := strings.Split(contents, "\n")
	for _, v := range lines {
		v := strings.TrimSpace(v)
		if v == ""{
			continue
		}
		l.idx = (l.idx + 1) % l.len
		color := logColors[l.idx]
		if color != TextWhite{
			v = TextColor(color, v)
		}
		os.Stdout.WriteString(fmt.Sprintf("%v\n", v))
	}
	return len(p), nil
}