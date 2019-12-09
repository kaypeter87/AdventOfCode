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
		_, err := fmt.Fscanf(file, "%d", &perline)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}

	nums[1] = 12
	nums[2] = 2

	for i := 0; i < len(nums); i += 4 {
		if nums[i] == 1 {
			add := nums[nums[i+1]] + nums[nums[i+2]]
			nums[nums[i+3]] = add
		}
		if nums[i] == 2 {
			multiply := nums[nums[i+1]] * nums[nums[i+2]]
			nums[nums[i+3]] = multiply
		}
		if nums[i] == 99 {
			break
		}
	}

	fmt.Println("Part 1 puzzle answer: ", nums[0])
}
