package strategy

import "fmt"

type SmsNotifier struct {
}

func (SmsNotifier) Send(message string) {
	fmt.Println("sending Message: ", message)

}
