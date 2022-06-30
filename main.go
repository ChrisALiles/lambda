// Dabbling in lambda calculus, based mainly on David Beazley's PYCON 2019 tutorial
// "Lambda Calculus from the Ground Up" (YouTube video).
package main

import (
	"fmt"
)

// Comment - all that exists are single-argument functions.

// Example 1 - modelling a switch, current goes either left or right.

func left(a any) func(b any) any {
	return func(b any) any {
		return a
	}
}

func right(a any) func(b any) any {
	return func(b any) any {
		return b
	}
}

// Example 2 - the same code can be used for true and false.

func lambdaTrue(a any) func(b any) any {
	return func(b any) any {
		return a
	}
}

func lambdaFalse(a any) func(b any) any {
	return func(b any) any {
		return b
	}
}

// Comment - we are modelling, not with data, but with behaviour.

// Example 3 - Boolean functions.

func lambdaNot(x func(any) func(any) any) any {
	return x(lambdaFalse)(lambdaTrue)
}

func lambdaAnd(x func(any) func(any) any) any {
	return func(y any) any {
		return x(y)(x)
	}
}

func lambdaOr(x func(any) func(any) any) any {
	return func(y any) any {
		return x(x)(y)
	}
}

// Example 4 - the integers.

// The incr and lambdaIncr functions are just helpers, not part of the LC.
func incr(x int) int {
	return x + 1
}

func lambdaIncr(x any) any {
	xa := x.(int)
	return xa + 1
}

func lambdaOne(f func(int) int) func(int) int {
	return func(x int) int {
		return f(x)
	}
}

func lambdaTwo(f func(int) int) func(int) int {
	return func(x int) int {
		return f(f(x))
	}
}

func lambdaTwoAny(f func(any) any) func(any) any {
	return func(x any) any {
		return f(f(x))
	}
}

func lambdaThree(f func(int) int) func(int) int {
	return func(x int) int {
		return f(f(f(x)))
	}
}

func lambdaThreeAny(f func(any) any) func(any) any {
	return func(x any) any {
		return f(f(f(x)))
	}
}

func lambdaZero(f func(int) int) func(int) int {
	return func(x int) int {
		return x
	}
}

// Example 5 - arithmetic

// Python: SUCC = lambda n:(lambda f:lambda x: f(n(f)(x)))
// Successor function.
// It took me 2 weeks to come up with this, and it's a long way
// from being good.
func lambdaSucr(twoF func(z func(any) any) func(any) any) any {
	return lambdaIncr(twoF(lambdaIncr)(0))
}

//
// Python
// Add     lambda x:lambda y:y(SUCC)(x)
// Mutiply lambda x:lambda y:lambda f:y(x(f))
func main() {
	// Example 1
	fmt.Println("Left", left("5V")("GRND"))
	fmt.Println("Right", right("loud")("soft"))

	// Example 2
	fmt.Println("True", lambdaTrue("true")("false"))
	fmt.Println("False", lambdaFalse("true")("false"))

	// Example 3
	x := lambdaNot(lambdaTrue)                   // Returns dynamic type = function, static type = interface.
	xc := x.(func(a any) func(b any) any)        // Type assertion to convert interface to function.
	y := lambdaNot(lambdaFalse)                  // Returns dynamic type = function, static type = interface.
	yc := y.(func(a any) func(b any) any)        // Type assertion to convert interface to function.
	fmt.Println("Not true", xc("arg1")("arg2"))  // Call the functions
	fmt.Println("Not false", yc("arg1")("arg2")) //   to observe their effects.

	x = lambdaAnd(lambdaTrue)               // For And and Or
	xnc := x.(func(y any) any)              //   we have to call and type
	xnnc := xnc(lambdaTrue)                 //   assert the (results of the) functions
	xc = xnnc.(func(a any) func(b any) any) //   one at a time.
	fmt.Println("And true true", xc("arg1")("arg2"))
	x = lambdaAnd(lambdaTrue)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaFalse)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("And true false", xc("arg1")("arg2"))
	x = lambdaAnd(lambdaFalse)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaTrue)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("And false true", xc("arg1")("arg2"))
	x = lambdaAnd(lambdaFalse)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaFalse)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("And false false", xc("arg1")("arg2"))

	x = lambdaOr(lambdaTrue)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaTrue)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("Or true true", xc("arg1")("arg2"))
	x = lambdaOr(lambdaTrue)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaFalse)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("Or true false", xc("arg1")("arg2"))
	x = lambdaOr(lambdaFalse)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaTrue)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("Or false true", xc("arg1")("arg2"))
	x = lambdaOr(lambdaFalse)
	xnc = x.(func(y any) any)
	xnnc = xnc(lambdaFalse)
	xc = xnnc.(func(a any) func(b any) any)
	fmt.Println("Or false false", xc("arg1")("arg2"))

	// Example 4
	fmt.Println("One", lambdaOne(incr)(0))
	fmt.Println("Two", lambdaTwo(incr)(0))
	fmt.Println("Two any", lambdaTwoAny(lambdaIncr)(0))
	fmt.Println("Two any variation", lambdaIncr(lambdaTwoAny(lambdaIncr)(0)))
	fmt.Println("Three", lambdaThree(incr)(0))
	fmt.Println("Zero", lambdaZero(incr)(0))

	// Example 5
	fmt.Println("lambdaSucr 2", lambdaSucr(lambdaTwoAny))
	fmt.Println("lambdaSucr 3", lambdaSucr(lambdaThreeAny))
	//fmt.Println("sucr sucr", lambdaSucr(lambdaSucr(lambdaThreeAny)))
}
