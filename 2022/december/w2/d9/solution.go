package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	m1 := calc(root.Left, root.Val, root.Val)
	m2 := calc(root.Right, root.Val, root.Val)
	if m1 > m2 {
		return m1
	}
	return m2
}

func calc(root *TreeNode, max, min int) int {
	if root == nil {
		return max - min
	}
	if root.Val > max {
		max = root.Val
	}
	if root.Val < min {
		min = root.Val
	}
	m1 := calc(root.Left, max, min)
	m2 := calc(root.Right, max, min)
	if m1 > m2 {
		return m1
	}
	return m2
}
