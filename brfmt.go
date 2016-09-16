package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func printNewLine(scope int) {
	fmt.Printf("\n")
	for i := 0; i < scope; i++ {
		fmt.Printf("  ")
	}
}

func main() {
	fileName := os.Args[1]

	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Couldn't read file:", err)
	}

	scope := 0
	for i := 0; i < len(dat); i++ {
		if dat[i] == '{' || dat[i] == '[' {
			fmt.Printf("%c", dat[i])
			scope++
			printNewLine(scope)
			continue
		}
		if dat[i] == '}' || dat[i] == ']' {
			scope--
			printNewLine(scope)
			fmt.Printf("%c", dat[i])
			continue
		}
		if dat[i] == ',' {
			fmt.Printf(",")
			printNewLine(scope)
			continue
		}
		if dat[i] == ':' {
			fmt.Printf(": ")
			continue
		}
		fmt.Printf("%c", dat[i])
	}
	fmt.Println()
}
