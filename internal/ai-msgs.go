package internal

import (
	"strings"

	"github.com/sashabaranov/go-openai"
)

const promptPrimer = `
You have access to most of human knowledge and are very pleased to assist in uncovering the correct and detailed answers for the topics you're asked about.
User will ask questions seeking concise and accurate responses
`

type ChatMsg struct {
	msg    openai.ChatCompletionMessage
	role   string
	tokens int
}

func introMsg(model string) ChatMsg {
	return systemMsg(model, strings.TrimSpace(promptPrimer))
}

func systemMsg(model string, q string) ChatMsg {
	return chatMsg(model, openai.ChatMessageRoleSystem, q)
}

func userMsg(model string, q string) ChatMsg {
	return chatMsg(model, openai.ChatMessageRoleUser, q)
}

func agentMsg(model string, q string) ChatMsg {
	return chatMsg(model, openai.ChatMessageRoleAssistant, q)
}

func chatMsg(model string, role string, content string) ChatMsg {
	msg := openai.ChatCompletionMessage{Role: role, Content: strings.TrimSpace(content)}
	tokens := countTokens([]openai.ChatCompletionMessage{msg}, model)
	return ChatMsg{msg: msg, role: role, tokens: tokens}
}

func chatMsgsToOpenAIMsgs(msgs []ChatMsg) []openai.ChatCompletionMessage {
	oaMsgs := make([]openai.ChatCompletionMessage, len(msgs))

	for i, chatMsg := range msgs {
		oaMsgs[i] = chatMsg.msg
	}

	return oaMsgs
}

func countSessionTokens(session []ChatMsg) int {
	tokens := 0
	for _, m := range session {
		tokens += m.tokens
	}
	return tokens
}
