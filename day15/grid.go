package main

import (
	"container/list"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type (
	branchCounter struct {
		c int32
	}
	branchInfo struct {
		length int
		pos    pos
		dir    direction
	}
)

func (b *branchCounter) add(v int32) int32 {
	return atomic.AddInt32(&b.c, v)
}

func (b *branchCounter) inc() int32 {
	return atomic.AddInt32(&b.c, 1)
}

func (b *branchCounter) dec() int32 {
	return atomic.AddInt32(&b.c, -1)
}

func findLongestPath(grid [][]rune, startPos pos) int {
	dl := getPossibleDirections(grid, startPos, 0)
	if len(dl) == 0 {
		return 0
	}

	branches := make(chan branchInfo)
	eol := make(chan int)
	bc := &branchCounter{}
	wg := &sync.WaitGroup{}

	// Start branch walkers..
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(id int) {
			for bi := range branches {
				processBranch(id, bc, grid, bi, branches, eol)
			}
			wg.Done()
		}(i)
	}

	// Collect branch lenghts and close channels when there is no more branches to walk..
	var maxLen int
	wg.Add(1)
	go func() {
		for l := range eol {
			if l > maxLen {
				maxLen = l
			}
			if w := bc.dec(); w == 0 {
				fmt.Printf("all branches checked, longest path is %d...\n", maxLen)
				close(branches)
				close(eol)
			}
		}
		wg.Done()
	}()

	// Add branch/branches from starting position..
	bc.add(int32(len(dl)))
	fmt.Printf("starting with %d branch(es)\n", len(dl))
	for _, d := range dl {
		branches <- branchInfo{
			length: 1,
			pos:    d.newPos(startPos),
			dir:    d,
		}
	}

	wg.Wait() // Wait for all GO routines to terminate before we return...
	fmt.Println("all GO routines have terminated..")
	return maxLen
}

func getPossibleDirections(grid [][]rune, pos pos, skipDir direction) []direction {
	var res []direction
	for d := direction(1); d <= direction(4); d++ {
		if p := d.newPos(pos); grid[p.y][p.x] != wall && d != skipDir {
			res = append(res, d)
		}
	}
	return res
}

func processBranch(id int, bc *branchCounter, grid [][]rune, bi branchInfo, branches chan<- branchInfo, out chan<- int) {
	l := bi.length
	p := bi.pos
	d := bi.dir
	localCache := list.New()
	for {
		dl := getPossibleDirections(grid, p, d.opposite())
		dll := len(dl)
		if dll == 0 {
			// Nowhere to go, send walking distance..
			out <- l

			// Any branches in the local cache..
			if localCache.Len() == 0 {
				return
			}

			// Try to see if we can give away some branches to another GO routines (if we have more than one)...
			b := localCache.Len() > 1
			for b {
				bi, e := getBranchInfo(localCache, false)
				select {
				case branches <- bi:
					localCache.Remove(e)
					b = localCache.Len() > 1
				default:
					b = false
				}

			}

			// Take first branch from the local cache, and restart walking from there..
			bi, _ := getBranchInfo(localCache, true)
			l = bi.length
			p = bi.pos
			d = bi.dir
			continue
		}

		// If trail is branching out, try to give branches away, or cache it locally if no one is able to take it off us right now.
		// Always keep the first branch for ourselves..
		if len(dl) > 1 {
			for i := 1; i < dll; i++ {
				bc.inc()
				bi = branchInfo{
					length: l + 1,
					pos:    dl[i].newPos(p),
					dir:    dl[i],
				}
				select {
				case branches <- bi:
				default:
					localCache.PushBack(bi)
				}
			}
		}
		d = dl[0]
		p = d.newPos(p)
		l++
	}
}

func getBranchInfo(cache *list.List, remove bool) (branchInfo, *list.Element) {
	e := cache.Front()
	if remove {
		cache.Remove(e)
	}
	bi, ok := e.Value.(branchInfo)
	if !ok {
		panic(e.Value)
	}
	return bi, e
}
