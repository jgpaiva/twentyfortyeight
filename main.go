package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"jgpaiva.com/2048/twentyfortyeight"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to 2048")
	rand.Seed(time.Now().UTC().UnixNano())
	b := twentyfortyeight.New()
	fmt.Println(&b)
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
		fmt.Println(&b)
		fmt.Print("> ")
	}
}
