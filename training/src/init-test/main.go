package main

import (
  "fmt"
  "init-test/module/init2"
  "module/init1"
)


func init () {
  fmt.Println("init main")
}

func main () {
  fmt.Println("--- main ---")
}

func init () {
  fmt.Println("end main")
}