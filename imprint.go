package main

import (
	"flag"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage of imprint:")
	flag.PrintDefaults()
}

func main() {
	boolPtr := flag.Bool("single", false, "a boolean for setting single line")
	flag.Parse()
	args := flag.Args()
	fmt.Println("args:", args)
	fmt.Println("single:", *boolPtr)
	fmt.Println("tail:", flag.Args())
	// fmt.Println(ip)
	if argLength := len(args); argLength < 1 {
		printUsage()
		os.Exit(1)
	}

}
