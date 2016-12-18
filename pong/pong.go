//pong.go
package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

type state struct {
	Player      CollisionableMovableObject
	Enemy       CollisionableMovableObject
	Ball        CollisionableMovableObject
	Shadows     []MovableObject
	TopLine     CollisionableObject
	BottomLine  CollisionableObject
	LeftLine    CollisionableObject
	RightLine   CollisionableObject
	ScorePlayer int
	ScoreEnemy  int
}

const (
	_temespan = 10
	_height   = 25
	_width    = 80
)

//timer event
func timerLoop(tch chan bool) {
	for {
		tch <- true
		time.Sleep(time.Duration(_temespan) * time.Millisecond)
	}
}

//key events
func keyEventLoop(kch chan termbox.Key) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			kch <- ev.Key
		default:
		}
	}
}

//draw console
func update(s state) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i := range s.Shadows {
		drawObj(s.Shadows[i])
	}
	drawObj(s.Player)
	drawObj(s.Enemy)
	drawObj(s.Ball)
	drawObj(s.TopLine)
	drawObj(s.BottomLine)
	drawObj(s.LeftLine)
	drawObj(s.RightLine)
	termbox.Flush()
}

//draw object
func drawObj(o Objective) {
	for w := 0; w < o.Size().Width; w++ {
		for h := 0; h < o.Size().Height; h++ {
			termbox.SetCell(o.Point().X+w, o.Point().Y+h,
				[]rune(o.Str())[0], termbox.ColorDefault, termbox.ColorDefault)
		}
	}
}

// controller
func controller(s state, kch chan termbox.Key, tch chan bool) {
	for {
		select {
		case key := <-kch: //key event
			switch key {
			case termbox.KeyEsc, termbox.KeyCtrlC: //game end
				return
			case termbox.KeyArrowUp:
				s.Player.Move(0, -1)
				break
			case termbox.KeyArrowDown:
				s.Player.Move(0, 1)
				break
			}
			update(s)
		case <-tch: //time event
			update(s)
			break
		default:
			break
		}
	}
}

func initState() state {
	s := state{}
	s.Player = NewCollisionableMovableObject(_width-3, _height/2, 2, 4, "|", 0, 0)
	return s
}

func start() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	s := initState()

	kch := make(chan termbox.Key)
	tch := make(chan bool)
	go keyEventLoop(kch)
	go timerLoop(tch)
	controller(s, kch, tch)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	defer termbox.Close()
}
