package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

// 游戏区域大小
const (
	Rows = 20
	Cols = 20
)

// 方向常量
const (
	Up = iota
	Right
	Down
	Left
)

// 游戏状态
const (
	StateInit = iota
	StatePlaying
	StateGameOver
)

// 游戏对象
type Game struct {
	// 游戏状态
	state int

	// 蛇的位置和方向
	snake struct {
		body      []struct{ x, y int }
		direction int
	}

	// 食物的位置
	food struct{ x, y int }

	// 游戏速度
	speed time.Duration
}

// 初始化游戏
func (g *Game) Init() {
	// 初始化蛇的位置和方向
	g.snake.body = []struct{ x, y int }{{Cols / 2, Rows / 2}}
	g.snake.direction = Right

	// 初始化食物的位置
	g.food = g.generateFood()

	// 初始化游戏状态
	g.state = StatePlaying

	// 初始化游戏速度
	g.speed = 200 * time.Millisecond

	// 初始化 termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// 监听键盘事件
	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					if g.snake.direction != Down {
						g.snake.direction = Up
					}
				case termbox.KeyArrowRight:
					if g.snake.direction != Left {
						g.snake.direction = Right
					}
				case termbox.KeyArrowDown:
					if g.snake.direction != Up {
						g.snake.direction = Down
					}
				case termbox.KeyArrowLeft:
					if g.snake.direction != Right {
						g.snake.direction = Left
					}
				case termbox.KeyEsc:
					g.state = StateGameOver
				case termbox.KeyEnter:
					if g.state == StateGameOver {
						g.Init()
					}
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}()

	// 开始游戏循环
	for g.state != StateGameOver {
		g.update()
		g.draw()
		time.Sleep(g.speed)
	}
}

// 更新游戏状态
func (g *Game) update() {
	// 如果游戏已结束，直接返回
	if g.state == StateGameOver {
		return
	}

	// 移动蛇的位置
	head := g.snake.body[0]
	switch g.snake.direction {
	case Up:
		head = struct{ x, y int }{head.x, head.y - 1}
	case Right:
		head = struct{ x, y int }{head.x + 1, head.y}
	case Down:
		head = struct{ x, y int }{head.x, head.y + 1}
	case Left:
		head = struct{ x, y int }{head.x - 1, head.y}
	}
	g.snake.body = append([]struct{ x, y int }{head}, g.snake.body...)
	if head.x == g.food.x && head.y == g.food.y {
		g.food = g.generateFood()
	} else {
		g.snake.body = g.snake.body[:len(g.snake.body)-1]
	}

	// 检查游戏是否结束
	if head.x < 0 || head.x >= Cols || head.y < 0 || head.y >= Rows {
		g.state = StateGameOver
	}
	for i := 1; i < len(g.snake.body); i++ {
		if head.x == g.snake.body[i].x && head.y == g.snake.body[i].y {
			g.state = StateGameOver
		}
	}

	// 调整游戏速度
	if len(g.snake.body) > 10 {
		g.speed = 100 * time.Millisecond
	} else if len(g.snake.body) > 5 {
		g.speed = 150 * time.Millisecond
	}
}

// 绘制游戏界面
func (g *Game) draw() {
	// 清空屏幕
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// 绘制蛇
	for _, p := range g.snake.body {
		termbox.SetCell(p.x, p.y, ' ', termbox.ColorGreen, termbox.ColorDefault)
	}

	// 绘制食物
	termbox.SetCell(g.food.x, g.food.y, ' ', termbox.ColorRed, termbox.ColorDefault)

	// 绘制游戏结束提示
	if g.state == StateGameOver {
		msg := "Game Over! Press Enter to restart or Esc to exit."
		for i, c := range msg {
			termbox.SetCell(Cols/2-len(msg)/2+i, Rows/2, c, termbox.ColorBlack, termbox.ColorDefault)
		}
	}

	// 刷新屏幕
	termbox.Flush()
}

// 生成新的食物
func (g *Game) generateFood() struct{ x, y int } {
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(Cols)
		y := rand.Intn(Rows)
		if !g.isColliding(struct{ x, y int }{x, y}) {
			return struct{ x, y int }{x, y}
		}
	}
}

// 检查是否与蛇的身体或边界相撞
func (g *Game) isColliding(p struct{ x, y int }) bool {
	if p.x < 0 || p.x >= Cols || p.y < 0 || p.y >= Rows {
		return true
	}
	for _, b := range g.snake.body {
		if p.x == b.x && p.y == b.y {
			return true
		}
	}
	return false
}

func main() {
	game := Game{}
	game.Init()
	fmt.Println("Game Over!")
}