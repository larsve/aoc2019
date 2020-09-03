package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	chemical         string
	reactionChemical struct {
		qty int
		chm chemical
	}
	reaction struct {
		input  []reactionChemical
		output reactionChemical
	}
	nanoFactory struct {
		reactions map[chemical]*reaction
		produced  map[chemical]int
		consumed  map[chemical]int
	}
)

const (
	fuel      chemical = "FUEL"
	ore       chemical = "ORE"
	oreSupply          = 1000000000000
)

var (
	errorInvalidReactionInput = errors.New("invalid reaction input")
	errorUnknownChemical      = errors.New("unknown chemical")
	reactionRegEx             = regexp.MustCompile(", | => |\n")
)

func (r *reaction) parse(line string) error {
	line = strings.TrimSpace(line)
	items := reactionRegEx.Split(line, -1)
	chemicals := []reactionChemical{}
	for _, s := range items {
		f := strings.Fields(strings.TrimSpace(s))
		if len(f) != 2 {
			return fmt.Errorf("parse: expected two tokens in %s, but got %d, %w", s, len(f), errorInvalidReactionInput)
		}
		q, err := strconv.Atoi(f[0])
		if err != nil {
			return fmt.Errorf("parse: expected integer as first value in %s, %w", s, err)
		}
		rc := reactionChemical{qty: q, chm: chemical(f[1])}
		chemicals = append(chemicals, rc)
	}
	if len(chemicals) < 2 {
		return fmt.Errorf("parse: must have at least two quantity/chemical groups per reaction, got %d from %s, %w", len(chemicals), line, errorInvalidReactionInput)
	}
	r.input = chemicals[:len(chemicals)-1]
	r.output = chemicals[len(chemicals)-1]
	return nil
}

func (n *nanoFactory) fuelForOre(oreSupply int) (int, error) {
	if err := n.need(1, fuel); err != nil {
		return 0, err
	}
	orePerFuel := n.consumed[ore]

	// Get a lower/upper starting point
	flower := oreSupply / orePerFuel
	fupper := int(float32(flower) * 1.1)
	//fmt.Printf("We should be able to get at least %d fuel\n", flower)
	for {
		n.reset()
		if err := n.need(fupper, fuel); err != nil {
			return 0, nil
		}
		o := n.consumed[ore]
		if o >= oreSupply {
			break
		}
		flower = fupper
		fupper = int(float32(flower) * 1.1)
	}

	//fmt.Printf("We can get somewhere between %d and %d fuel\n", flower, fupper)
	for {
		if fupper-flower <= 1 {
			return flower, nil
		}
		mp := (fupper-flower)/2 + flower
		n.reset()
		if err := n.need(mp, fuel); err != nil {
			return 0, nil
		}
		if n.consumed[ore] < oreSupply {
			flower = mp
		} else if n.consumed[ore] > oreSupply {
			fupper = mp
		}
	}
}

func (n *nanoFactory) inStock(c chemical) int {
	return n.produced[c] - n.consumed[c]
}

func (n *nanoFactory) need(q int, c chemical) error {
	if c == ore {
		n.consumed[c] += q
		return nil
	}

	if sq := n.inStock(c); sq > 0 {
		if sq > q {
			sq = q
		}
		q -= sq
		n.consumed[c] += sq
		if q == 0 {
			return nil
		}
	}

	r, ok := n.reactions[c]
	if !ok {
		return fmt.Errorf("need: chemical %s is unknown, %w", c, errorUnknownChemical)
	}

	qf := 1
	if q > r.output.qty {
		qf = (q + r.output.qty - 1) / r.output.qty
	}
	for _, nc := range r.input {
		if err := n.need(nc.qty*qf, nc.chm); err != nil {
			return err
		}
	}
	n.produced[c] += qf * r.output.qty
	n.consumed[c] += q
	return nil
}

func (n *nanoFactory) readReactions(rd io.Reader) error {
	r := bufio.NewReader(rd)
	for {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			reaction := &reaction{}
			if err := reaction.parse(line); err != nil {
				return fmt.Errorf("readReactions: failed to parse line [%s], %w", line, err)
			}
			n.reactions[reaction.output.chm] = reaction
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("readReactions: error while reading input, %w", err)
		}
	}
	return nil
}

func (n *nanoFactory) reset() {
	n.consumed = make(map[chemical]int)
	n.produced = make(map[chemical]int)
}

func newNanoFactory(reactions io.Reader) (*nanoFactory, error) {
	nf := &nanoFactory{
		reactions: make(map[chemical]*reaction),
		produced:  make(map[chemical]int),
		consumed:  make(map[chemical]int),
	}

	// Read and parse reaction recipes
	if err := nf.readReactions(reactions); err != nil {
		return nil, err
	}

	return nf, nil
}

func nanoFactoryFromFile(filename string) (*nanoFactory, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("nanoFactoryFromFile: %w", err)
	}
	defer f.Close()
	return newNanoFactory(f)
}

func part1(filename string) *nanoFactory {
	nf, err := nanoFactoryFromFile(filename)
	if err != nil {
		fmt.Printf("Failed to create nano factory, error: %v\n", err)
		return nil
	}
	if err = nf.need(1, "FUEL"); err != nil {
		return nil
	}
	fmt.Printf("Part1 needs %d ORE\n", nf.consumed[ore])
	return nf
}

func part2(filename string) (*nanoFactory, int) {
	nf, err := nanoFactoryFromFile(filename)
	if err != nil {
		fmt.Printf("Failed to create nano factory, error: %v\n", err)
		return nil, 0
	}
	fuel, err := nf.fuelForOre(oreSupply)
	if err != nil {
		fmt.Printf("Failed to calculate fuel amount, error: %v\n", err)
		return nil, 0
	}
	fmt.Printf("Part2 can produce %d fuel for %d ORE\n", fuel, oreSupply)
	return nf, fuel
}
