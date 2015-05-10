package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
	flag.Parse()
	src, err := os.Stat(flag.Arg(0))
	if err != nil {		// if existing
		log.Fatal(err)
	}

	if src.IsDir() {
		err = os.RemoveAll(flag.Arg(0)) // default with -r
	} else {
		err = os.Remove(flag.Arg(0))
	}

	if err != nil {
		fmt.Println(err)
	}
}
