package lib

import (
	"fmt"
	"sort"
)

/*
The problem:

Give an array of ribbon sizes and an expected ribbons count. The function will cut the ribbons form the array (but will not join them) and will return the size of each of the new ribbons. While if the given array of ribbon sizes can't produce as much ribbon as the user wants then the function will return an empty array with a message saying can't product that much ribbon.
*/
func ProcessRibbons(arr []int, desiredCount int) (int, string) {
	var (
		size      int
		msg       string
		currCount int
	)
	// Sort the sizes
	sort.Ints(arr)
	for _, v := range arr {
		currCount = 0
		for _, r := range arr {
			dividedRibbons := r / v
			currCount += dividedRibbons
		}
		if currCount == desiredCount {
			size = v
			msg = fmt.Sprintf("Minimum ribbon size for %v ribbons from array %v is %v\n", desiredCount, arr, size)
			break
		} else {
			msg = fmt.Sprintf("Can't convert %v array into %v ribbons\n", arr, desiredCount)
		}
	}

	fmt.Print(msg)
	return size, msg
}
