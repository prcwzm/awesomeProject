package algorithm

import "fmt"

func Fib(times int) {

	f0 := 0
	f1 := 1
	fib := f1
	if times == 1 {
		fmt.Println(f0)
	}
	if times == 2 {
		fmt.Println(f1)
	}

	times = times - 1

	for i := 0; i < times; i++ {
		fib = f0 + f1
		f0 = f1
		f1 = fib
	}
	println(fib)
}
