// Dabbling in lambda calculus, based initially on David Beazley's PYCON 2019 tutorial
// "Lambda Calculus from the Ground Up" (YouTube video).
// This was then followed up by a study of 4 YouTube videos :
//      CSE 340 S16: 4-15-16 "Lambda Calculus Pt.1" and parts 2, 3 and 4.
//      Adam Doupe - Arizona State University
// These clarified a lot for me.
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

// The incr function is just a helper, not part of the LC.
func incr(x int) int {
	return x + 1
}

// The commented out functions were based on the Python Beazley code.

//func lambdaIncr(x any) any {
//	xa := x.(int)
//	return xa + 1
//}

//func lambdaOne(f func(int) int) func(int) int {
//	return func(x int) int {
//		return f(x)
//	}
//}

//func lambdaTwo(f func(int) int) func(int) int {
//	return func(x int) int {
//		return f(f(x))
//	}
//}

//func lambdaTwoAny(f func(any) any) func(any) any {
//	return func(x any) any {
//		return f(f(x))
//	}
//}

//func lambdaThree(f func(int) int) func(int) int {
//	return func(x int) int {
//		return f(f(f(x)))
//	}
//}

//func lambdaThreeAny(f func(any) any) func(any) any {
//	return func(x any) any {
//		return f(f(f(x)))
//	}
//}

//func lambdaZero(f func(int) int) func(int) int {
//	return func(x int) int {
//		return x
//	}
//}
// These functions were written after I took Adam Doupe's course.

//Zero ƛf.ƛx.x
//One  ƛf.ƛx.fx
//Two  ƛf.ƛx.ffx

func lambdaOne(f func(int) int, x int) int {
	return f(x)
}

func lambdaTwo(f func(int) int, x int) int {
	return f(f(x))
}

func lambdaThree(f func(int) int, x int) int {
	return f(f(f(x)))
}

func lambdaZero(f func(int) int, x int) int {
	return x
}

// Example 5

//Successor ƛn.ƛf.ƛx.f(nfx)
func lambdaSucc(cn func(func(int) int, int) int,
	f func(int) int,
	x int) func(func(int) int, int) int {
	return func(func(int) int, int) int {
		return f(cn(f, x))
	}
}

// T ƛx.ƛy.x
// F ƛx.ƛy.y
// AND ƛa.ƛb.abF
// NOT ƛa.aFT

func lambdaT(x, y any) any {
	return x
}

func lambdaF(x, y any) any {
	return y
}

func lambdaNOT(x func(a, b any) any) any {
	return x(lambdaF, lambdaT)
}

func lambdaAND(x, y func(a, b any) any) any {
	return x(y, lambdaF)
}

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
	fmt.Println("One", lambdaOne(incr, 0))
	fmt.Println("Two", lambdaTwo(incr, 0))
	fmt.Println("Three", lambdaThree(incr, 0))
	fmt.Println("Zero", lambdaZero(incr, 0))

	// Example 5
	newCN := lambdaSucc(lambdaOne, incr, 0)
	fmt.Println("successor of 1", newCN(incr, 0))
	newCN = lambdaSucc(lambdaThree, incr, 0)
	fmt.Println("successor of 3", newCN(incr, 0))
	fmt.Println("T", lambdaT(true, false))
	fmt.Println("F", lambdaF(true, false))
	notRes := lambdaNOT(lambdaT)
	notResA := notRes.(func(a, b any) any)
	fmt.Println("Not T", notResA(true, false))
	notRes = lambdaNOT(lambdaF)
	notResA = notRes.(func(a, b any) any)
	fmt.Println("Not F", notResA(true, false))
	andRes := lambdaAND(lambdaT, lambdaF)
	andResA := andRes.(func(a, b any) any)
	fmt.Println("And T F", andResA(true, false))
	andRes = lambdaAND(lambdaT, lambdaT)
	andResA = andRes.(func(a, b any) any)
	fmt.Println("And T T", andResA(true, false))
	andRes = lambdaAND(lambdaF, lambdaT)
	andResA = andRes.(func(a, b any) any)
	fmt.Println("And F T", andResA(true, false))
	andRes = lambdaAND(lambdaF, lambdaF)
	andResA = andRes.(func(a, b any) any)
	fmt.Println("And F F", andResA(true, false))
}
