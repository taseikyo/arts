/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-24 15:35:06
 * @link    github.com/taseikyo
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ca, cb := 0, 0
	pa, pb := headA, headB
	for pa != nil {
		ca++
		pa = pa.Next
	}
	for pb != nil {
		cb++
		pb = pb.Next
	}
	pa, pb = headA, headB
	if ca > cb {
		for i := ca - cb; i > 0; i-- {
			pa = pa.Next
		}
	} else if ca < cb {
		for i := cb - ca; i > 0; i-- {
			pb = pb.Next
		}
	}

	for pa != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}

	return nil
}
