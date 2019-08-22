package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a" ,6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}


func ExampleRepeat() {
	s := Repeat("a",6)
	fmt.Print(s)
	//Output:aaaaaa
}

