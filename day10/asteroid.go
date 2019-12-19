package main

type asteroid struct {
	mapA      int     // Linear address for this asteroid on the map
	mapX      int     // X coordinate for this asteroid on the map
	mapY      int     // Y coordinate for this asteroid on the map
	distance  int     // (Manhattan) Distance from new station
	angle     float64 // Angle from the monitoring station
	destroyed bool    // Keeps track if this asteroid is destroyed or not
}

func newAsteroid(a, x, y int) *asteroid {
	return &asteroid{mapA: a, mapX: x, mapY: y}
}
