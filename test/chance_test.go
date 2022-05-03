package chance_test

import (
	"math"
	"testing"

	chance "github.com/rwhelan/chance/pkg"
)

func IntAbs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func TestPercent(t *testing.T) {
	t.Parallel()
	rounds := 1000000000

	p50 := chance.Percent(50)
	tr, fl := 0, 0

	for i := 0; i < rounds; i++ {
		if p50() {
			tr++
		} else {
			fl++
		}
	}

	delta := IntAbs(tr - fl)
	if float64(delta)/(float64(rounds)/2) > 0.0025 {
		t.Fatalf("Delta outside of acceptable range")
	}
}

func TestSelector(t *testing.T) {
	t.Parallel()
	s := chance.NewSelector().AddOption(
		&chance.Option{Name: "25", Weight: 25},
		&chance.Option{Name: "75", Weight: 75},
	)

	rounds := 100000000
	selects := 0
	for i := 0; i < rounds; i++ {
		if s.GetOption().Name == "25" {
			selects++
		}
	}

	if math.Round(float64(rounds)/float64(selects)) != 4.0 {
		t.Fatalf("Fail")
	}
}
