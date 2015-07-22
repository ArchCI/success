package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("Add test fail")
	} else {
	  t.Log("Add test passed")
	}
}
