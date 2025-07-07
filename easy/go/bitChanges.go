func minChanges(n int, k int) int {
     if (k | n) != n {
        return -1
    }

    flips := n & ^k
    count := 0
    for flips > 0 {
        count += flips & 1
        flips >>= 1
    }

    return count
}
