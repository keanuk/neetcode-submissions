/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	curr := head
	curr1 := list1
	curr2 := list2
	for curr1 != nil || curr2 != nil {
		if curr2 == nil {
			curr.Next = curr1
			return head.Next
		} else if curr1 == nil {
			curr.Next = curr2
			return head.Next
		}
		if curr1.Val <= curr2.Val {
			curr.Next = curr1
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr2 = curr2.Next
		}
		curr = curr.Next
	}
	return head.Next
}
