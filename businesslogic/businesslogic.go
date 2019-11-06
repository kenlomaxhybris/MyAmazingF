package businesslogic

import "time"

// F is a function that returns "Hello World"
func F() string {
	time.Sleep(time.Millisecond * 50)
	return "Hello World"
}
