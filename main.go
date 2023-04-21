package main

import (
	"chat-go/internal"
	"context"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	openAiKey := os.Getenv("OPENAI_API_KEY")
	ctx := context.Background()

	internal.Run(ctx, openAiKey)
}
