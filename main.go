package main

import (
	"flag"
	"fmt"
	"log"
	"tg-bot-urls/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	tgClient := telegram.New(tgBotHost, mustToken())

	fmt.Println(tgClient)

	// fetcher = fetcher.New()
	// processor = processor.New(tgClient)
	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("t", "", "telegram bot token")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not found")
	}
	return *token
}
