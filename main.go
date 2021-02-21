package main

import (
	"flag"
	"fmt"
	"github.com/J-Guenther/GoRasterize/input"
	"io/ioutil"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Only one file allowed.")
		os.Exit(1)
	}
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("File can not be read.")
		os.Exit(1)
	}
	input.ReadGeoJson(content)
}
