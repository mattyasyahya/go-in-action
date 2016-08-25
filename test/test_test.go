package test

import "testing"

// TestSimple test simple for fun
func TestSimple(t *testing.T) {
	t.Log("Given Sum() function")
	{
		t.Log("\tWhen value 1 and 2")
		{
			sum := Sum(1, 2)
			if sum != 3 {
				t.Fatal("\t\tShould get result 3")
			}
			t.Log("\t\tShould get result 3")
		}
	}
}

// TestTable test using table test
func TestTable(t *testing.T) {
	scenario := []struct {
		a      int
		b      int
		result int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
	}

	t.Log("Given Sum() function")
	{
		for _, value := range scenario {
			t.Logf("\tWhen value %d and %d", value.a, value.b)
			{
				sum := Sum(value.a, value.b)
				if sum != value.result {
					t.Fatalf("\t\tShould get result %d", value.result)
				}
				t.Logf("\t\tShould get result %d", value.result)
			}
		}
	}
}
