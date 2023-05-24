package easy

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	slc := []int{1, 1, 2}
	res := RemoveDuplicates(slc)
	if res != 2 {
		t.Errorf("wrong, res: %d", res)
	}

	slc = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res = RemoveDuplicates(slc)
	if res != 5 {
		t.Errorf("Wrong, res: %d", res)
	}
}

func TestRemoveDuplicates2(t *testing.T) {
	slc := []int{1, 1, 2}
	res := RemoveDuplicates2(slc)
	if res != 2 {
		t.Errorf("wrong, res: %d", res)
	}

	slc = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res = RemoveDuplicates2(slc)
	if res != 5 {
		t.Errorf("Wrong, res: %d", res)
	}
}
