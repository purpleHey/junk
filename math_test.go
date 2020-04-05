package math

import (
	"testing"

	"/github.com/purple/junk"
)

func TestSum(t *testing.T)  {
	sum := math.Sum([]int{5, 8, -3})

	if sum != 10 {
		t.Errorf("Expected 10, got %d", sum)
	}
}