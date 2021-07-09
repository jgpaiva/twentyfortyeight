package twentyfortyeight

import (
	"fmt"
	"math/rand"
)

type Board struct {
	B    [4][4]int8
	free int8
}

func New() Board {
	b := Board{free: 16}
	b = b.nextBoard()
	b = b.nextBoard()
	return b
}

func fromArray(b [4][4]int8) *Board {
	r := Board{B: b}
	r.updateFree()
	return &r
}

func (b *Board) updateFree() {
	b.free = 16
	for _, line := range b.B {
		for _, item := range line {
			if item != 0 {
				b.free--
			}
		}
	}
}

func (b *Board) nextBoard() Board {
	spot := int8(rand.Intn(int(b.free)))
	var value int8 = 2
	// 4 is supposed to happen with probability 1/10
	if rand.Intn(10) == 9 {
		value = 4
	}
	var count int8
	for i, line := range b.B {
		for j, item := range line {
			if item == 0 {
				if count == spot {
					b.B[i][j] = value
				}
				count += 1
			}
		}
	}
	b.free -= 1
	return *b
}

type direction int

const (
	Up    = iota
	Right = iota
	Down  = iota
	Left  = iota
)

func (b *Board) move(direction direction) (ret *Board, moved bool) {
	ret = fromArray(b.B)
	if direction == Right {
		for i := range ret.B {
			target := 3
			// move everything to the right
			for j := 3; j >= 0; j-- {
				if ret.B[i][j] != 0 {
					if target != j {
						ret.B[i][target] = ret.B[i][j]
						ret.B[i][j] = 0
						moved = true
					}
					target--
				}
			}
			needMove := false
			// join adjacent
			for j := 3; j > 0; j-- {
				if ret.B[i][j] == ret.B[i][j-1] {
					ret.B[i][j] = ret.B[i][j] * 2
					ret.B[i][j-1] = 0
					needMove = true
				}
			}
			// if anything was joined, move to the right again
			if needMove {
				target := 3
				for j := 3; j >= 0; j-- {
					if ret.B[i][j] != 0 {
						if target != j {
							ret.B[i][target] = ret.B[i][j]
							ret.B[i][j] = 0
							moved = true
						}
						target--
					}
				}
			}
		}
	}
	return ret, moved
}

func (b *Board) String() (ret string) {
	for _, line := range b.B {
		for _, value := range line {
			if value == 0 {
				ret += "   ."
			} else {
				ret += fmt.Sprintf("%4d", value)
			}
		}
		ret += "\n"
	}
	return
}
