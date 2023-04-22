package main

import (
	"context"
	"log"
	"os/user"

	"github.com/alex-parra/chatli/internal"
)

func main() {
	ctx := context.Background()

	internal.Startup()

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	openAIKey, err := internal.GetOpenAIKey(currentUser.Uid)
	if err != nil {
		log.Fatal("Failed getting OpenAI key")
	}

	internal.Chat(ctx, openAIKey)
}
