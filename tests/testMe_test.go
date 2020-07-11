package tests

import "testing"

func TestF1(t *testing.T) {
	if F1(0) != 0 {
		t.Error(`f1(0) != 0`)
	}
	if F1(1) != 1 {
		t.Error(`f1(1) != 1`)
	}

	if F1(2) != 1 {
		t.Error(`f1(2) != 1`)
	}

	if F1(10) != 55 {
		t.Error(`f1(10) != 55`)
	}
}

func TestF2(t *testing.T) {
	if F2(0) != 0 {
		t.Error(`f1(0) != 0`)
	}
	if F2(1) != 1 {
		t.Error(`f1(1) != 1`)
	}

	if F2(2) != 1 {
		t.Error(`f1(2) != 1`)
	}

	if F2(10) != 55 {
		t.Error(`f1(10) != 55`)
	}
}

func TestS1(t *testing.T) {
	if S1("123456789") != 9 {
		t.Error(`s1("123456789") != 9`)
	}

	if S1("") != 0 {
		t.Error(`s1("") != 0`)
	}
}

func TestS2(t *testing.T) {
	if S2("123456789") != 9 {
		t.Error(`s2("123456789") != 9`)
	}

	if S2("") != 0 {
		t.Error(`s2("") != 0`)
	}
}
