/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-14 16:44:25
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
	cntA, cntB := 0, 0
	pa, pb := headA, headB
	for pa != nil {
		cntA++
		pa = pa.Next
	}
	for pb != nil {
		cntB++
		pb = pb.Next
	}
	pa, pb = headA, headB
	if cntA > cntB {
		for i := cntA - cntB; i > 0; i-- {
			pa = pa.Next
		}
	} else if cntA < cntB {
		for i := cntB - cntA; i > 0; i-- {
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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
