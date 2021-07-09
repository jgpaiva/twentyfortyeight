package main

import (
	"fmt"
	"math/rand"
	"time"

	"jgpaiva.com/2048/twentyfortyeight"
)

func main() {
	fmt.Println("Welcome to 2048")
	rand.Seed(time.Now().UTC().UnixNano())
	b := twentyfortyeight.New()
	fmt.Println(&b)
}
