package main

import "testing"

func TestSimple(t *testing.T) {
	t.Log("Given Sum() function")
	{
		t.Log("\tWhen value 1 and 2")
		{
			sum := 1 + 2
			if sum != 3 {
				t.Fatal("\t\tShould get result 3")
			}
			t.Log("\t\tShould get result 3")
		}
	}
}
