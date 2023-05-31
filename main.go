package main

import (
	"bufio"
	_ "fmt"
	"io"
	_ "log"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("data/in.txt")
	defer file.Close()
	readFile := bufio.NewReader(file)
	newFile, _ := os.Create("data/data_out.txt")
	defer newFile.Close()
	for {
		line, err := readFile.ReadString('\n')
		modLine := "Name: "
		var k int
		var typedString string
		for _, symbol := range line {
			typedString += string(symbol)
			if symbol == '|' {
				if typedString == "|" {
					panic("parse error: empty field on string")
				}
				k += 1
				switch {
				case k == 1:
					modLine += typedString + " Address: "
					typedString = ""

				case k == 2:
					modLine += typedString + " City: "
					typedString = ""
					k += 1
				}
			}
		}
		if typedString == "" {
			panic("parse error: empty field on string")
		}
		modLine += typedString
		strings.Split(modLine, "|")
		for _, words := range modLine {
			newFile.WriteString(string(words))
		}
		if err == io.EOF {
			break
		}

	}

}
