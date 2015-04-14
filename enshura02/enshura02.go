package enshura02

import (
	"io"
	"fmt"
)

type TotalPerDay [7]int

func Count(integers []int) (result TotalPerDay) {
	for idx, number := range integers {
		result[idx%7] += number
	}
	return
}

func Parse(r io.Reader) ([]int, error) {
	var n int

	fmt.Fscanf(r, "%d\n", &n)
	numbers := make([]int, 0, n)

	var number int
	for i := 0; i < n; i++ {
		if pn, err := fmt.Fscanf(r, "%d\n", &number); err != nil || pn != 1 {
			return numbers, err
		}
		numbers = append(numbers, number)
	}

	return numbers, nil
}

func Main(r io.Reader) error {
	numbers, err := Parse(r)
	if err != nil {
		fmt.Errorf("error occured in parsing")
		return err
	}

	totals_pd := Count(numbers)

	for _, total := range totals_pd {
		fmt.Printf("%d\n", total)
	}

	return nil
}
