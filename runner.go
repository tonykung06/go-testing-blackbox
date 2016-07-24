package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

func main() {
	val, ok := quick.Value(reflect.TypeOf(MyStruct{}), rand.New(rand.NewSource(time.Now().Unix())))

	if ok {
		fmt.Println(val.Interface().(MyStruct))
	}
}

type MyStruct struct {
	MyInt    int
	MyString string
	MySlice  []float32
}

func TestOddMultipleOfThree(t *testing.T) {
	f := func(x int) bool {
		y := funcBeingTested(x)
		return y%2 == 1 && y%3 == 0
	}
	//quick.Check will generate many random int to run the f
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
