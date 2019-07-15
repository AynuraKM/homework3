package main

import (
	"fmt"
)

func main() {
	b := Bird{}
	b.color = "black"
	b.weight = 228
	b.tail.length = 322
	b.tail.width = 233
	b.wing.length = 13.37
	b.wing.width = 33.31
	fmt.Println(b)
	fmt.Println(b.Distance(Wing{13, 33}, 2))

}

type (
	Wing struct {
		width float32
		length float32
	}

	Tail struct {
		width  float32
		length float32
	}
)

type (
	Bird struct {
		tail   Tail
		wing   Wing
		color string
		weight float32
	}
)

func (b *Bird) Distance(wing Wing, t int) int {
	// если птица большая, то скорость полета150км/ч
	v := 1
	s := b.wing.length * b.wing.width
	if s > 100 {
		v = t * 150
	}
	return v
}

