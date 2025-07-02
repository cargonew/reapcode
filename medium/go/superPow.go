
package superPow

func superPow(a int, b []int) int {
	const MOD = 1337

	modPow := func(x, n, mod int) int {
		result := 1
		x %= mod
		for n > 0 {
			if n%2 == 1 {
				result = (result * x) % mod
			}
			x = (x * x) % mod
			n /= 2
		}
		return result
	}

	result := 1
	for _, digit := range b {
		result = modPow(result, 10, MOD) * modPow(a, digit, MOD) % MOD
	}

	return result
}

