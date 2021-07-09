package twentyfortyeight

import (
	"math/rand"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	b := New()
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	var free int
	var occupied int
	for _, line := range b.B {
		for _, item := range line {
			if item != 0 {
				occupied++
			} else {
				free++
			}
		}
	}
	if free != 14 {
		t.Errorf("twentyfortyeight.New() free items was %d, expected 14", free)
	}
	if occupied != 2 {
		t.Errorf("twentyfortyeight.New() occupied items was %d, expected 2", occupied)
	}
}

func TestMove(t *testing.T) {
	s := []struct {
		b          *Board
		direction  direction
		expectedOk bool
		expected   *Board
	}{
		{
			fromArray([4][4]int8{
				{2, 0, 0, 0},
				{2, 0, 4, 0},
				{16, 8, 4, 2},
				{0, 2, 8, 0}}),
			Right,
			true,
			fromArray([4][4]int8{
				{0, 0, 0, 2},
				{0, 0, 2, 4},
				{16, 8, 4, 2},
				{0, 0, 2, 8}}),
		},
		{
			fromArray([4][4]int8{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}}),
			Right,
			false,
			nil,
		},
		{
			fromArray([4][4]int8{
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{2, 4, 8, 16}}),
			Right,
			false,
			nil,
		},
		{
			fromArray([4][4]int8{
				{2, 2, 0, 0},
				{0, 4, 4, 0},
				{0, 8, 0, 8},
				{4, 4, 4, 4}}),
			Right,
			true,
			fromArray([4][4]int8{
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
				{0, 0, 8, 8}}),
		}}
	for _, s := range s {
		res, ok := s.b.move(s.direction)
		if ok != s.expectedOk {
			t.Errorf("move(Right) on:\n%v was %v, expected %v", s.b, ok, s.expectedOk)
		}
		if ok && s.expectedOk && res.B != s.expected.B {
			t.Errorf("move(Right) on:\n%v returned:\n%v expected:\n%v", s.b, res, s.expected)
		}
	}
}
