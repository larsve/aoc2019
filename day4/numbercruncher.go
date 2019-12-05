package main

type password struct {
	digits [6]int
}

func split(n int) [6]int {
	d := [6]int{}
	d[0] = n / 100000
	d[1] = n / 10000 % 10
	d[2] = n / 1000 % 10
	d[3] = n / 100 % 10
	d[4] = n / 10 % 10
	d[5] = n % 10
	return d
}

func (p *password) haveIncreasingOrEqualDigits() bool {
	return p.digits[0] <= p.digits[1] &&
		p.digits[1] <= p.digits[2] &&
		p.digits[2] <= p.digits[3] &&
		p.digits[3] <= p.digits[4] &&
		p.digits[4] <= p.digits[5]
}

func (p *password) haveAdjacentDoubleDigitsPart1() bool {
	return p.digits[0] == p.digits[1] ||
		p.digits[1] == p.digits[2] ||
		p.digits[2] == p.digits[3] ||
		p.digits[3] == p.digits[4] ||
		p.digits[4] == p.digits[5]
}

func (p *password) haveAdjacentDoubleDigitsPart2() bool {
	if !p.haveAdjacentDoubleDigitsPart1() {
		return false
	}
	val := 0
	cnt := 1
	found := false
	for i, dig := range p.digits {
		if i > 0 {
			if dig == val {
				cnt++
			} else {
				found = found || cnt == 2
				cnt = 1
			}
		}
		val = dig
	}
	found = found || cnt == 2
	return found
}

func getPossiblePasswordsPart1(start, stop int) int {
	m := 0
	for i := start; i <= stop; i++ {
		p := &password{digits: split(i)}
		if p.haveIncreasingOrEqualDigits() && p.haveAdjacentDoubleDigitsPart1() {
			m++
		}
	}
	return m
}

func getPossiblePasswordsPart2(start, stop int) int {
	m := 0
	for i := start; i <= stop; i++ {
		p := &password{digits: split(i)}
		if p.haveIncreasingOrEqualDigits() && p.haveAdjacentDoubleDigitsPart2() {
			m++
		}
	}
	return m
}
