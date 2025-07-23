package superPow
import (
	"container/heap"
)

// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/
// Time: O(n log n)
// Space: O(n)
// Given an integer array nums sorted in non-decreasing order, return the sum of all possible
// subarray sums of nums (i.e., the sum of sums of all non-empty subarrays of nums).
// Since the answer may be too large, return it modulo 10^9 + 7.
// A subarray is a contiguous non-empty sequence of elements within an array.	
// Example 1:
// Input: nums = [1,2,3,4]
// Output: 50
// Explanation: The subarrays of nums are [1], [2], [3], [4], [1,2], [1,2,3], [1,2,3,4], [2,3], [2,3,4], [3,4].
// The sums of these subarrays are 1, 2, 3, 4
// 1+2=3, 1+2+3=6, 1+2+3+4=10, 2+3=5, 2+3+4=9, 3+4=7.
// The sum of all subarray sums is 1 + 2 + 3 + 4 + 3 + 6 + 10 + 5 + 9 + 7
// = 50.

const MOD = 1_000_000_007

type Item struct {
	sum int
	start int
	end int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func rangeSum(nums []int, n int, left int, right int) int {
	h := &MinHeap{}

	for i := 0; i < n; i++ {
		heap.Push(h, Item{nums[i], i, i})
	}

	res := 0
	for count := 1; count <= right; count++ {
		curr := heap.Pop(h).(Item)
		if count >= left {
			res = (res + curr.sum) % MOD
		}
		if curr.end+1 < n {
			newSum := curr.sum + nums[curr.end+1]
			heap.Push(h, Item{newSum, curr.start, curr.end + 1})
		}
	}

	return res
}



 



