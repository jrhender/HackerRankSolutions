package main

import (
	"fmt"
)

func main() {
	q := [...]int32{1, 2, 5, 3, 7, 8, 6, 4}

	//Setup expectedLink
	var expectedLine []int32
	for i := int32(0); i < int32(len(q)); i++ {
		expectedLine = append(expectedLine, i+1)
	}

	totalBribes := int32(0)
	for i := int32(0); i < int32(len(q)); i++ {
		if expectedLine[i] != q[i] {
			if expectedLine[i+1] == q[i] {
				expectedLine[i+1] = expectedLine[i]
				expectedLine[i] = q[i]
				totalBribes++
			} else if expectedLine[i+2] == q[i] {
				expectedLine[i+2] = expectedLine[i+1]
				expectedLine[i+1] = expectedLine[i]
				expectedLine[i] = q[i]
				totalBribes += 2
			} else {
				fmt.Println(i)
				fmt.Println("Too chaotic")
				return
			}
		}
	}
	fmt.Println(totalBribes)
}
