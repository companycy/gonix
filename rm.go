package main

import (
  "flag"
  "os"
  "fmt"
)

func main() {

  flag.Parse()
  err := os.Remove(flag.Arg(0))

  if err != nil {
      fmt.Println(err)
  }
}
