package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// no file: reprint waht you write. IO
// when there is a file print the file

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		echoLogging()
		os.Exit(0)
	}

	printFile(args[0])
}

func echoLogging() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}

func printFile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("There is no file named %s", file)
	}

	fmt.Print(string(data))
}
