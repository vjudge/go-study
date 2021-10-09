package main_test

import (
	"fmt"
	"testing"
)

func TestRune (t *testing.T) {
	str := "放假就开始学 Go"
	fmt.Printf("str: %q\n", str)
	fmt.Printf("str-char: %q\n", []rune(str)) // char
	fmt.Printf("str-hex: %x\n", []rune(str)) // hex
	fmt.Printf("str-hex: [% x]\n", []byte(str)) // hex
}
