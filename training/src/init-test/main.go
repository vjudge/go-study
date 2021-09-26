package main

import (
  "fmt"
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