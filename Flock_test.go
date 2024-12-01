package flock

import "fmt"

func ExampleFlock() {
	var f1 = New("./Flock_test.go")
	var f2 = New("./Flock_test.go")
	f1.Lock()
	f1.Close()
	f2.Lock()
	f2.Close()
	fmt.Println("finished")
	// Output:
	// finished
}
