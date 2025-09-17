package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sungales/gobot/bot"
)

func main() {

	 err := godotenv.Load()
	 if err != nil {
		log.Fatal("Error loading .env file")
	 }
	 botToken := os.Getenv("BOT_TOKEN")

	gobot.BotToken = botToken
	gobot.Run()
}