package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum([]int{5, 8, -3})

	if sum != 10 {
		t.Errorf("Expected 10, got %d", sum)
	}
}
