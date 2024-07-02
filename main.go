package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 3; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}

	fmt.Println(float64(sum) / 3)
	fmt.Println("Congratulations, you are accepted!")
}
