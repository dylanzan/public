package decorator

//装饰器模式

//IDRaw
type IDraw interface {
	Draw() string
}

//Square 正方形
type Square struct {
}

// Draw Draw
func (s Square) Draw() string {
	return "this is a sqare"
}

// ColorSquare 有色的正方形
type ColorSquare struct {
	Square IDraw
	color  string
}

//NewColorSquare NewColorSquare
func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{color: color, Square: square}
}

// Draw Draw
func (c ColorSquare) Draw() string {
	return c.Square.Draw() + "color is " + c.color
}
