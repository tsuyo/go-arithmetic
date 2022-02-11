package arithmetic

import "math"

// Checks if a number is prime or not
func IsPrime(num int) bool {
	upto := int(math.Sqrt(float64(num)))
	for i := 2; i <= upto; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
