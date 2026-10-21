/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-31 16:58:30
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil {
		return []int{-1, -1}
	}
	pre, cur := head, head.Next
	firstIndex, preIndex, index := 0, 0, 1
	minDistance := 1000000
	for cur.Next != nil {
		if pre.Val < cur.Val && cur.Val > cur.Next.Val {
			if firstIndex == 0 {
				firstIndex = index
			}
			if preIndex != 0 {
				minDistance = min(minDistance, index-preIndex)
			}
			preIndex = index
		}
		if pre.Val > cur.Val && cur.Val < cur.Next.Val {
			if firstIndex == 0 {
				firstIndex = index
			}
			if preIndex != 0 {
				minDistance = min(minDistance, index-preIndex)
			}
			preIndex = index
		}
		cur = cur.Next
		pre = pre.Next
		index++
	}
	if minDistance == 1000000 {
		return []int{-1, -1}
	}

	return []int{minDistance, preIndex - firstIndex}
}
