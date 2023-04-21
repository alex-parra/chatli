package internal

import "fmt"

func shError(prefix string, err error) {
	fmt.Println(shColor("red", prefix+":"), err)
}
