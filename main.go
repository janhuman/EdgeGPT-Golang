package main

import (
	"log"

	EdgeGPT "github.com/wildy2828/EdgeGPT-Golang/Res"
)

func callback(answer *EdgeGPT.Answer) {
	if answer.IsDone() {
		log.Println(answer.NumUserMessages(), answer.MaxNumUserMessages(), answer.Text())
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	bot := EdgeGPT.NewChatBot("config.json", []map[string]interface{}{}) // proxy deleted
	err := bot.Init()
	if err != nil {
		panic(err)
	}
	err = bot.Ask("give me a joke", EdgeGPT.Creative, callback)
	if err != nil {
		panic(err)
	}
	err = bot.Ask("It's not funny", EdgeGPT.Creative, callback)
	if err != nil {
		panic(err)
	}
}
