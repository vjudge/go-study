package main_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestMyTest (t *testing.T)  {
	err := errors.New("This is error.")
	t.Log("err:", err)
	if err != nil {
		t.Fail()
		t.Error("This is Error.")
		t.Errorf("This is %s Errorf.", "custom")
	}
	//if err != nil {
	//	t.FailNow()
	//	t.Fatal("This is Fatal.")
	//	t.Fatalf("This is %s Fatal.", "custom")
	//}
	if err != nil {
		t.SkipNow()
		t.Skip("This is Skip.")
	}
	t.Log("This is test Log.")
	t.Logf("This is %s's Logf", "vjude")
}

func BenchmarkMyTest (t *testing.B)  {
	fmt.Println("this is my test Benchmark.")
}

func ExampleMyTest (t *testing.T)  {
	fmt.Println("this is my test Example.")
}


