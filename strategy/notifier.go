package strategy

type Notifier interface {
	Send(message string)
}
