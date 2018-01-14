package restfight

import "math/rand"

// Generate key.
func generateKey(base int, length int) int {
	return base*length + rand.Intn(length/10)
}
