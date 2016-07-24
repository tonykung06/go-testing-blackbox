package main

import (
	"fmt"
	"math/rand"
	"reflect"
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
	MyInt int
	// MyString string
	// MySlice  []float32
}

func (ms MyStruct) Generate(rand *rand.Rand, size int) reflect.Value {
	result := rand.Float32()*20. - 10.
	return reflect.ValueOf(MyStruct{int(result)})
}

//quick.Check
// func TestOddMultipleOfThree(t *testing.T) {
// 	f := func(x int) bool {
// 		y := funcBeingTested(x)
// 		return y%2 == 1 && y%3 == 0
// 	}
// 	//quick.Check will generate many random int to run the f
// 	if err := quick.Check(f, nil); err != nil {
// 		t.Error(err)
// 	}
// }

// //quick.CheckEqual
// func square(x int) int {
// 	return x * x
// }

// func square2(x int) int {
// 	if x > 0 {
// 		return x * x
// 	} else {
// 		return 0
// 	}
// }

// func Test2() {
// 	//passing a large number of inputs and assert two functions return same values with same inputs
// 	fmt.Println(quick.CheckEqual(square, square2, nil))
// }
