package random_test

import (
	"fmt"
	"testing"

	chance "github.com/rwhelan/chance/pkg"
)

func TestRandomInt(t *testing.T) {
	v, err := chance.RandomInt(65536)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("FOO", v)
}
