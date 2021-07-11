package twentyfortyeight

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDirectionString(t *testing.T) {
	directions := []struct {
		val      direction
		expected string
	}{
		{Right, "Right"},
		{Left, "Left"},
		{Up, "Up"},
		{Down, "Down"},
	}
	for _, dir := range directions {
		if res := fmt.Sprintf("%v", dir.val); res != dir.expected {
			t.Errorf("fmt.Sprintf(\"%%v\", %v) got %v, expected %v", dir.val, res, dir.expected)
		}
	}
}

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
			fromArray([4][4]int16{
				{2, 0, 0, 0},
				{2, 0, 4, 0},
				{16, 8, 4, 2},
				{0, 2, 8, 0}}),
			Right,
			true,
			fromArray([4][4]int16{
				{0, 0, 0, 2},
				{0, 0, 2, 4},
				{16, 8, 4, 2},
				{0, 0, 2, 8}}),
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 2},
				{0, 2, 0, 4},
				{16, 8, 4, 2},
				{0, 2, 8, 0}}),
			Left,
			true,
			fromArray([4][4]int16{
				{2, 0, 0, 0},
				{2, 4, 0, 0},
				{16, 8, 4, 2},
				{2, 8, 0, 0}}),
		},
		{
			fromArray([4][4]int16{
				{2, 4, 2, 0},
				{0, 0, 4, 8},
				{0, 2, 8, 2},
				{0, 0, 16, 0}}),
			Down,
			true,
			fromArray([4][4]int16{
				{0, 0, 2, 0},
				{0, 0, 4, 0},
				{0, 4, 8, 8},
				{2, 2, 16, 2}}),
		},
		{
			fromArray([4][4]int16{
				{2, 0, 0, 0},
				{2, 0, 0, 0},
				{4, 2, 0, 0},
				{8, 4, 0, 0}}),
			Down,
			true,
			fromArray([4][4]int16{
				{0, 0, 0, 0},
				{4, 0, 0, 0},
				{4, 2, 0, 0},
				{8, 4, 0, 0}}),
		},
		{
			fromArray([4][4]int16{
				{0, 0, 2, 0},
				{0, 4, 4, 8},
				{0, 0, 8, 2},
				{2, 2, 16, 0}}),
			Up,
			true,
			fromArray([4][4]int16{
				{2, 4, 2, 8},
				{0, 2, 4, 2},
				{0, 0, 8, 0},
				{0, 0, 16, 0}}),
		},
		{
			fromArray([4][4]int16{
				{4, 16, 8, 0},
				{64, 4, 0, 0},
				{128, 2, 0, 0},
				{256, 2, 0, 0}}),
			Up,
			true,
			fromArray([4][4]int16{
				{4, 16, 8, 0},
				{64, 4, 0, 0},
				{128, 4, 0, 0},
				{256, 0, 0, 0}}),
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}}),
			Right,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}}),
			Left,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}}),
			Down,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}}),
			Up,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 2},
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{2, 4, 8, 16}}),
			Right,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{2, 0, 0, 0},
				{4, 0, 0, 0},
				{8, 0, 0, 0},
				{2, 4, 8, 16}}),
			Left,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{0, 0, 0, 18},
				{0, 0, 0, 8},
				{0, 0, 0, 4},
				{2, 4, 8, 2}}),
			Down,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{2, 4, 8, 18},
				{0, 0, 0, 8},
				{0, 0, 0, 4},
				{0, 0, 0, 2}}),
			Up,
			false,
			nil,
		},
		{
			fromArray([4][4]int16{
				{2, 2, 0, 0},
				{0, 4, 4, 0},
				{0, 8, 0, 8},
				{4, 4, 4, 4}}),
			Right,
			true,
			fromArray([4][4]int16{
				{0, 0, 0, 4},
				{0, 0, 0, 8},
				{0, 0, 0, 16},
				{0, 0, 8, 8}}),
		},
		{
			fromArray([4][4]int16{
				{0, 0, 2, 2},
				{0, 4, 4, 0},
				{2, 8, 0, 8},
				{4, 4, 4, 4}}),
			Left,
			true,
			fromArray([4][4]int16{
				{4, 0, 0, 0},
				{8, 0, 0, 0},
				{2, 16, 0, 0},
				{8, 8, 0, 0}}),
		},
		{
			fromArray([4][4]int16{
				{4, 2, 0, 0},
				{4, 8, 4, 0},
				{4, 0, 4, 2},
				{4, 8, 0, 2}}),
			Down,
			true,
			fromArray([4][4]int16{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{8, 2, 0, 0},
				{8, 16, 8, 4}}),
		},
		{
			fromArray([4][4]int16{
				{4, 8, 0, 2},
				{4, 8, 4, 2},
				{4, 0, 4, 0},
				{4, 2, 0, 0}}),
			Up,
			true,
			fromArray([4][4]int16{
				{8, 16, 8, 4},
				{8, 2, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}}),
		}}
	for _, s := range s {
		res, ok := s.b.Move(s.direction)
		if ok != s.expectedOk {
			t.Errorf("move(%v) on:\n%v OK was %v, expected %v", s.direction, s.b, ok, s.expectedOk)
		}
		if ok && s.expectedOk && res.B != s.expected.B {
			t.Errorf("move(%v) on:\n%v returned:\n%v expected:\n%v", s.direction, s.b, res, s.expected)
		}
	}
}
