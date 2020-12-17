package sto

import (
	"testing"
)

func Test01(t *testing.T) {
	data := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(data)
	if res != 2 && res != 3 {
		t.Fail()
	}

	res2 := findRepeatNumber2(data)
	if res2 != 2 && res2 != 3 {
		t.Fail()
	}
}