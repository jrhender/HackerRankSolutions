package main

import (
	"fmt"
	"sort"
)

func getMedian(s []int) float64 {
	l := len(s)
	if l%2 == 1 {
		i := l / 2
		return float64(s[i])
	}
	middle1 := float64(s[l/2-1])
	middle2 := float64(s[l/2])
	result := (middle1 + middle2) / 2.0 //Need to consider decimals
	return result
}

func addNewValue(trailing []int, new int) []int {
	trailing = append(trailing, new)

	// Insertion sort
	for i := len(trailing) - 2; i >= 0; i-- { // Going down through trailing to sort
		if new >= trailing[i] {
			return trailing
		}
		trailing[i+1] = trailing[i]
		trailing[i] = new
	}
	return trailing
}

func removeOldValue(a []int, toRemove int) []int {
	i := sort.SearchInts(a, toRemove)

	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = 0      // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.

	return a
}

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	//Initial setup of trailing expenditures
	var trailing []int
	for i := int32(0); i < d; i++ {
		trailing = append(trailing, int(expenditure[i]))
	}
	sort.Ints(trailing)

	//Notification calc
	numNotifications := 0
	for i := int(d); i < len(expenditure); i++ {
		median := getMedian(trailing)
		if float64(expenditure[i]) >= 2*median {
			numNotifications++
		}
		trailing = addNewValue(trailing, int(expenditure[i]))
		trailing = removeOldValue(trailing, int(expenditure[i-int(d)]))
	}
	return int32(numNotifications)
}

func main() {
	input := []int32{1, 2, 3, 4, 4}
	d := int32(4)
	fmt.Println(activityNotifications(input, d))
}
