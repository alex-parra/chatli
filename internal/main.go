package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/sashabaranov/go-openai"
)

const cmdNewSession = "/new"
const cmdQuit = "/q"

func Startup() {
	shClear()
	fmt.Println()
	fmt.Println(shColor("yellow:bold", Config.AppName))
	fmt.Println("Created by", shColor("green", Config.CreatedBy), "·", Config.CreatedByLinks[0], "·", Config.CreatedByLinks[1])
	fmt.Println()
	fmt.Println(shColor("gray", "-------------"))
	fmt.Println()
}

func Chat(ctx context.Context, openAIKey string) {
	if openAIKey == "" {
		shError("Error", fmt.Errorf("OpenAI key missing"))
		return
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() { <-signalChannel; shutdown() }()

	c := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	session := startSession(model)

	for {
		userInput := shAsk("", false)

		if userInput == cmdQuit {
			shutdown()
		}

		if userInput == cmdNewSession {
			session = startSession(model)
			continue
		}

		session = append(session, userMsg(model, userInput))
		req := openai.ChatCompletionRequest{Model: model, Stream: true, Messages: chatMsgsToOpenAIMsgs(session)}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			shError("Error", fmt.Errorf("failed to connect to OpenAI: %w", err))
			session = session[:len(session)-1] // remove last question from session
			continue
		}
		defer stream.Close()

		fmt.Println(shColor("gray", "---"))

		completeResponse := ""
		for {
			res, err := stream.Recv()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					shError("Error", fmt.Errorf("failed getting response: %w", err))
				}
				break
			}

			completeResponse += res.Choices[0].Delta.Content
			fmt.Print(res.Choices[0].Delta.Content)
		}

		session = append(session, agentMsg(model, completeResponse))
		tokens := countSessionTokens(session)

		fmt.Print("\n\n")
		fmt.Println(shColor("whitesmoke", "- Tokens: %d -", tokens), "  ", shColor("gray", "(enter '%s' to clear context or '%s' to quit)", cmdNewSession, cmdQuit))

		if Config.ModelMaxTokens[model]-tokens <= Config.SessionTokensThreshold {
			fmt.Println(" ", shColor("gray", "Reaching max tokens: deleting the oldest message from context."))
			session = append(session[0:1], session[2:]...)
		}

		fmt.Println("")
	} // loop
}

// Helpers ------------------------------------

func startSession(model string) []ChatMsg {
	shClear()
	fmt.Println()
	fmt.Println("👋 Welcome to", Config.AppName)
	fmt.Println(shColor("gray", "-------------"))
	fmt.Println(shColor("yellow:bold", "What do you want to talk about?"), "  ", shColor("gray", "(enter '%s' to quit)", cmdQuit))

	return []ChatMsg{introMsg(model)}
}

func shutdown() {
	fmt.Println()
	fmt.Println(shColor("gray", "-------------"))
	fmt.Println(shColor("whitesmoke", "Thank you for using %s", Config.AppName))
	fmt.Println("Created by", shColor("green", Config.CreatedBy), "·", Config.CreatedByLinks[0], "·", Config.CreatedByLinks[1])

	os.Exit(0)
}
