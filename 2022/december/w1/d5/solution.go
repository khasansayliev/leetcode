package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	i := 0
	for head != nil {
		m[i] = head
		i++
		head = head.Next
	}
	length := len(m)

	if length == 0 {
		return head
	}
	part := length / 2
	if part*2 == length {
		return m[part]
	}
	return m[part]
}
