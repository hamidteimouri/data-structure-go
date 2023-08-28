package main

func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, ok := numToIndex[complement]; ok {
			return []int{index, i}
		}
		numToIndex[num] = i
	}

	return nil // No solution found
}

/*
func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)

    if result != nil {
        fmt.Printf("Indices: %v\n", result)
        fmt.Printf("Numbers: %d, %d\n", nums[result[0]], nums[result[1]])
    } else {
        fmt.Println("No solution found.")
    }
}
*/
