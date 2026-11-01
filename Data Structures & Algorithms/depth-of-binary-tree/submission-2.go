/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
    return max(sideDepth(root.Left, 1), sideDepth(root.Right, 1))
}

func sideDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	return max(sideDepth(root.Left, depth), sideDepth(root.Right, depth))
}