package main

import (
	"fmt"
	"testing"
)

// MyError a simple implement of error interface
type MyError struct {
	desc string
}

func (e *MyError) Error() string {
	return e.desc
}

func check() error {
	var err *MyError

	if 1 > 2 { // will not execute
		err = &MyError{"hello"}
	}

	if err != nil {
		fmt.Println("inner check, err is not nil")
		return err
	}

	fmt.Println("inner check, err is nil")
	return err // return err or nil is matter
}

func TestError(t *testing.T) {
	err := check()
	if err != nil {
		t.Fatalf("it should be nil, but got:%v", err)
	}
	//assert.Nil(t, err)
}
