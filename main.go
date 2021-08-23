package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
	"jgpaiva.com/2048/twentyfortyeight"
)

func printBoard(b *twentyfortyeight.Board) (ret string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	two := color.New(color.BgRed).SprintFunc()
	four := color.New(color.BgGreen).SprintFunc()
	eight := color.New(color.BgYellow).SprintFunc()
	sixteen := color.New(color.BgBlue).SprintFunc()
	thirtytwo := color.New(color.BgMagenta).SprintFunc()
	sixtyfour := color.New(color.BgCyan).SprintFunc()
	for _, line := range b.B {
		for _, value := range line {
			if value == 0 {
				ret += yellow("    .")
			} else {
				ret += " "
				if value == 2 {
					ret += two(fmt.Sprintf("%4d", value))
				} else if value == 4 {
					ret += four(fmt.Sprintf("%4d", value))
				} else if value == 8 {
					ret += eight(fmt.Sprintf("%4d", value))
				} else if value == 16 {
					ret += sixteen(fmt.Sprintf("%4d", value))
				} else if value == 32 {
					ret += thirtytwo(fmt.Sprintf("%4d", value))
				} else if value == 64 {
					ret += sixtyfour(fmt.Sprintf("%4d", value))
				} else {
					ret += fmt.Sprintf("%4d", value)
				}
			}
		}
		ret += "\n"
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to 2048")
	rand.Seed(time.Now().UTC().UnixNano())
	b := twentyfortyeight.New()
	fmt.Println(printBoard(&b))
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Up" || text == "up" || text == "w" {
			if n, ok := b.Move(twentyfortyeight.Up); ok {
				b = n.NextBoard()
			}
		} else if text == "Down" || text == "down" || text == "s" {
			if n, ok := b.Move(twentyfortyeight.Down); ok {
				b = n.NextBoard()
			}
		} else if text == "Left" || text == "left" || text == "a" {
			if n, ok := b.Move(twentyfortyeight.Left); ok {
				b = n.NextBoard()
			}
		} else if text == "Right" || text == "right" || text == "d" {
			if n, ok := b.Move(twentyfortyeight.Right); ok {
				b = n.NextBoard()
			}
		} else {
			fmt.Println("invalid input: ", text)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(printBoard(&b))
		fmt.Print("> ")
	}
}
