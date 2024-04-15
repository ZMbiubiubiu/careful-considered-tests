package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("hello  world")

	var receiver = make([]byte, 5)
	buffer.Read(receiver)
	fmt.Println(receiver)
	fmt.Println(buffer.String())

	// 将结构体编码为 JSON
	buffer.Reset()
	p := Person{"Alice", 25}
	enc := json.NewEncoder(&buffer)
	enc.Encode(p)
	fmt.Println(buffer.String()) // 输出：{"Name":"Alice","Age":25}

	// 从 JSON 解码为结构体
	var p2 Person
	dec := json.NewDecoder(&buffer)
	dec.Decode(&p2)
	fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age) // 输出：Name: Alice, Age: 25
}
