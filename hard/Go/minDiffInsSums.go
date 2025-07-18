package go_test



import (
	"container/heap"
)
// This code defines a solution to the problem of finding the minimum minimumDifference	
// between the sums of two groups of numbers in an array, where the array is divided into
// three equal parts. The solution uses two heaps: a max-heap for the left parts
// and a min-heap for the right parts. It calculates the sums of the two groups and finds
// the minimum difference between them.
// The function minimumDifference takes an array of integers and returns the minimum minimumDifference
// between the sums of two groups formed by dividing the array into three equal parts.
// The algorithm uses heaps to efficiently manage the sums of the two groups.
 



type MaxHeap []int
func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] } 
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

type MinHeap []int
func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] } 
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func minimumDifference(nums []int) int64 {
	n := len(nums) / 3
	leftSums := make([]int64, 2*n+1)
	rightSums := make([]int64, 2*n+1)

	leftHeap := &MaxHeap{}
	heap.Init(leftHeap)
	leftSum := int64(0)

	for i := 0; i < 2*n; i++ {
		heap.Push(leftHeap, nums[i])
		leftSum += int64(nums[i])
		if leftHeap.Len() > n {
			leftSum -= int64(heap.Pop(leftHeap).(int))
		}
		if leftHeap.Len() == n {
			leftSums[i] = leftSum
		}
	}

	rightHeap := &MinHeap{}
	heap.Init(rightHeap)
	rightSum := int64(0)

	for i := 3*n - 1; i >= n; i-- {
		heap.Push(rightHeap, nums[i])
		rightSum += int64(nums[i])
		if rightHeap.Len() > n {
			rightSum -= int64(heap.Pop(rightHeap).(int))
		}
		if rightHeap.Len() == n {
			rightSums[i] = rightSum
		}
	}

	
	res := int64(1<<63 - 1)
	for i := n - 1; i < 2*n; i++ {
		if rightSums[i+1] > 0 {
			diff := leftSums[i] - rightSums[i+1]
			if diff < res {
				res = diff
			}
		}
	}

	return res
}

