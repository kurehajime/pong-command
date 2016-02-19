package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"time"
  	  "sync"
)

var (
    	m *sync.Mutex
	me_y    int   = 12
	enemy_y int   = 12
	ball    []int = []int{40, 12}
	vector  []int = []int{-1, 1}
	level   int   = 0
	score   []int = []int{0, 0}
	shadow  [][]int
	ipAddr  string
	clear   bool = false
)

const (
	WALL_LEFT   = 0
	WALL_RIGHT  = 79
	WALL_TOP    = 1
	WALL_BOTTOM = 23
	ME_X        = 76
	ENEMY_X     = 2
	BAR         = 4
)

func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i += 1 {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}

func draw() {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if clear == true {
		return
	}
	drawLine(WALL_LEFT, WALL_TOP-1, fmt.Sprintf("                                                                     %03d - %03d", score[0], score[1]))
	drawLine(WALL_LEFT, WALL_TOP, fmt.Sprintf("--------------------------------------------------------------------------------"))
	drawLine(WALL_LEFT, WALL_BOTTOM, fmt.Sprintf("--------------------------------------------------------------------------------"))
	drawLine(WALL_LEFT, WALL_BOTTOM+1, fmt.Sprintf("EXIT : ESC KEY"))
	drawLine(ball[0], ball[1], fmt.Sprintf("*"))

	for i, _ := range shadow {
		drawLine(shadow[i][0], shadow[i][1], fmt.Sprintf(string(ipAddr[len(ipAddr)-i-1])))

	}

	for i := 0; i < BAR; i++ {
		drawLine(ME_X, me_y+i, fmt.Sprintf("||"))
		drawLine(ENEMY_X, enemy_y+i, fmt.Sprintf("||"))
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
				if me_y > WALL_TOP+1 {
					me_y--
				}
			case termbox.KeyArrowDown:
				if me_y < WALL_BOTTOM-BAR {
					me_y++
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

		if ball[1] <= WALL_TOP+1 || ball[1] >= WALL_BOTTOM-1 {
			vector[1] *= -1
		}

		if ball[0] <= WALL_LEFT+1 || ball[0] >= WALL_RIGHT-1 {
			vector[0] *= -1
			if ball[0] <= WALL_LEFT+1 {
				score[1]++
			}
			if ball[0] >= WALL_RIGHT-1 {
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

		vec = ball[1] - (enemy_y + 2)

		if enemy_y > WALL_TOP+1 && vec < 0 {
			enemy_y--
		}
		if enemy_y < WALL_BOTTOM-BAR && vec > 0 {
			enemy_y++
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
	if vector[0] == 1 && ball[0] > WALL_RIGHT-10 {
		if (ball[0] == ME_X || ball[0] == ME_X-1) && ball[1] >= me_y && ball[1] <= me_y+BAR {
			vector[0] *= -1
			level = (level + 1) % 10
		}
	}
	if vector[0] == -1 && ball[0] < WALL_LEFT+10 {
		if (ball[0] == ENEMY_X+1 || ball[0] == ENEMY_X+2) && ball[1] >= enemy_y && ball[1] <= enemy_y+BAR {
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
	for i, _ := range shadow {
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
