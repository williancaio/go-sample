package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(sum(1, 2))

	f := func(text string) {
		fmt.Println(text)
	}

	f("test")

	sum, sub := calc(1, 1)

	fmt.Println(sum, sub)

	sum1, _ := calc(1, 1)

	fmt.Println(sum1)
	fmt.Println("==========")
	fmt.Println(namedCalc(1, 1))

	for i := 0; i < 20; i++ {
		fmt.Println(fibonacci(i))
	}

}

func calc(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func namedCalc(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func manyParams(num ...int) {

}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}
