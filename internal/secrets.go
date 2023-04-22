package internal

import (
	"fmt"

	"github.com/zalando/go-keyring"
)

func GetOpenAIKey(storageKey string) (string, error) {
	service := Config.AppName + "-OpenAi-API-key"

	if storageKey == "" {
		return "", fmt.Errorf("storageKey is required")
	}

	openAiKey, err := keyring.Get(service, storageKey)
	if err != nil {
		for {
			fmt.Println("Enter your OpenAI key: ", "  ", shColor("gray", "key will be stored in your Keychain with name: %s", service))
			openAiKey = shAsk("", true)
			if openAiKey != "" {
				break
			}
			fmt.Println("OpenAI key can't be empty.")
		}

		err := keyring.Set(service, storageKey, openAiKey)
		if err != nil {
			return "", err
		}
	}

	return openAiKey, nil
}
