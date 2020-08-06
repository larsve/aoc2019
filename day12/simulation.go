package main

import "fmt"

type (
	simulation struct {
		moons     []*moon
		stepCount int
	}
	state struct {
		pos int
		vel int
	}
)

func (s *simulation) addMoon(name string, x, y, z int) {
	s.moons = append(s.moons, &moon{name: name, pos: vector{x, y, z}, vel: vector{0, 0, 0}})
}

func (s *simulation) applyGravity() {
	for _, m1 := range s.moons {
		for _, m2 := range s.moons {
			if m1 == m2 {
				continue
			}
			m1.calcVelosity(*m2)
		}
	}
}

func (s *simulation) applyVelocity() {
	for _, m1 := range s.moons {
		m1.applyVelocity()
	}
}

func (s *simulation) currentState() (x, y, z [4]state) {
	for i, m := range s.moons {
		x[i].pos = m.pos.x
		x[i].vel = m.vel.x
		y[i].pos = m.pos.y
		y[i].vel = m.vel.y
		z[i].pos = m.pos.z
		z[i].vel = m.vel.z
	}
	return
}

func (s *simulation) energy() int {
	tot := 0
	for _, m := range s.moons {
		tot += m.energy()
	}
	return tot
}

func (s *simulation) loop() int {
	xStates := map[[4]state]int{}
	yStates := map[[4]state]int{}
	zStates := map[[4]state]int{}
	xs, ys, zs := s.currentState()
	i := 0
	for {
		// Save current states
		xStates[xs] = i
		yStates[ys] = i
		zStates[zs] = i

		// Step
		s.step()
		i++
		xs, ys, zs = s.currentState()

		// Check if current states have been hit before
		xi, xok := xStates[xs]
		yi, yok := yStates[ys]
		zi, zok := zStates[zs]
		if xok && yok && zok {
			return lcm(i-xi, i-yi, i-zi)
		}
	}
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func (s *simulation) step() {
	s.applyGravity()
	s.applyVelocity()
	s.stepCount++
}

func (s *simulation) steps(count int) {
	for i := 0; i < count; i++ {
		s.step()
	}
}

func part1() *simulation {
	sim := &simulation{}
	sim.addMoon("Io", 3, 2, -6)
	sim.addMoon("Europa", -13, 18, 10)
	sim.addMoon("Ganymede", -8, -1, 13)
	sim.addMoon("Callisto", 5, 10, 4)
	sim.steps(1000)
	fmt.Println("Part 1 total energy:", sim.energy())
	return sim
}

func part2() (*simulation, int) {
	sim := &simulation{}
	sim.addMoon("Io", 3, 2, -6)
	sim.addMoon("Europa", -13, 18, 10)
	sim.addMoon("Ganymede", -8, -1, 13)
	sim.addMoon("Callisto", 5, 10, 4)
	steps := sim.loop()
	fmt.Printf("Part 2 total steps: %v (simulation steps: %v)\n", steps, sim.stepCount)
	return sim, steps
}
