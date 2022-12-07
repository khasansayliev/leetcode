// 2256. Minimum Average Difference
package main

func main() {
	minimumAverageDifference([]int{2, 5, 3, 9, 5, 3})
}

func minimumAverageDifference(nums []int) int {
	min, minIndex := 1000000, -1
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum, rightSum := 0, totalSum
	for i, num := range nums {
		leftSum += num
		rightSum -= num
		leftAvg := leftSum / (i + 1)
		rightAvg := 0
		if len(nums)-i-1 != 0 {
			rightAvg = rightSum / (len(nums) - i - 1)
		}
		dif := leftAvg - rightAvg
		if dif < 0 {
			dif = dif * -1
		}
		if dif < min {
			min = dif
			minIndex = i
		}
	}
	return minIndex
}
