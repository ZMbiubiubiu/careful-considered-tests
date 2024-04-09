package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
	"unicode"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	// 缓冲区太小，导致太多的read系统调用
	// read这种blocked的系统调用多，导致G与M频繁分离
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func main() {
	//  profile 一下 CPU
	cpuProfile, _ := os.Create("cpu_profile")
	pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("could not open file %q: %v", os.Args[1], err)
	}

	words := 0
	inword := false
	for {
		r, err := readbyte(f)
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
