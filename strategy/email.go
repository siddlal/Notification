package strategy

import (
	"fmt"
)

type EmailNotifier struct {
}

func (EmailNotifier) Send(message string) {
	fmt.Println("sending Email: ", message)
}
