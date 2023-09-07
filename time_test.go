package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnixZeroTime(t *testing.T) {
	// when timestamp is 0, it means at 1970-01-01 08:00:00
	ts := time.Unix(0, 0)
	fmt.Printf("format:%s ts:%d \n", ts.Format("2006-01-02 15:04:05"), ts.Unix())
	assert.Equal(t, "1970-01-01 08:00:00", ts.Format("2006-01-02 15:04:05"))
}
