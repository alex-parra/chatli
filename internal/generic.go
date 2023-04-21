package internal

import "fmt"

// sf is short-hand for string format and calls fmt.Sprintf if arguments are passed
func sf(str string, args ...any) string {
	if len(args) == 0 {
		return str
	}
	return fmt.Sprintf(str, args...)
}
