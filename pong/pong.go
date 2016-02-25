package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	m      *sync.Mutex
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
	WallLeft   = 0
	WallRight  = 79
	WallTop    = 1
	WallBottom = 23
	MeX        = 76
	EnemyX     = 2
	Bar        = 4
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
	drawLine(WallLeft, WallTop-1, fmt.Sprintf("                                                                     %03d - %03d", score[0], score[1]))
	drawLine(WallLeft, WallTop, fmt.Sprintf("--------------------------------------------------------------------------------"))
	drawLine(WallLeft, WallBottom, fmt.Sprintf("--------------------------------------------------------------------------------"))
	drawLine(WallLeft, WallBottom+1, fmt.Sprintf("EXIT : ESC KEY"))
	drawLine(ball[0], ball[1], fmt.Sprintf("*"))

	for i := range shadow {
		drawLine(shadow[i][0], shadow[i][1], fmt.Sprintf(string(ipAddr[len(ipAddr)-i-1])))

	}

	for i := 0; i < Bar; i++ {
		drawLine(MeX, meY+i, fmt.Sprintf("||"))
		drawLine(EnemyX, enemyY+i, fmt.Sprintf("||"))
	}
	m.Lock()
	defer m.Unlock()
	termbox.Flush()
}

func keyEvent() {
	draw()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				clear = true
				draw()
				return
			case termbox.KeyArrowUp:
				if meY > WallTop+1 {
					meY--
				}
			case termbox.KeyArrowDown:
				if meY < WallBottom-Bar {
					meY++
				}
			default:
			}
			draw()
		default:
		}
		hitTest()
	}
}

func moveBall() {
	for {

		ball[0] += vector[0]
		ball[1] += vector[1]
		hitTest()
		draw()

		ball[0] += vector[0]
		hitTest()
		recMove()
		draw()

		if ball[1] <= WallTop+1 || ball[1] >= WallBottom-1 {
			vector[1] *= -1
		}

		if ball[0] <= WallLeft+1 || ball[0] >= WallRight-1 {
			vector[0] *= -1
			if ball[0] <= WallLeft+1 {
				score[1]++
			}
			if ball[0] >= WallRight-1 {
				score[0]++
			}
			initGame()
			draw()
			time.Sleep(time.Duration(500) * time.Millisecond)
		}

		time.Sleep(time.Duration(100-level*5) * time.Millisecond)
	}
}
func recMove() {
	shadow = append(shadow, []int{ball[0], ball[1]})
	shadow = shadow[1:]
}
func moveEnemy() {
	vec := 0
	for {

		vec = ball[1] - (enemyY + 2)

		if enemyY > WallTop+1 && vec < 0 {
			enemyY--
		}
		if enemyY < WallBottom-Bar && vec > 0 {
			enemyY++
		}
		hitTest()
		draw()

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
	if vector[0] == 1 && ball[0] > WallRight-10 {
		if (ball[0] == MeX || ball[0] == MeX-1) && ball[1] >= meY && ball[1] <= meY+Bar {
			vector[0] *= -1
			level = (level + 1) % 10
		}
	}
	if vector[0] == -1 && ball[0] < WallLeft+10 {
		if (ball[0] == EnemyX+1 || ball[0] == EnemyX+2) && ball[1] >= enemyY && ball[1] <= enemyY+Bar {
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
	m = new(sync.Mutex)
	initGame()
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	go moveBall()
	go moveEnemy()
	defer termbox.Close()

	keyEvent()
}
