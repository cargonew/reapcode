package easy

func findTarget(root *TreeNode, k int) bool {
    nums := []int{}

    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        nums = append(nums, node.Val)
        inorder(node.Right)
    }

    inorder(root)

    left, right := 0, len(nums)-1
    for left < right {
        sum := nums[left] + nums[right]
        if sum == k {
            return true
        } else if sum < k {
            left++
        } else {
            right--
        }
    }

    return false
}

  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

