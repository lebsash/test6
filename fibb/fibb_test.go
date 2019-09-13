package fibb

import (
	"testing"
)

func TestFibb(t *testing.T) {

	res := []int64{0, 1, 1, 2, 3, 5, 8, 13}

	Calc("1", 10)

	var val []int64
	result := false
	wrong := false
	for result == false {
		val, result = Status("1")
	}

	for i, item := range val {
		if res[i] != item {
			wrong = true
		}
	}

	if wrong {
		t.Errorf("Wrong result, \n Expected: %v \n Getted: %v", res, val)
	}

}
