package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cccat <filepath>")
		os.Exit(1)
	}
	var filepath *string
	filepath = flag.String("f", "", "specify the filepath")
	flag.Parse()
	if *filepath == "" {
		*filepath = flag.Arg(0)
	}
  // fmt.Println("Filepath: ", *filepath)

	f, err := os.Open(*filepath)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(f)
  reader.WriteTo(os.Stdout)
}
