package main

type spaceImageLayer struct {
	rawData     []byte
	distibution map[byte]int
}

func newLayer(layerData []byte) *spaceImageLayer {
	dataSize := len(layerData)
	l := &spaceImageLayer{rawData: make([]byte, dataSize), distibution: map[byte]int{}}
	for i, d := range layerData {
		pixel := d - '0' // Convert from string ("n") to byte (n), where n 0..9
		l.rawData[i] = pixel
		if c, ok := l.distibution[pixel]; ok {
			l.distibution[pixel] = c + 1
		} else {
			l.distibution[pixel] = 1
		}
	}
	return l
}
