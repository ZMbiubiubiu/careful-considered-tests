package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Printf("in displayName:%p\t", &d)
	fmt.Println("My Name Is", d.name)
}
func (d *data) setAge(age int) {
	d.age = age
	fmt.Printf("in setAge:%p\t", d)
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {
	// Methods are really just functions that provide syntactic sugar to provide the
	// ability for data to exhibit behavior.

	// value receiver
	d := data{
		name: "Bill",
	}
	f0 := func(d data) func() {
		return func() {
			d := d
			fmt.Printf("in displayName:%p\t", &d)
			fmt.Println("f0 My Name Is", d.name)
		}
	}(d)
	// f1 诞生的这一刻，它就无法改了
	f1 := d.displayName
	f0()
	f1()
	d.name = "Joan"
	f0()
	f1()

	d2 := &data{
		name: "Bill",
	}
	f2 := d2.displayName
	f2()
	d2.name = "Joan"
	f2()

	//	 pointer receiver
	d3 := &data{
		name: "Bill",
	}
	f3 := d3.setAge
	f3(45)
	d3.name = "Jane"
	f3(45)
}
