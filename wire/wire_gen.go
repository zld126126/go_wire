// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package wire

// Injectors from wire.go:

func InitializeEvent(msg string) Event {
	message := NewMessage(msg)
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event
}
