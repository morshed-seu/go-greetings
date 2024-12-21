package greetings

import "fmt"

func Hello(name String) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}


