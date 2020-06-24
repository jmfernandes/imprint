package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const header = "bbbxxxxxxxxxxxxxxxxxxxxx\nxx\nxx file: fff\nxx\nxx author: aaa\nxx\nxx created: ttt\nxx\nxxxxxxxxxxxxxxxxxxxxxeee\n"

var language = map[string][]string{
	".py":  {"#", "#", "#"},
	".cpp": {"/*", "*", "*/"},
}

func main() {
	boolPtr := flag.Bool("single", false, "a boolean for setting single line")
	flag.Parse()
	args := flag.Args()

	if *boolPtr {
		fmt.Println("boolen present")
	}

	if argLength := len(args); argLength < 1 {
		printUsage()
		os.Exit(1)
	}

	for i := 0; i < len(args); i++ {
		currentArg := args[i]
		_, file := filepath.Split(currentArg)
		if file == "" || isDirectory(currentArg) {
			fmt.Println("Only enter filepaths, not directory paths")
			os.Exit(2)
		}
	}

	for i := 0; i < len(args); i++ {
		currentArg := args[i]
		file, _ := filepath.Abs(currentArg)
		writeFile(file, createString(file))
	}
}

func createString(file string) string {
	currentTime := time.Now()
	_, fileName := filepath.Split(file)
	fileExtention := filepath.Ext(file)
	author := "josh"
	values, found := language[fileExtention]
	if !found {
		values = []string{"#", "#", "#"}
	}
	r := strings.NewReplacer(
		"bbb", values[0],
		"x", values[1],
		"eee", values[2],
		"fff", fileName,
		"aaa", author,
		"ttt", currentTime.Format("January 2, 2006"))
	return (r.Replace(header))
}

func writeFile(path string, content string) {
	f, err := os.Create(path)
	defer f.Close()
	check(err)
	_, err = f.WriteString(content)
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

func printUsage() {
	fmt.Println("Usage of imprint:")
	flag.PrintDefaults()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
