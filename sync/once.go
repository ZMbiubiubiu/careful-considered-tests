package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func NewPerson() *Person {
	one.Do(func() {
		fmt.Println("create person")
		person = &Person{
			Name: "bingo",
			Age:  31,
		}
	})
	return person
}

var (
	person *Person
	one    sync.Once
)

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			NewPerson()
			wg.Done()
		}()
	}
	wg.Wait()
}
