/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (64.34%)
 * Likes:    643
 * Dislikes: 0
 * Total Accepted:    155.5K
 * Total Submissions: 232.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 方法一 —— recursion
// func swapPairs(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	n := head.Next
// 	head.Next = swapPairs(head.Next.Next)
// 	n.Next = head
// 	return n
// }

// 方法二 —— iteration
func swapPairs(head *ListNode) *ListNode {
	prev := &ListNode{Next: head}
	head = prev
	for prev.Next != nil && prev.Next.Next != nil {
		cur := prev.Next
		next := prev.Next.Next
		prev.Next, cur.Next, next.Next = next, next.Next, cur
		prev = cur
	}
	return head.Next
}

// @lc code=end

