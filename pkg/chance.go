package chance

import (
	"math/rand"
	"time"
)

type Option struct {
	Name   string
	Weight int
}

type Selector struct {
	r       *rand.Rand
	options []*Option
}

func NewSelector() *Selector {
	return &Selector{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Selector) AddOption(option ...*Option) *Selector {
	s.options = append(s.options, option...)
	return s
}

func (s *Selector) GetOption() *Option {
	c := make([]int, len(s.options))
	c[0] = s.options[0].Weight

	for i := 1; i < len(s.options); i++ {
		c[i] = c[i-1] + s.options[i].Weight
	}

	rint := s.r.Intn(c[len(c)-1])
	for i, v := range c {
		if rint < v {
			return s.options[i]
		}
	}

	// fmt.Println("THIS SHOULDN'T HAPPEN")
	return s.options[len(s.options)-1]
}

func Percent(i int) func() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() bool {
		return r.Intn(100) < i
	}
}

func Chance(s, i int) func() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() bool {
		return r.Intn(s) < i
	}
}
