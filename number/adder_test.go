package number

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4
	if sum != expect {
		t.Errorf("expext'%d' but got '%d'", expect, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Print(sum)
	//Output:6
}