package mostFrequent

func getMinimumDifference(root *TreeNode) int {
    var prev *TreeNode
    minDiff := 1<<31 - 1  // big number

    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }

        inorder(node.Left)

        if prev != nil {
            diff := node.Val - prev.Val
            if diff < 0 {
                diff = -diff
            }
            if diff < minDiff {
                minDiff = diff
            }
        }
        prev = node

        inorder(node.Right)
    }

    inorder(root)
    return minDiff
}

