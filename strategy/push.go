package strategy

import "fmt"

type PushNotifier struct {
}

func (PushNotifier) Send(message string) {
	fmt.Println("Sending Push: ", message)
}
