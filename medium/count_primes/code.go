package count_primes

func countPrimes(n int) int {
	arr := make([]bool, n)
	for i := 2; i < n; i++ {
		arr[i] = true
	}

	for i := 2; i*i < n; i++ {
		if arr[i] {
			for j := i * i; j < n; j += i {
				arr[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if arr[i] {
			count++
		}
	}

	return count
}
