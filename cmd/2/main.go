package main

import "fmt"

const fourMillion = 4000000

func main() {
	sum := 2
	prevs := []int{1, 2}

	for i := 1; ; i++ {
		if i < 3 {
			fmt.Println(i)
			continue
		}

		n := prevs[0] + prevs[1]
		if n > fourMillion {
			break
		}

		fmt.Println(n)

		prevs[0] = prevs[1]
		prevs[1] = n

		if n%2 == 0 {
			sum += n
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
