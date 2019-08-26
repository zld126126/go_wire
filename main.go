package main

import "awesomeProject/wire"

func main() {
	// 使用wire前
	message := wire.NewMessage("default wire-使用wire前")
	greeter := wire.NewGreeter(message)
	event1 := wire.NewEvent(greeter)
	event1.Start()

	// 使用wire后
	event2 := wire.InitializeEvent("my_wire-使用wire后")
	event2.Start()
}
