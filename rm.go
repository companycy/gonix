package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func main() {
	recursivePtr := flag.Bool("r", false, "recursively delete a directory")

	flag.Parse()

	src, err := os.Stat(flag.Arg(0))
	if err != nil {		// if existing
		log.Fatal(err)
	}

	if src.IsDir() {
		if *recursivePtr {
			err = os.RemoveAll(flag.Arg(0))
		} else {
			fmt.Printf("rm: %s: is a directory\n", flag.Arg(0))
		}
	} else {
		err = os.Remove(flag.Arg(0))
	}

	if err != nil {
		fmt.Println(err)
	}
}
