package main

import (
	"github.com/siddlal/notification/config"
	"github.com/siddlal/notification/decorator"
	"github.com/siddlal/notification/strategy"
)

func main() {
	cfg := config.GetConfig()
	println("App:", cfg.AppName, "| Version:", cfg.Version)
	var notifier strategy.Notifier = strategy.EmailNotifier{}
	loggedNotifier := decorator.LoggerDecorator{Notifier: notifier}

	loggedNotifier.Send("Hello SIddharth! your order has been shipped")
}
