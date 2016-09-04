package main

import (
	"github.com/pierrec/xxHash/xxHash32"
	"bytes"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	buf := bytes.NewBufferString("this is a test")
	fmt.Printf("\"%x\"", xxHash32.Checksum(buf.Bytes(), 0xCAFE)) // hash the bytes with a seed
	fmt.Printf("  Elapsed: %v\n", time.Now().Sub(t))
}
