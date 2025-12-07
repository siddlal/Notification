package decorator

import (
	"fmt"

	"github.com/siddlal/notification/strategy"
)

type LoggerDecorator struct {
	Notifier strategy.Notifier
}

func (d LoggerDecorator) Send(message string) {
	fmt.Println("LOG sending notification: ")
	d.Notifier.Send(message)
	fmt.Println("LOG: notification sent")
}
