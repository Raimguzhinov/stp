package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
)

func productOfNonZero(v []int, ind []int) (int, error) {
	if len(v) == 0 || len(ind) == 0 {
		return 0, fmt.Errorf("empty array")
	}
	product := 1
	hasNonZero := false
	for _, i := range ind {
		if i < 0 || i >= len(v) {
			return 0, fmt.Errorf("index %d is out of bounds for array v", i)
		}
		if v[i] != 0 {
			product *= v[i]
			hasNonZero = true
		}
	}
	if !hasNonZero {
		return 0, nil
	}
	return product, nil
}

func findMinAndIndex(arr []int) (int, int, error) {
	if len(arr) == 0 {
		return 0, -1, errors.New("array is empty")
	}

	minVal := arr[0]
	minIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minVal = arr[i]
			minIndex = i
		}
	}

	return minVal, minIndex, nil
}

func reverseArray(arr []float64) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
}

type Options struct {
	Values []int `short:"v" long:"values" description:"Array of values" required:"true"`
	Ind    []int `short:"i" long:"indexes" description:"Array of indexes" required:"true"`
}

func main() {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("Failed to parse flags: %v", err)
	}

	result, err := productOfNonZero(opts.Values, opts.Ind)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Product of non-zero elements is:", result)
	}
}
