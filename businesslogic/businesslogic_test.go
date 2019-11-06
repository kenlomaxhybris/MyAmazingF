package businesslogic

import (
	"fmt"
	"testing"
)

func checkMyLogic(s string, t *testing.T) {
	t.Helper()
	if s != "Hello World" {
		t.Error("Output is not what we expected")
	}
}

func TestF(t *testing.T) {
	var s = F()
	checkMyLogic(s, t)
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		F()
	}
}

func ExampleF() {
	// godoc -http ":8081"
	// go test ./... -cover -race -v
	// I would like a function F that returns "Hello World"
	fmt.Println(F())

	// Output:
	// Hello World
}
