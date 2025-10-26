
package mostFrequent

func isTrionic(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	p := 0
	for p+1 < n && nums[p] < nums[p+1] {
		p++
	}

	if p == 0 || p == n-1 {
		return false
	}

	q := p
	for q+1 < n && nums[q] > nums[q+1] {
		q++
	}

	if q == p || q == n-1 {
		return false
	}

	// strictly increasing again
	for q+1 < n && nums[q] < nums[q+1] {
		q++
	}

	return q == n-1
}

