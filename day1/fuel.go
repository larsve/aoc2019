package main

func requiredFuel(mass int64) int64 {
	if rf := mass/3 - 2; rf >= 0 {
		return rf
	}
	return 0
}

func totalRequiredFuel(masses []int64) (tot int64) {
	for _, mass := range masses {
		tot += requiredFuel(mass)
	}
	return tot
}

func totalMassCompensatedFuel(masses []int64) (tot int64) {
	for _, mass := range masses {
		tot += massCompensatedFuel(requiredFuel(mass))
	}
	return tot
}

func massCompensatedFuel(mass int64) (tot int64) {
	tot = mass
	additionalFuel := requiredFuel(tot)
	for additionalFuel > 0 {
		tot += additionalFuel
		additionalFuel = requiredFuel(additionalFuel)
	}
	return tot
}
