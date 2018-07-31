/*
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (42.42%)
 * Total Accepted:    380.2K
 * Total Submissions: 896.2K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	res := &ListNode{}
	current := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			current.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else if l2 == nil {
			current.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				current.Next = &ListNode{Val: l1.Val}
				l1 = l1.Next
			} else {
				current.Next = &ListNode{Val: l2.Val}
				l2 = l2.Next
			}
		}
		current = current.Next
	}
	return res.Next
}