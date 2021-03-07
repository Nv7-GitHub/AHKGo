package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "File to run")
}

func handle(err error) {
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(0)
	}
}

func main() {
	flag.Parse()

	script, err := ioutil.ReadFile(filename)
	handle(err)

	tm := time.Now()
	exec := createExecutable(string(script))
	fmt.Println(time.Now().Sub(tm))

	functions["Loop1"]()

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	evalExec(exec)
}
