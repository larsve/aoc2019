package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type object struct {
	name     string
	level    int
	parent   string
	children []string
}

type orbitMap struct {
	comName            string
	objects            map[string]*object
	orbitCountChecksum int
	pcom               map[string]string
}

func (o *orbitMap) addMapLine(line string) {
	objPair := strings.Split(strings.Trim(line, " \n"), ")")
	if len(objPair) == 2 {
		parent := objPair[0]
		child := objPair[1]
		obj, ok := o.objects[parent]
		if ok {
			obj.children = append(obj.children, child)
		} else {
			o.objects[parent] = &object{name: parent, children: []string{child}}
			o.pcom[parent] = child
		}

		obj, ok = o.objects[child]
		if ok {
			obj.parent = parent
			delete(o.pcom, child)
		} else {
			o.objects[child] = &object{name: child, parent: parent, children: []string{}}
		}
	}
}

func (o *orbitMap) readMap(s io.Reader) error {
	b := bufio.NewReader(s)
	for {
		l, e := b.ReadString('\n')
		o.addMapLine(l)
		if e != nil {
			if e == io.EOF {
				return nil
			}
			return e
		}
	}
}

func (o *orbitMap) validateCom() error {
	com := make([]string, len(o.pcom))
	i := 0
	for pc := range o.pcom {
		com[i] = pc
		i++
	}
	if len(com) != 1 {
		return fmt.Errorf("Faild to read map, must have exactly one COM. COM: %v", com)
	}
	o.comName = com[0]

	return nil
}

func (o *orbitMap) setDistance(name string, distance int) {
	obj, ok := o.objects[name]
	if !ok {
		return
	}
	obj.level = distance
	for _, c := range obj.children {
		o.setDistance(c, distance+1)
	}
}

func (o *orbitMap) calcChecksum() {
	sum := 0
	for _, l := range o.objects {
		sum += l.level
	}
	o.orbitCountChecksum = sum
}

func (o *orbitMap) getDistance(firstObject, secondObject string) int {
	o1, o1ok := o.objects[firstObject]
	o2, o2ok := o.objects[secondObject]
	if !o1ok || !o2ok {
		return -1
	}
	cd := 0
	o1d := o1.level - 1
	o2d := o2.level - 1
	// Search for first common branch..
	for {
		if o1.level > o2.level {
			o1 = o.objects[o1.parent]
		} else if o1.level < o2.level {
			o2 = o.objects[o2.parent]
		} else {
			o1 = o.objects[o1.parent]
			o2 = o.objects[o2.parent]
		}
		if o1 == o2 {
			cd = o1.level
			break
		}
	}
	return o1d - cd + o2d - cd
}
func newMap(mapSource io.Reader) (o *orbitMap, err error) {
	o = &orbitMap{objects: map[string]*object{}, pcom: map[string]string{}}
	if err = o.readMap(mapSource); err != nil {
		return
	}
	if err = o.validateCom(); err != nil {
		return
	}
	o.setDistance(o.comName, 0)
	o.calcChecksum()
	return
}
