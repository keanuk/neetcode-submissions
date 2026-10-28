/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var curr *ListNode
	next := head
	for next != nil {
		tmp := next.Next
		next.Next = curr
		curr = next
		next = tmp
	}
	return curr
}
