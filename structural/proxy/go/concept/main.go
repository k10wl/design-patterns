package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

type Logger interface {
	Println(...any)
}

type ProxyLogger struct {
	logger Logger
}

func (p ProxyLogger) Println(args ...any) {
	s := args[0].(string)
	if ok, _ := regexp.MatchString("[0-5]", s); ok {
		p.logger.Println("can't log up to 5")
		return
	}
	p.logger.Println(s)
}

func main() {
	fmt.Println(">>> Proxy demo")
	logger := log.New(os.Stdout, "[empty]", 0)
	proxy := ProxyLogger{logger: logger}
	fmt.Println("\n>>> Using proxy")
	makeSomeLogs(proxy)
	fmt.Println("\n>>> Using real")
	makeSomeLogs(logger)
}

func makeSomeLogs(logger Logger) {
	for i := 0; i < 10; i++ {
		logger.Println(fmt.Sprintf("making log #%d", i))
	}
}
