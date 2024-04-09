package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"unicode"
)

var (
	buf [1]byte
)

func readbyteV3(r io.Reader) (rune, error) {
	// 经过编译分析，该buf每次都会被分配到堆上
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	//  profile 一下 CPU
	//cpuProfile, _ := os.Create("cpu_buffer_profile")
	//pprof.StartCPUProfile(cpuProfile)
	//defer pprof.StopCPUProfile()

	memProfile, _ := os.Create("mem_profile")
	runtime.MemProfileRate = 1

	defer func() {
		pprof.WriteHeapProfile(memProfile)
		memProfile.Close()
	}()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}
	b := bufio.NewReader(f)

	words := 0
	inword := false
	for {
		r, err := readbyteV2(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%q: %d words\n", os.Args[1], words)
}
