package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0
	if n == 1 {
		return 0, nil
	}
	if n <= 0 {
		return 0, errors.New("n must be positive integer")
	}
	for i := 0; n > 1; i++ {
		if n%2 == 0 {
			n /= 2
			count++
		} else {
			n = n*3 + 1
			count++
		}
	}
	return count, nil
}
