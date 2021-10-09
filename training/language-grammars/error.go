package main

import (
	"errors"
	"fmt"
	"os"
)

func print(req string) (res string, err error) {
	if req == "" {
		err = errors.New("req is empty")
		return
	}
	fmt.Println("req: ", req)
	return
}

func main ()  {
	_, err1 := print("vjudge coming")
	fmt.Println("err:", err1)
	_, err2 := print("")
	fmt.Println("err:", err2)
	os.Exit(0)
}