package internal

import (
	"log"
	"strings"

	"github.com/fatih/color"
)

// shColor adds ansi escape codes to colorize shell output
func shColor(fx, str string, a ...any) string {
	if color.NoColor {
		return str
	}

	opts := strings.Split(fx, ":")
	colorName := sliceAt(opts, 0, "reset")
	effect := sliceAt(opts, 1, "")

	whiteSmoke := func(s string, a ...any) string {
		if len(a) > 0 {
			s = sf(s, a...)
		}
		return "\033[38;2;180;180;180m" + s + "\033[39m"
	}

	gray := func(s string, a ...any) string {
		if len(a) > 0 {
			s = sf(s, a...)
		}
		return "\033[38;2;85;85;85m" + s + "\033[39m"
	}

	colors := map[string]func(s string, a ...any) string{
		"red":        color.RedString,
		"green":      color.GreenString,
		"yellow":     color.YellowString,
		"blue":       color.BlueString,
		"purple":     color.MagentaString,
		"cyan":       color.CyanString,
		"white":      color.WhiteString,
		"whitesmoke": whiteSmoke,
		"gray":       gray,
		"reset":      sf,
	}

	if effect == "bold" {
		str = "\033[1m" + str + "\033[22m"
	}

	if _, ok := colors[colorName]; !ok {
		log.Printf("WARN: unsupported color '%s'", colorName)
		colorName = "reset"
	}

	return colors[colorName](str, a...)
}
