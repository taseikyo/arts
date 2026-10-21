/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2026-08-26 15:16:19
 * @link    github.com/taseikyo
 */

func verifyTreeOrder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	return verify(postorder, 0, len(postorder)-1)
}

func verify(postorder []int, left, right int) bool {
	// 递归终止条件：区间内没有节点或只有一个节点
	if left >= right {
		return true
	}

	root := postorder[right] // 根节点是区间最后一个元素

	// 1. 找到左右子树的分界点：第一个大于根节点的位置
	mid := left
	for mid < right && postorder[mid] < root {
		mid++
	}

	// 2. 检查右子树中的所有节点是否都大于根节点
	for i := mid; i < right; i++ {
		if postorder[i] < root {
			return false
		}
	}

	// 3. 递归验证左子树 [left, mid-1] 和右子树 [mid, right-1]
	return verify(postorder, left, mid-1) && verify(postorder, mid, right-1)
}
