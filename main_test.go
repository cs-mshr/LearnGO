package main

import "testing"

var tests = []struct {
	name        string
	dividend    float32
	divisor     float32
	expected    float32
	errExpected bool
}{
	{"valid-data", 100, 10, 10, false},
	{"invalid-data", 100, 0, -1, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		val, err := divide(tt.dividend, tt.divisor)
		if tt.errExpected {
			if err == nil {
				t.Error("Error is expected")
			}
		} else {
			if err != nil {
				t.Error("Error is not expected", err.Error())
			}
		}

		if val != tt.expected {
			t.Errorf("Expected value is %v, but got %v", tt.expected, val)
		}
	}

}

func TestDivide(t *testing.T) {
	_, err := divide(100, 10)
	if err != nil {
		t.Error("Expected error to be nil")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(100, 0)
	if err == nil {
		t.Error("Error is expected")
	}
}
