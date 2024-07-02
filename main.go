package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 3; i++ {
		var num int
		fmt.Scan(&num)
		sum += num
	}

	var score = float64(sum) / 3
	fmt.Println(score)

	if score >= 60 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}
