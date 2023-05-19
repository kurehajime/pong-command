// main.go
package main

import "flag"

func main() {
	ipAddr := flag.String("ball", "*", "the characters representing the ball")
    target := flag.Int("target", 0, "the first to reach this score wins. 0 means the game is infinite")
    flag.Parse()
	start(*ipAddr, *target)
}
