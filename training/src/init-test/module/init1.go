package init1

import "fmt"

func init () {
  fmt.Println("--- init1 ---")
}

func GetName () string {
  return "init1"
}

