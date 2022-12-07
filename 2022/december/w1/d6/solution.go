package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	i := 1
	var odd, even, result *ListNode
	for head != nil {
		if i%2 == 0 {
			even = &ListNode{
				Val:  head.Val,
				Next: even,
			}
		} else {
			odd = &ListNode{
				Val:  head.Val,
				Next: odd,
			}
		}
		head = head.Next
		i++
	}

	for even != nil {
		result = &ListNode{
			Val:  even.Val,
			Next: result,
		}
		even = even.Next
	}

	for odd != nil {
		result = &ListNode{
			Val:  odd.Val,
			Next: result,
		}
		odd = odd.Next
	}

	return result
}
