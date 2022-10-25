package main

import "testing"

func TestAdd(t *testing.T) {

	var testCases = []struct {
		input1   int
		input2   int
		expected int
	}{
		{1, 2, 3},
		{2, 4, 6},
		{5, 6, 11},
		{36, 36, 72},
	}

	for _, tests := range testCases {
		addition := Add(tests.input1, tests.input2)
		if addition != tests.expected {
			t.Errorf("FAILED: Expected: %d , received: %d\n", tests.expected, addition)
		} else {
			t.Logf("PASSED: Expected %d , received %d", tests.expected, addition)
		}
	}

}
func TestDivide(t *testing.T) {
	var tests = []struct {
		inp1 int
		inp2 int
		exp  int
	}{
		{4, 2, 2},
		{6, 3, 2},
		{9, 2, 4},
	}
	for _, v := range tests {
		division := Divide(v.inp1, v.inp2)
		if division != v.exp {
			t.Errorf("FAILED: Expected %d , received %d", v.exp, division)
		} else {
			t.Logf("PASSED: Expected %d , received %d", v.exp, division)
		}
	}
}
func TestMultiply(t *testing.T) {
	var tests = []struct {
		inp1 int
		inp2 int
		exp  int
	}{
		{33, 45, 1485},
		{12, 56, 672},
		{85, 98, 8330},
	}

	for _, v := range tests {
		pdt := Multiply(v.inp1, v.inp2)
		if pdt != v.exp {
			t.Errorf("FAILED: Expected %d , received %d", v.exp, pdt)
		} else {
			t.Logf("PASSED: Expected %d , received %d", v.exp, pdt)
		}
	}
}
