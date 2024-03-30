package main

import (
	"fmt"
	"io"
	"time"
)

type Logger struct {
	writers []io.Writer
}

func NewLogger(writers ...io.Writer) Logger {
	return Logger{writers: writers}
}

func (l Logger) getFormattedTime() string {
	return time.Now().Format("2006-01-02T15:04:05.000Z07:00")
}

func (l Logger) generateLog(s string) string {
	return fmt.Sprintf("%s - %s\n", l.getFormattedTime(), s)
}

func (l Logger) Log(s string) {
	log := l.generateLog(s)
	multi := io.MultiWriter(l.writers...)
	fmt.Fprint(multi, log)
}
