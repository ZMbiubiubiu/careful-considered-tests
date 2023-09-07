package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
	Age  int
	Sex  int8
}

func TestJSONUnmarshal(t *testing.T) {
	var err error
	p := Person{
		Name: "zm",
		Age:  29,
		Sex:  1,
	}

	t.Run("x", func(t *testing.T) {
		// notice name empty; Age 29->30; sex:1->2
		data := []byte(`{"Age":30,"Sex":2}`)
		err = json.Unmarshal(data, &p)
		assert.Nil(t, err)

		assert.Equal(t, "zm", p.Name)
		assert.Equal(t, 30, p.Age)
		assert.Equal(t, int8(2), p.Sex)
	})
	t.Run("valid json unmarshal to p2", func(t *testing.T) {
		data := []byte(`{"Name":"zm","Age":29,"Sex":1}`)
		var p2 Person
		err = json.Unmarshal(data, &p2)
		assert.Nil(t, err)

		assert.Equal(t, p.Name, p2.Name)
		assert.Equal(t, p.Age, p2.Age)
		assert.Equal(t, p.Sex, p2.Sex)
	})
	t.Run("empty json unmarshal to zero value struct", func(t *testing.T) {
		data := []byte(`{}`)
		var p3 Person
		err = json.Unmarshal(data, &p3)
		assert.Nil(t, err)

		assert.Equal(t, "", p3.Name)
		assert.Equal(t, 0, p3.Age)
		assert.Equal(t, int8(0), p3.Sex)
	})

	t.Run("empty json unmarshal to value struct", func(t *testing.T) {
		data := []byte(`{}`)
		err = json.Unmarshal(data, &p)
		assert.Nil(t, err)

		assert.Equal(t, "", p.Name)
		assert.Equal(t, 0, p.Age)
		assert.Equal(t, int8(0), p.Sex)
	})
}
