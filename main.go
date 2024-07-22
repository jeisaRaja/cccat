package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cccat <filepath>")
		os.Exit(1)
	}

	for _, fp := range os.Args[1:] {
		if fp == "" {
			reader := bufio.NewReader(os.Stdin)
			reader.WriteTo(os.Stdout)
		} else {
			f, err := os.Open(fp)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			reader := bufio.NewReader(f)
			_, err = reader.WriteTo(os.Stdout)
		}
	}
}
