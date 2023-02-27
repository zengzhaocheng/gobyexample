// Go supports [_anonymous functions_](https://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.
// 一个函数内引用了外部的局部变量，这种现象，就称之为闭包。
package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func incr() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
func func1() (i int) {
	i = 10
	defer func() {
		i += 1
	}()
	return 5
}

func main() {
	//闭包对外层词法域变量是引用的
	//a closure get a variable's reference
	fmt.Println("======1======")
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
	fmt.Println(a("All"))
	println("=======2======")

	i := incr()
	println(i())
	println(i())
	println(i())
	println(incr()())
	println(incr()())
	println(incr()())
	fmt.Println("=====3======")

	x := 1
	f := func() {
		println(x)
	}
	x = 2
	x = 3
	f() // 3

	fmt.Println("=====4======")
	x1 := 1
	func() {
		println(x1) // 1
	}() //这里已经解引用了, get the value here
	x1 = 2
	x1 = 3

	//等价于 equals to
	x2 := 1
	f2 := func() {
		println(x2)
	}
	f2() // 1
	x2 = 2
	x2 = 3

	x3 := 1
	func() {
		println(&x3) //
	}()
	println(&x3) // will be same as x3 printed in func()

	fmt.Println("=====5======")
	//循环里闭包
	for i := 0; i < 3; i++ {
		func() {
			println(i) // 0, 1, 2
		}()
	}
	fmt.Println("===========")
	var dummy [3]int
	for i := 0; i < len(dummy); i++ {
		println(i) // 0, 1, 2
	}
	fmt.Println("===========")
	var f4 func()
	for i := 0; i < len(dummy); i++ {
		f4 = func() {
			println(i)
		}
	}
	f4() // 3 i 逃逸了？ i escaped?
	fmt.Println("===========")
	var f5 func()
	for i := range dummy {
		f5 = func() {
			println(i)
		}
	}
	f5() // 2

	fmt.Println("=====6======")

	var funcSlice []func()
	for i := 0; i < 3; i++ {
		funcSlice = append(funcSlice, func() {
			println(i)
		})
	}
	for j := 0; j < 3; j++ {
		funcSlice[j]() // 3, 3, 3
	}
	fmt.Println("===========")
	var funcSlice1 []func()
	for i := 0; i < 3; i++ {
		println(&i) // 0xc0000ac1d0 0xc0000ac1d0 0xc0000ac1d0
		funcSlice1 = append(funcSlice1, func() {
			println(&i)
		})

	}
	for j := 0; j < 3; j++ {
		funcSlice1[j]() // 0xc0000ac1d0 0xc0000ac1d0 0xc0000ac1d0
	}
	fmt.Println("===========")

	var funcSlice2 []func()

	for i := 0; i < 3; i++ {
		j := i
		funcSlice2 = append(funcSlice2, func() {
			println(j)
		})
	}
	for j := 0; j < 3; j++ {
		funcSlice2[j]() // 0, 1, 2
	}
	fmt.Println("===========")
	var funcSlice3 []func()
	for i := 0; i < 3; i++ {
		func(i int) {
			funcSlice3 = append(funcSlice3, func() {
				println(i)
			})
		}(i)

	}
	for j := 0; j < 3; j++ {
		funcSlice3[j]() // 0, 1, 2
	}
	fmt.Println("=====7======")
	f6 := func() (i int) {
		i = 10
		defer func() {
			i += 1
		}()
		return 5
	}
	closure := f6()
	fmt.Println(closure) //6
	fmt.Println("=====8======")

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())

}
