package stylus

import (
	"fmt"
	"testing"
)

func TestCompile(t *testing.T) {
	pathc, errc := Compile("test/1.styl")

	var path string

	select {
	case path = <-pathc:
	case err := <-errc:
		t.Error(err)
	}

	fmt.Println(string(path))
}
