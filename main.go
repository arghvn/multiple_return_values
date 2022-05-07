package main

import "fmt"

//The Go programming language allows us to have multiple outputs for each function.
// For example, we can take the output of the function and the error from the function at the same time,
// and like some languages, we do not need to have one output and we evaluate whether the output is error or not. Not !

// The '(int, int)' in this function signature shows that  the function returns 2 'int's.

// Here we use the 2 different return values from the call with _multiple assignment_.

// If you only want a subset of the returned values, use the blank identifier '_'.
func vals() (int, int) {

	return 3, 7

}

// Go has built-in support for _multiple return values_This feature is used often in idiomatic Go,
// for example to return both result and error values from a function.

func main() {

	a, b := vals()

	fmt.Println(a)

	fmt.Println(b)

	_, c := vals()

	fmt.Println(c)

}

// output :
// 3
// 7
// 7

// more details :
//In the above code, we created a function called vals, but this function has two outputs,
// so how to define the output data type is different than a normal function with
//The syntax (..., type-data, type-data) is defined, and its output must be separated by.

//In the first example, we obtained two outputs of the vals function using two variables a and b.
//In the next example, we extracted only the second output with variable c,
//and used the blank identifier _ instead of the first output so that the value of the first output would not be received or extracted.
