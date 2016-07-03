package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	meY    = 12
	enemyY = 12
	ball   = []int{40, 12}
	vector = []int{-1, 1}
	level  = 0
	score  = []int{0, 0}
	shadow [][]int
	ipAddr string
	clear  = false
)

const (
	wallLeft   = 0
	wallRight  = 79
	wallTop    = 1
	wallBottom = 23
	meX        = 76
	enemyX     = 2
	bar        = 4
)

func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}

func draw() {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if clear == true {
		return
	}
	drawLine(wallLeft, wallTop-1, fmt.Sprintf("                                                                     %03d - %03d", score[0], score[1]))
	drawLine(wallLeft, wallTop, fmt.Sprintf("--------------------------------------------------------------------------------"))
	drawLine(wallLeft, wallBottom, fmt.Sprintf("--------------------------------------------------------------------------------"))
	drawLine(wallLeft, wallBottom+1, fmt.Sprintf("EXIT : ESC KEY"))
	drawLine(ball[0], ball[1], fmt.Sprintf("*"))

	for i := range shadow {
		drawLine(shadow[i][0], shadow[i][1], fmt.Sprintf(string(ipAddr[len(ipAddr)-i-1])))

	}

	for i := 0; i < bar; i++ {
		drawLine(meX, meY+i, fmt.Sprintf("||"))
		drawLine(enemyX, enemyY+i, fmt.Sprintf("||"))
	}
	termbox.Flush()
}

func keyEvent(c chan bool) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				clear = true
				c <- false
				return
			case termbox.KeyArrowUp:
				if meY > wallTop+1 {
					meY--
				}
			case termbox.KeyArrowDown:
				if meY < wallBottom-bar {
					meY++
				}

			default:
			}
			c <- true
		default:
		}
		hitTest()
	}
}

func moveBall(c chan bool) {
	for {

		ball[0] += vector[0]
		ball[1] += vector[1]
		hitTest()
		c <- true

		ball[0] += vector[0]
		hitTest()
		recMove()
		c <- true

		if ball[1] <= wallTop+1 || ball[1] >= wallBottom-1 {
			vector[1] *= -1
		}

		if ball[0] <= wallLeft+1 || ball[0] >= wallRight-1 {
			vector[0] *= -1
			if ball[0] <= wallLeft+1 {
				score[1]++
			}
			if ball[0] >= wallRight-1 {
				score[0]++
			}
			initGame()
			c <- true
			time.Sleep(time.Duration(500) * time.Millisecond)
		}
		c <- true
		time.Sleep(time.Duration(100-level*5) * time.Millisecond)
	}
}
func recMove() {
	shadow = append(shadow, []int{ball[0], ball[1]})
	shadow = shadow[1:]
}
func moveEnemy(c chan bool) {
	vec := 0
	for {

		vec = ball[1] - (enemyY + 2)

		if enemyY > wallTop+1 && vec < 0 {
			enemyY--
		}
		if enemyY < wallBottom-bar && vec > 0 {
			enemyY++
		}
		hitTest()
		c <- true

		switch true {
		case ball[0] < 30:
			time.Sleep(100 * time.Millisecond)
		case ball[0] < 50:
			time.Sleep(150 * time.Millisecond)
		case ball[0] < 80:
			time.Sleep(180 * time.Millisecond)
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func hitTest() {
	if vector[0] == 1 && ball[0] > wallRight-10 {
		if (ball[0] == meX || ball[0] == meX-1) && ball[1] >= meY && ball[1] <= meY+bar {
			vector[0] *= -1
			level = (level + 1) % 10
		}
	}
	if vector[0] == -1 && ball[0] < wallLeft+10 {
		if (ball[0] == enemyX+1 || ball[0] == enemyX+2) && ball[1] >= enemyY && ball[1] <= enemyY+bar {
			vector[0] *= -1
			level = (level + 1) % 10
		}
	}
}

func initGame() {
	level = 0
	rand.Seed(time.Now().UnixNano())
	ball = []int{40, 5 + rand.Intn(15)}
	if len(os.Args) >= 2 {
		ipAddr = os.Args[1]
	} else {
		ipAddr = ""
	}
	shadow = make([][]int, len(ipAddr))
	for i := range shadow {
		shadow[i] = []int{ball[0], ball[1]}
	}
}

func main() {
	initGame()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	signal := make(chan bool)

	go moveBall(signal)
	go moveEnemy(signal)
	go keyEvent(signal)

	for {
		s := <-signal
		draw()
		if s == false {
			close(signal)
			break
		}
	}

	defer termbox.Close()

}
