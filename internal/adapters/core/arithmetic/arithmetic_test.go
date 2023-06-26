package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
		require.Equal(t, answer, int32(2))
	}
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
		require.Equal(t, answer, int32(0))
	}
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, _ := arith.Multiplication(3, 3)
	//if err != nil {
	//	t.Fatalf("Expected: %v, got: %v", nil, err)
	//	require.Equal(t, answer, int32(6))
	//}
	if answer != 9 {
		t.Fatalf("ERROR. Expected: %v, got: %v", 9, answer)
	}
	t.Logf("PASS. Expected: %v, got: %v", 9, answer)
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(4, 2)
	if err != nil {
		t.Fatalf("Expected: %v, got: %v", nil, err)
		require.Equal(t, answer, int32(2))
	}
}
