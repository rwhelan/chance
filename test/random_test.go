package random_test

import (
	"testing"

	chance "github.com/rwhelan/chance/pkg"
)

func TestRandomInt(t *testing.T) {
	chance.RandomInt(65536)
}
