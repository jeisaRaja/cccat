package program

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Program struct {
	filepaths []string
	LineBool  bool
	line      int
}

func (p *Program) ReadFiles() {
	for _, fp := range p.filepaths {
		if fp == "-" {
			reader := bufio.NewReader(os.Stdin)
			p.write(reader)
			return
		}
		file, err := os.Open(fp)
		if err != nil {
			fmt.Println("error while opening file, ", err)
			break
		}
		reader := bufio.NewReader(file)
		p.write(reader)
	}
}

func (p *Program) AddFilepaths(path string) {
	p.filepaths = append(p.filepaths, path)
	fmt.Println(p.filepaths)
}

func (p *Program) write(reader *bufio.Reader) {
	for {
		p.line++
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("error while reading line, ", err)
		}
		line = strings.TrimRight(line, "\n")
		if p.LineBool {
			fmt.Println(p.line, line)
		} else {
			fmt.Println(line)
		}
	}
}
