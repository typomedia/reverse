package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"log"
	"os"
	"strings"
)

const version = "1.0"

const usage = `Usage: reverse [options] [file ...]

Reverse lines characterwise.

Options:
 -h, --help     display this help
 -V, --version  display version
`

func main() {
	flag.BoolP("help", "h", false, "")
	flag.Usage = func() {
		fmt.Print(usage)
	}
	info := flag.BoolP("version", "V", false, "")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 && !*info {
		flag.Usage()
	}

	if *info {
		fmt.Println("remove by Philipp Speck [Version " + version + "]")
		fmt.Println("Copyright (C) 2022 Typomedia Foundation.")
	}

	if len(args) > 0 {
		for _, path := range args {
			content, err := os.ReadFile(path)
			if err != nil {
				log.Fatal("Error when opening file: ", err)
			}

			fmt.Println(strings.TrimSpace(reverse(string(content))))
		}
	}

}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
