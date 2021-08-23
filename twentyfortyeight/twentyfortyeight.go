package twentyfortyeight

import (
	"fmt"
	"math/rand"
)

type Board struct {
	B    [4][4]int16
	free int16
}

func New() Board {
	b := Board{free: 16}
	b = b.NextBoard()
	b = b.NextBoard()
	return b
}

func fromArray(b [4][4]int16) *Board {
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

func (b *Board) NextBoard() Board {
	spot := int16(rand.Intn(int(b.free)))
	var value int16 = 2
	// 4 is supposed to happen with probability 1/10
	if rand.Intn(10) == 9 {
		value = 4
	}
	var count int16
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

func (d direction) String() string {
	if d == Up {
		return "Up"
	} else if d == Down {
		return "Down"
	} else if d == Right {
		return "Right"
	} else if d == Left {
		return "Left"
	}
	panic("unreachable")
}

func moveRight(line *[4]int16, moved *bool) {
	target := 3
	for j := 3; j >= 0; j-- {
		if line[j] != 0 {
			if target != j {
				line[target] = line[j]
				line[j] = 0
				*moved = true
			}
			target--
		}
	}
}

func moveLeft(line *[4]int16, moved *bool) {
	target := 0
	for j := 0; j <= 3; j++ {
		if line[j] != 0 {
			if target != j {
				line[target] = line[j]
				line[j] = 0
				*moved = true
			}
			target++
		}
	}
}
func moveDown(b *[4][4]int16, moved *bool, i int) {
	target := 3
	for j := 3; j >= 0; j-- {
		if b[j][i] != 0 {
			if target != j {
				b[target][i] = b[j][i]
				b[j][i] = 0
				*moved = true
			}
			target--
		}
	}
}
func moveUp(b *[4][4]int16, moved *bool, i int) {
	target := 0
	for j := 0; j <= 3; j++ {
		if b[j][i] != 0 {
			if target != j {
				b[target][i] = b[j][i]
				b[j][i] = 0
				*moved = true
			}
			target++
		}
	}
}

func (b *Board) Move(direction direction) (ret *Board, moved bool) {
	ret = fromArray(b.B)
	if direction == Right {
		for i := range ret.B {
			moveRight(&ret.B[i], &moved)
			needMove := false
			// join adjacent
			for j := 3; j > 0; j-- {
				if ret.B[i][j] == ret.B[i][j-1] && ret.B[i][j] != 0 {
					ret.B[i][j] = ret.B[i][j] * 2
					ret.B[i][j-1] = 0
					needMove = true
					moved = true
					ret.free++
				}
			}
			// if anything was joined, move right again
			if needMove {
				moveRight(&ret.B[i], &moved)
			}
		}
	}
	if direction == Left {
		for i := range ret.B {
			moveLeft(&ret.B[i], &moved)
			needMove := false
			// join adjacent
			for j := 0; j < 3; j++ {
				if ret.B[i][j] == ret.B[i][j+1] && ret.B[i][j] != 0 {
					ret.B[i][j] = ret.B[i][j] * 2
					ret.B[i][j+1] = 0
					needMove = true
					moved = true
					ret.free++
				}
			}
			// if anything was joined, move left again
			if needMove {
				moveLeft(&ret.B[i], &moved)
			}
		}
	}
	if direction == Down {
		for i := range ret.B {
			moveDown(&ret.B, &moved, i)
			needMove := false
			// join adjacent
			for j := 3; j > 0; j-- {
				if ret.B[j][i] == ret.B[j-1][i] && ret.B[j][i] != 0 {
					ret.B[j][i] = ret.B[j][i] * 2
					ret.B[j-1][i] = 0
					needMove = true
					moved = true
					ret.free++
				}
			}
			// if anything was joined, move down again
			if needMove {
				moveDown(&ret.B, &moved, i)
			}
		}
	}
	if direction == Up {
		for i := range ret.B {
			moveUp(&ret.B, &moved, i)
			needMove := false
			// join adjacent
			for j := 0; j < 3; j++ {
				if ret.B[j][i] == ret.B[j+1][i] && ret.B[j][i] != 0 {
					ret.B[j][i] = ret.B[j][i] * 2
					ret.B[j+1][i] = 0
					needMove = true
					moved = true
					ret.free++
				}
			}
			// if anything was joined, move down again
			if needMove {
				moveUp(&ret.B, &moved, i)
			}
		}
	}
	return ret, moved
}

func (b *Board) String() (ret string) {
	for _, line := range b.B {
		for _, value := range line {
			if value == 0 {
				ret += "    ."
			} else {
				ret += fmt.Sprintf("%5d", value)
			}
		}
		ret += "\n"
	}
	return
}
