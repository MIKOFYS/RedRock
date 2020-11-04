package main

import (
	"bytes"
	"fmt"
	"os"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func main(){
	var writer bytes.Buffer
	user := "don't communicate by sharing memory share memory by communicating"
	p, _ := os.Create("proverb.txt")
	for _,u := range user{
		n, err := writer.Write([]byte(u))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(u) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	}



