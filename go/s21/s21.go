package s21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dh := &ListNode{}
	curr := dh

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			curr.Next = list2
			list2 = list2.Next
		} else {
			curr.Next = list1
			list1 = list1.Next
		}

		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	}

	if list2 == nil {
		curr.Next = list1
	}

	return dh.Next
}
