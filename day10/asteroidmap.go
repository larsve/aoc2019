package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"
)

type asteroidMap struct {
	x, y              int
	width             int
	height            int
	asteroids         map[int]*asteroid
	monitoringStation *asteroid
}

func (a *asteroidMap) checkAsteroid(ast *asteroid) int {
	myX := ast.mapX
	myY := ast.mapY
	vec := map[int]int{}
	for _, o := range a.asteroids {
		if o == ast {
			continue
		}
		v := a.getVector(myX, myY, o.mapX, o.mapY)
		vec[v]++
	}
	return len(vec)
}

func (a *asteroidMap) getBestPlacement() (bestX, bestY, asteroidCount int) {
	startTime := time.Now()
	aChan := make(chan *asteroid, 10)
	rChan := make(chan [2]int, 10)
	wg := sync.WaitGroup{}
	bestAddr := -1
	bestCount := -1

	// Send asteroids too check to achan, and close channel when all asteroids are sent
	go func() {
		for _, a := range a.asteroids {
			aChan <- a
		}
		close(aChan)
	}()

	// Start workers that check asteroids
	nproc := runtime.NumCPU()
	wg.Add(nproc)
	for i := 0; i < nproc; i++ {
		go func() {
			defer wg.Done()
			for asteroid := range aChan {
				rChan <- [2]int{asteroid.mapA, a.checkAsteroid(asteroid)}
			}
		}()
	}

	// Monitor workers and close rchan when all workers are done
	go func() {
		wg.Wait()
		close(rChan)
	}()

	// Start to collect results
	for res := range rChan {
		if res[1] > bestCount {
			bestCount = res[1]
			bestAddr = res[0]
		}
	}

	bestAsteroid := a.asteroids[bestAddr]
	a.monitoringStation = bestAsteroid
	computeTime := time.Since(startTime)
	fmt.Printf("All asteroids checked.. Best count (%v) at %v X %v, Compute time: %v\n", bestCount, bestAsteroid.mapX, bestAsteroid.mapY, computeTime)
	return bestAsteroid.mapX, bestAsteroid.mapY, bestCount
}

func (a *asteroidMap) getVector(myX, myY, aX, aY int) int {
	oX := aX - myX
	oY := aY - myY
	if oX == 0 || oY == 0 {
		return a.getZeroVector(oX, oY)
	}
	amv := min(abs(oX), abs(oY))
	for i := amv; i > 1; i-- {
		if oX%i == 0 && oY%i == 0 {
			oX /= i
			oY /= i
			break
		}
	}

	return oY*a.width + oX
}

func (a *asteroidMap) getZeroVector(oX, oY int) int {
	if oX == 0 && oY == 0 {
		return 0
	} else if oX == 0 {
		if oY > 0 {
			oY = 1
		} else {
			oY = -1
		}
	} else {
		if oX > 0 {
			oX = 1
		} else {
			oX = -1
		}
	}
	return oY*a.width + oX
}

func angle(cX, cY, tX, tY int) float64 {
	// Calc radians for the two points
	r := math.Atan2(float64(tY-cY), float64(tX-cX))
	// Convert radians to degrees
	d := r * (180 / math.Pi)
	// Adjust 0 degrees to be "north"
	d += 90
	// Make sure result is always a positive value between 0-359 degrees
	if d < 0 {
		d += 360
	}
	return d
}

func (a *asteroidMap) setAnglesAndDistanceFrom(station *asteroid) {
	myX := station.mapX
	myY := station.mapY
	angles := map[int]float64{}
	for _, o := range a.asteroids {
		if o == station {
			continue
		}
		oX := o.mapX
		oY := o.mapY
		v := a.getVector(myX, myY, oX, oY)
		a, ok := angles[v]
		if !ok {
			a = angle(myX, myY, oX, oY)
			angles[v] = a
		}
		o.angle = a
		o.distance = abs(myX-oX) + abs(myY-oY)
	}
}

func (a *asteroidMap) vaporizeAsteroidsFrom(monitoringStation *asteroid) []*asteroid {
	fmt.Printf("Monitoring station placed at %vX%v\n", monitoringStation.mapX, monitoringStation.mapY)
	fmt.Printf("Asteroids on map: %v\n", len(a.asteroids))

	// Calculate angles and distances from the monitoring station
	a.setAnglesAndDistanceFrom(monitoringStation)

	// Create a slice sorted by angle and distance
	asteroids := make([]*asteroid, 0, len(a.asteroids)-1)
	for _, ast := range a.asteroids {
		if ast == monitoringStation {
			continue // Don't want to vaporize the monitoring station
		}
		ast.destroyed = false // Reset destryed in case we'll run this multiple times
		asteroids = append(asteroids, ast)
	}
	sort.Slice(asteroids, func(i, j int) bool {
		ai := asteroids[i]
		aj := asteroids[j]
		if ai.angle == aj.angle {
			return ai.distance < aj.distance
		}
		return ai.angle < aj.angle
	})

	// Create a slice for asteroid vaporize order, with the capacity to hold all asteroids
	vOrder := make([]*asteroid, 0, len(a.asteroids))

	// Fire up laser and lets vaporize some asteroids...
	toVaporize := len(asteroids)
	rev := 0
	for toVaporize > 0 {
		rev++
		cv := -1.0 // Set initial value outside the valid 0-359 degree range at the beginning of each revolution
		for _, ast := range asteroids {
			if ast.destroyed || ast.angle == cv {
				continue // We'll catch this on the next revolution
			}
			cv = ast.angle
			vOrder = append(vOrder, ast)
			ast.destroyed = true
			toVaporize--
		}
	}
	fmt.Printf("Vaporized %v asteroids in %v revolutions\n", len(vOrder), rev)
	return vOrder
}

func (a *asteroidMap) part2Answer() (res int) {
	startTime := time.Now()
	res = -1
	if a.monitoringStation == nil {
		a.getBestPlacement()
	}
	vo := a.vaporizeAsteroidsFrom(a.monitoringStation)
	if len(vo) >= 200 {
		ta := vo[199]
		res = ta.mapX*100 + ta.mapY
	}
	computeTime := time.Since(startTime)
	fmt.Printf("All asteroids vaporized in %v\nResult: %v", computeTime, res)
	return
}

func abs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func (a *asteroidMap) processBuffer(buffer []byte) {
	for _, b := range buffer {
		// Ignore anything that isn't '.' or '#'..
		if !(b == '.' || b == '#') {
			continue
		}
		if b == '#' {
			addr := a.y*a.width + a.x
			a.asteroids[addr] = newAsteroid(addr, a.x, a.y)
		}
		a.x++
		if a.x >= a.width {
			a.x = 0
			a.y++
		}
	}
}

func newMap(width, height int, source io.Reader) (*asteroidMap, error) {
	amap := &asteroidMap{width: width, height: height, asteroids: map[int]*asteroid{}}

	buf := make([]byte, 4095)
	for {
		br, e := source.Read(buf)
		amap.processBuffer(buf[:br])
		if e == io.EOF {
			break
		}
		if e != nil {
			return nil, e
		}
	}
	return amap, nil
}

func newMapFromFile(with, height int, filename string) (*asteroidMap, error) {
	f, e := os.Open(filename)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	return newMap(with, height, f)
}
