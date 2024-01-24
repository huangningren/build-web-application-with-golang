package main

import (
	"fmt"
	"net"
)

func main() {
	//err := Decode(1)
	//fmt.Println(err == nil)
	//
	//err2 := Decode2(1)
	//fmt.Println(err2 == nil)

	err3 := UseNetError(-1)
	fmt.Println(err3 == nil)
	if nerr, ok := err3.(net.Error); ok && nerr.Temporary() {
		fmt.Println("Temporary is true")
	}

}

type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string {
	return e.msg
}

func Decode(i int16) *SyntaxError {
	var err *SyntaxError
	if i < 0 {
		fmt.Println("小于0")
		err = &SyntaxError{
			msg:    "小于0",
			Offset: 10,
		}
	}

	return err
}

func Decode2(i int16) error {
	var err *SyntaxError
	if i < 0 {
		fmt.Println("小于0")
		err = &SyntaxError{
			msg:    "小于0",
			Offset: 10,
		}
	}

	// return nil
	return err
}

func UseNetError(i int16) error {
	if i < 0 {
		fmt.Println("小于0")
		err := &net.DNSError{
			Err:         "小于于0",
			IsTimeout:   true,
			IsTemporary: true,
			IsNotFound:  true,
		}

		return err
	}

	return nil
}
