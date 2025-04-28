package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "World", "a name to say hello to")
	help := flag.Bool("help", false, "display help")

	flag.Parse()

	if *help {
		fmt.Println("Usage: go run main.go [--name your_name] [--help]")
		os.Exit(0)
	}

	fmt.Printf("Hello, %s!\n", *name)
}
