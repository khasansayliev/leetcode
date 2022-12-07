package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	helper(root, L, R, &sum)
	return sum
}

func helper(root *TreeNode, L int, R int, sum *int) {
	if root == nil {
		return
	}
	if root.Val < L {
		helper(root.Right, L, R, sum)
		return
	}
	if root.Val > R {
		helper(root.Left, L, R, sum)
		return
	}

	*sum += root.Val
	helper(root.Left, L, R, sum)
	helper(root.Right, L, R, sum)

}
