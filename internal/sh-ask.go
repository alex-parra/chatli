package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func shAsk(def string, secret bool) string {
	input := ""
	if secret {
		b, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err == nil {
			input = string(b)
		}
		fmt.Println()

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		ok := scanner.Scan()
		if ok {
			input = strings.TrimRight(scanner.Text(), "\r\n")
		}
	}

	if input == "" {
		input = def
	}

	return input
}
