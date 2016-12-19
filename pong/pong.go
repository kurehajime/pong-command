//pong.go
package main

import (
	"fmt"
	"math"
	"time"

	"math/rand"

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
	Count       int
}

var (
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
	drawObj(s.Ball)
	for i := range s.Shadows {
		drawObj(s.Shadows[i])
	}
	drawObj(s.LeftLine)
	drawObj(s.RightLine)
	drawObj(s.TopLine)
	drawObj(s.BottomLine)
	drawLine(1, 0, "EXIT : ESC KEY")
	drawLine(_width-10, 0, fmt.Sprintf("%03d - %03d", s.ScoreEnemy, s.ScorePlayer))
	drawObj(s.Player)
	drawObj(s.Enemy)
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

//drawLine
func drawLine(x, y int, str string) {
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], termbox.ColorDefault, termbox.ColorDefault)
	}
}

// controller
func controller(s state, kch chan termbox.Key, tch chan bool) {
	var ballMaxTime = 9
	var ballTime = ballMaxTime
	var enemyMaxTime = 7
	var enemyTime = enemyMaxTime
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
			s = updateStatus(s)
		case <-tch: //time event
			ballTime = ballTime - 1
			if ballTime < 0 {
				ballTime = ballMaxTime - int(math.Min(float64(s.Count), 8))
				prevX := s.Ball.Point().X
				s.Ball.Next()
				s.Shadows = nextShadow(s.Shadows, s.Ball)
				nextX := s.Ball.Point().X
				s = updateStatus(s)
				if _width/2 >= s.Ball.Point().X && prevX < nextX {
					s.Ball.Move(1, 0)
					s = updateStatus(s)
				} else if _width/2 <= s.Ball.Point().X && prevX > nextX {
					s.Ball.Move(-1, 0)
					s = updateStatus(s)
				}
			}
			enemyTime = enemyTime - 1
			if enemyTime < 0 {
				enemyTime = enemyMaxTime
				s.Enemy.Move(0, enemyMove(s.Enemy, s.Ball, s.Player))
				s = updateStatus(s)
			}
			break
		default:
			break
		}
		update(s)
	}
}

//enemyMove
func enemyMove(enemy Objective, ball Objective, player Objective) int {
	if ball.Point().X <= _width/2 {
		if enemy.Point().Y+2 < ball.Point().Y {
			return 1
		} else if enemy.Point().Y+2 > ball.Point().Y {
			return -1
		}
		return 0
	}
	p := (player.Point().Y + ball.Point().Y) / 2
	if int(math.Abs(float64(enemy.Point().Y-p))) < 1 {
		return 0
	} else if enemy.Point().Y <= p {
		return 1
	} else if enemy.Point().Y >= p {
		return -1
	}
	return 0
}

//updateStatus
func updateStatus(s state) state {
	if s.Ball.Collision(s.Player) || s.Ball.Collision(s.Enemy) {
		s.Ball.Prev()
		s.Shadows = nextShadow(s.Shadows, s.Ball)
		s.Ball.Turn(VERTICAL)
		s.Ball.Next()
		s.Shadows = nextShadow(s.Shadows, s.Ball)
		s.Count++
	}
	if s.Ball.Collision(s.TopLine) || s.Ball.Collision(s.BottomLine) {
		s.Ball.Prev()
		s.Shadows = nextShadow(s.Shadows, s.Ball)
		s.Ball.Turn(HORIZONAL)
		s.Ball.Next()
		s.Shadows = nextShadow(s.Shadows, s.Ball)
	}
	if s.Ball.Collision(s.LeftLine) {
		s.Ball = inirBall()
		s.ScorePlayer++
		s.Count = 0
	}

	if s.Ball.Collision(s.RightLine) {
		s.Ball = inirBall()
		s.ScoreEnemy++
		s.Count = 0
	}
	return s
}

//initState
func initState(keyword string) state {
	s := state{}
	_width, _height = termbox.Size()
	s.TopLine = NewCollisionableObject(0, 1, _width, 1, "-")
	s.BottomLine = NewCollisionableObject(0, _height-2, _width, 1, "-")
	s.LeftLine = NewCollisionableObject(0, 0, 1, _height, " ")
	s.RightLine = NewCollisionableObject(_width-1, 0, 1, _height, " ")
	s.Player = NewCollisionableMovableObject(_width-3, _height/2-2, 2, 4, "|", 0, 0)
	s.Enemy = NewCollisionableMovableObject(1, _height/2-2, 2, 4, "|", 0, 0)
	s.Ball = inirBall()
	for i := range keyword {
		s.Shadows = append(s.Shadows, NewMovableObject(s.Ball.Point().X, s.Ball.Point().Y, s.Ball.Size().Width, s.Ball.Size().Height, string(keyword[i]), 0, 0))
	}
	return s
}

//nextShadow
func nextShadow(shadow []MovableObject, ball CollisionableMovableObject) []MovableObject {
	for i := len(shadow) - 1; i >= 0; i-- {
		if i == 0 {
			shadow[i] = NewMovableObject(ball.Point().X, ball.Point().Y, 1, 1, shadow[i].Str(), 0, 0)
		} else {
			shadow[i] = NewMovableObject(shadow[i-1].Point().X, shadow[i-1].Point().Y, 1, 1, shadow[i].Str(), 0, 0)
		}
	}
	return shadow
}

//inirBall
func inirBall() CollisionableMovableObject {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(_height / 3)
	vec := 0
	if rand.Intn(100) <= 50 {
		vec = 1
	} else {
		vec = -1
	}
	return NewCollisionableMovableObject(_width/2, _height/3+r1, 1, 1, " ", -1, vec)
}

//start
func start(keyword string) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	s := initState(keyword)

	kch := make(chan termbox.Key)
	tch := make(chan bool)
	go keyEventLoop(kch)
	go timerLoop(tch)
	controller(s, kch, tch)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	defer termbox.Close()
}
