package restfight

import (
	"time"
)

// Generate key.
func generateKey(base int, length int) int {

	return int(time.Now().UnixNano())

}
