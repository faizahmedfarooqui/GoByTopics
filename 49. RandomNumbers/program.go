package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	p := fmt.Println

	// 0 <= i < 100
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	p()

	// 0.0 <= f < 1.0
	p(rand.Float64())

	// 5.0 <= f' < 10.0
	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)
	p()

	// The default number generator is deterministic, so it’ll produce the
	// same sequence of numbers each time by default. To produce varying
	// sequences, give it a seed that changes. Note that this is not safe
	// to use for random numbers you intend to be secret, use crypto/rand
	// for those.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	p()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	p()

	// If you seed a source with the same number, it produces the same
	// sequence of random numbers.
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))
}
