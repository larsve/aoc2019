package main

type direction int

const (
	north direction = 1 + iota
	south
	west
	east
)

func (d direction) breadCrumb() rune {
	return [...]rune{'▲', '▼', '◄', '►'}[d-1]
}

func (d direction) String() string {
	return [...]string{"North", "South", "West", "East"}[d-1]
}

func (d direction) newPos(p pos) pos {
	switch d {
	case north:
		return pos{x: p.x, y: p.y - 1}
	case south:
		return pos{x: p.x, y: p.y + 1}
	case west:
		return pos{x: p.x - 1, y: p.y}
	case east:
		return pos{x: p.x + 1, y: p.y}
	default:
		return p
	}
}

func (d direction) opposite() direction {
	switch d {
	case north:
		return south
	case east:
		return west
	case south:
		return north
	case west:
		return east
	}
	return d
}

func (d direction) turnRight() direction {
	switch d {
	case north:
		return east
	case east:
		return south
	case south:
		return west
	case west:
		return north
	}
	return d
}

func (d direction) turnLeft() direction {
	switch d {
	case north:
		return west
	case west:
		return south
	case south:
		return east
	case east:
		return north
	}
	return d
}
