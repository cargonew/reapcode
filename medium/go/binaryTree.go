/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode) 
    isChild := make(map[int]bool)    

    for _, desc := range descriptions {
        parentVal, childVal, isLeft := desc[0], desc[1], desc[2]

        
        if _, exists := nodes[parentVal]; !exists {
            nodes[parentVal] = &TreeNode{Val: parentVal}
        }
        
        if _, exists := nodes[childVal]; !exists {
            nodes[childVal] = &TreeNode{Val: childVal}
        }

        
        if isLeft == 1 {
            nodes[parentVal].Left = nodes[childVal]
        } else {
            nodes[parentVal].Right = nodes[childVal]
        }

        isChild[childVal] = true
    }

  
    for val, node := range nodes {
        if !isChild[val] {
            return node
        }
    }
    return nil 
}

