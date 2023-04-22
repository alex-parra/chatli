package internal

import "github.com/sashabaranov/go-openai"

type config struct {
	AppName        string
	CreatedBy      string
	CreatedByLinks []string

	SessionTokensThreshold int
	ModelMaxTokens         map[string]int
}

var Config config = config{
	AppName:        "Chatli",
	CreatedBy:      "Alex Parra",
	CreatedByLinks: []string{"linkedin.com/in/alexpds", "github.com/alex-parra"},

	SessionTokensThreshold: 500,
	ModelMaxTokens: map[string]int{
		openai.GPT3Dot5Turbo: 1096,
		openai.GPT4:          8192,
	},
}
