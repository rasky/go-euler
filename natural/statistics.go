package natural

// Returns the i-th permutation of the array [1,2,...,n]
func IthPermutation(n, i int64) []int64 {
	fact := make([]int64, n)
	perm := make([]int64, n)

	fact[0] = int64(1)
	for k := int64(1); k < n; k++ {
		fact[k] = fact[k-1] * k
	}

	for k := int64(0); k < n; k++ {
		perm[k] = i / fact[n-1-k]
		i = i % fact[n-1-k]
	}

	for k := n - 1; k > 0; k-- {
		for j := k - 1; j >= 0; j-- {
			if perm[j] <= perm[k] {
				perm[k]++
			}
		}
	}

	return perm
}

// Compute the next permutation of the array (the first permutaiton is assumed
// to be the sorted one). Return false if there are no more permutations left
func NextPermutation(data []int64) bool {
	n := len(data)
	var k, l int
	for k = n - 2; k >= 0; k-- {
		if data[k] < data[k+1] {
			break
		}
	}
	if k < 0 {
		return false
	}
	for l = n - 1; !(data[k] < data[l]); l-- {

	}
	data[k], data[l] = data[l], data[k]
	for i, j := k+1, n-1; i < j; i++ {
		data[i], data[j] = data[j], data[i]
		j--
	}
	return true
}

// Compute the next k-combination of the array.
func NextCombination(data []int64, k int64) bool {
	n := int64(len(data))

	var h, t int64
	for h = k - 1; h >= 0; h-- {
		if data[h] < data[n-1] {
			break
		}
	}
	if h < 0 {
		return false
	}

	for t = n - 1; t > k; t-- {
		if data[t-1] <= data[h] {
			break
		}
	}

	data[h], data[t] = data[t], data[h]
	h++
	t++
	if h < k && t < n {
		j := t
		for h != j {
			data[h], data[j] = data[j], data[h]
			h++
			j++
			if h == k {
				h = t
			}
			if j == n {
				j = t
			} else if h == t {
				t = j
			}
		}
	}
	return true
}
