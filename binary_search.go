package main

import {
  "fmt"
}

func BinarySearchInt(searchInt int, sortedArr []int) (int, error) {
	// get array length
	arrLen := len(sortedArr)

	// get first element and last element of array
	first_elem := 0
	last_elem := arrLen - 1

	for first_elem <= last_elem {
		midpoint := (first_elem + last_elem) / 2
		guess := sortedArr[midpoint]

		// return element location if item is found
		if guess == searchInt {
			return midpoint, nil
		}

		if guess < searchInt {
			last_elem = midpoint - 1
		} else {
			first_elem = midpoint + 1
		}
	}

	return -1, fmt.Errorf("Element %v not found", searchInt) 
}
