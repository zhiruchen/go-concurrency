package examples

import "fmt"

func FanInFanout() {
	nums := []int{2, 3, 4, 5, 8, 6}

	numberChan := numberChan(nums)
	c1 := sq(numberChan)
	c2 := sq(numberChan)

	c := fanIn(c1, c2)

	var sum int
	for i := 1; i <= len(nums); i++ {
		sum += <-c
	}

	fmt.Println(sum)
}

func numberChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

func sq(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}

		close(out)
	}()

	return out
}

func fanIn(c1, c2 <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case n := <-c1:
				c <- n
			case n := <-c2:
				c <- n
			}
		}
	}()

	return c
}
