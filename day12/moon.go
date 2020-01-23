package main

type (
	vector struct {
		x, y, z int
	}
	moon struct {
		name string
		pos  vector
		vel  vector
	}
)

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}

func (m *moon) calcVelosity(other moon) {
	newVelocity := func(v, s, o int) int {
		if s < o {
			return v + 1
		} else if s > o {
			return v - 1
		}
		return v
	}
	m.vel.x = newVelocity(m.vel.x, m.pos.x, other.pos.x)
	m.vel.y = newVelocity(m.vel.y, m.pos.y, other.pos.y)
	m.vel.z = newVelocity(m.vel.z, m.pos.z, other.pos.z)
}

func (m *moon) applyVelocity() {
	m.pos.x += m.vel.x
	m.pos.y += m.vel.y
	m.pos.z += m.vel.z
}

func (m *moon) energy() int {
	e := func(p vector) int {
		return abs(p.x) + abs(p.y) + abs(p.z)
	}
	pot := e(m.pos)
	kin := e(m.vel)
	return pot * kin
}
