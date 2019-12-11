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

	copySlice := make([]int, len(nums))
	computeOutput := func(o, p int) int {
		copy(copySlice, nums)
		copySlice[1] = o
		copySlice[2] = p

		for i := 0; i < len(nums); i += 4 {
			if copySlice[i] == 1 {
				add := copySlice[copySlice[i+1]] + copySlice[copySlice[i+2]]
				copySlice[copySlice[i+3]] = add
			}
			if copySlice[i] == 2 {
				multiply := copySlice[copySlice[i+1]] * copySlice[copySlice[i+2]]
				copySlice[copySlice[i+3]] = multiply
			}
			if copySlice[i] == 99 {
				break
			}
		}
		return copySlice[0]
	}
	computeCode := func(q, w int) int {
		copy(copySlice, nums)
		copySlice[1] = 12
		copySlice[2] = 2
		for i := 0; i < len(nums); i += 4 {
			if copySlice[i] == 1 {
				add := copySlice[copySlice[i+1]] + copySlice[copySlice[i+2]]
				copySlice[copySlice[i+3]] = add
			}
			if copySlice[i] == 2 {
				multiply := copySlice[copySlice[i+1]] * copySlice[copySlice[i+2]]
				copySlice[copySlice[i+3]] = multiply
			}
			if copySlice[i] == 99 {
				break
			}
		}
		return copySlice[0]
	}

	fmt.Println(computeCode(12, 2))

	output := 19690720
	for o := 0; o <= 99; o++ {
		for p := 0; p <= 99; p++ {
			f := computeOutput(o, p)
			if f == output {
				fmt.Println(o, p, 100*o+p)
			}
		}
	}
}
