package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	arr := getLastNumbers(root1, []int{})
	arr2 := getLastNumbers(root2, []int{})
	if len(arr) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] != arr2[i] {
			return false
		}
	}
	return true
}

func getLastNumbers(root *TreeNode, result []int) []int {
	if root == nil {
		return result
	}
	if root.Right == nil && root.Left == nil {
		return append(result, root.Val)
	}
	left := getLastNumbers(root.Left, result)
	right := getLastNumbers(root.Right, result)
	left = append(left, right...)
	return append(result, left...)
}
