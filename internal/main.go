package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sashabaranov/go-openai"
)

func Run(ctx context.Context, openAiKey string) {
	if openAiKey == "" {
		shError("Error", fmt.Errorf("OpenAI key missing"))
		return
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go shutdown(signalChannel)

	c := openai.NewClient(openAiKey)
	model := openai.GPT3Dot5Turbo
	session := startSession()

	for {
		userInput := shAsk("", false)

		if userInput == "new-session" {
			session = startSession()
			continue
		}

		session = append(session, userMsg(userInput))
		req := openai.ChatCompletionRequest{Model: model, Stream: true, Messages: session}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			shError("Error", fmt.Errorf("failed to connect to OpenAi: %w", err))
			session = session[:len(session)-1] // remove last question from session
			continue
		}
		defer stream.Close()

		fmt.Println(shColor("gray", "---"))

		fullResponse := ""
		for {
			response, err := stream.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					shError("Error", fmt.Errorf("failed getting response: %w", err))
				}
				break
			}

			fullResponse += response.Choices[0].Delta.Content
			fmt.Print(response.Choices[0].Delta.Content)
		}

		session = append(session, assistantMsg(fullResponse))

		fmt.Print("\n\n")
		fmt.Println(shColor("whitesmoke", "- Tokens: %d -", countTokens(session, model)), shColor("gray", "   (type 'new-session' to clear context | CTRL+C to quit)"))
		fmt.Println("")
	} // loop
}

// Helpers ------------------------------------

func startSession() []openai.ChatCompletionMessage {
	shClear()
	fmt.Println()
	fmt.Println(shColor("yellow:bold", "What do you want to talk about?"), shColor("whitesmoke", "   (type CTRL+C to quit)"))

	introMsg := systemMsg(strings.TrimSpace(`
You have access to most of human knowledge and are very pleased to assist in uncovering the correct and detailed answers for the topics you're asked about.
User will ask questions seeking concise and accurate responses
	`))

	return []openai.ChatCompletionMessage{introMsg}
}

func systemMsg(q string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: q}
}

func userMsg(q string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: q}
}

func assistantMsg(q string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: q}
}

func shutdown(signalChannel chan os.Signal) {
	<-signalChannel

	fmt.Println()
	fmt.Println(shColor("gray", "-------------"))
	fmt.Println(shColor("whitesmoke", "Thank you for using ChatGo"))
	fmt.Println("Created by", shColor("green", "Alex Parra · github.com/alex-parra"))

	os.Exit(0)
}
