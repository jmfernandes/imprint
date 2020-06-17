package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func printUsage() {
	fmt.Println("Usage of imprint:")
	flag.PrintDefaults()
}

func writeFile(path string) {
	f, err := os.Create(path)
	defer f.Close()
	check(err)
	_, err = f.WriteString("writes\n")
	check(err)
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	status := false
	if err == nil {
		if fileInfo.Mode()&os.ModeType == os.ModeDir {
			status = true
		}
	}
	return status
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
	// stuff, _ := filepath.Abs(args[0])
	// fmt.Println(len(args))
	// fmt.Println("filepath:", stuff)
	for i := 0; i < len(args); i++ {
		currentArg := args[i]
		_, file := filepath.Split(currentArg)
		if file == "" || isDirectory(currentArg) {
			fmt.Println("Only enter filepaths, not directory paths")
			os.Exit(2)
		}
	}

	for i := 0; i < len(args); i++ {
		file, _ := filepath.Abs(args[i])
		fmt.Println(file)
		writeFile(file)
	}
}
