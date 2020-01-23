package main

import "fmt"

type (
	simulation struct {
		moons     []*moon
		stepCount int
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

func (s *simulation) currentState() int {
	state := 0
	for i, m := range s.moons {
		state *= i
		state += m.energy()
		state *= 100
		state += m.pos.x
		state *= 100
		state += m.pos.y
		state *= 100
		state += m.pos.z
		state *= 100
		state += m.vel.x
		state *= 100
		state += m.vel.y
		state *= 100
		state += m.vel.z
	}
	return state
}

func (s *simulation) energy() int {
	tot := 0
	for _, m := range s.moons {
		tot += m.energy()
	}
	return tot
}

func (s *simulation) loop() int {
	states := map[int]int{}
	states[s.currentState()] = 0
	for {
		s.step()
		state := s.currentState()
		if i, ok := states[state]; ok {
			return s.stepCount - i
		}
		states[state] = s.stepCount
	}
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

func part2() *simulation {
	sim := &simulation{}
	sim.addMoon("Io", 3, 2, -6)
	sim.addMoon("Europa", -13, 18, 10)
	sim.addMoon("Ganymede", -8, -1, 13)
	sim.addMoon("Callisto", 5, 10, 4)
	steps := sim.loop()
	fmt.Printf("Part 2 total steps: %v (steps between repeating patterns: %v)\n", sim.stepCount, steps)
	return sim
}
