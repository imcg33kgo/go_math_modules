package Go_Math_Modules

type Point struct {
	x int
	y int
}

func (self Point) distance(p Point) int {
	//var x int
	x := self.x - p.x
	y := self.y - p.y
	if x < 0 {
		x = 0 - x
	}

	if y < 0 {
		y = 0 - y
	}

	return (x*x + y*y)
}

func (self Point) equal(p Point) bool {
	if self.x == p.x && self.y == p.y {
		return true
	} else {
		return false
	}
}