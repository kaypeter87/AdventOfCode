package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		nums = append(nums, perline)
	}

	var sum1 int = 0
	var sum2 int = 0
	var temp int = 0

	for i := 0; i < len(nums); i++ {
		sum1 += nums[i]/3 - 2
		temp = nums[i]
		for temp/3-2 > 0 {
			temp = temp/3 - 2
			sum2 += temp
		}
	}

	fmt.Println("Part 1 puzzle answer: ", sum1)
	fmt.Println("Part 2 puzzle answer: ", sum2)
}
