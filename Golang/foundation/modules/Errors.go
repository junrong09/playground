package modules

import "fmt"

// Basically, "error" type is just a interface as follows:
//type error interface {
//	Error() string
//}
// To create customize error, it is just a struct with Error() method implemented

func RunErrors() {
	i, err := run()
	if err == nil {
		fmt.Println(i)
	} else {
		fmt.Println(err)
	}
}

type MyError struct {
	msg string
}

// Method to implement for any error
func (err *MyError) Error() string {
	// Never fmt.Sprint(err), because it will call err.Error() and forms a endless callstack
	return fmt.Sprintln(err.msg)
}

func run() (int, error) {
	return 0, &MyError{"Error msg 124"}
}