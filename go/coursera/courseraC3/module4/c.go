package main

import (
	"fmt"
	"sync"
)

const (
	numPhilosophers = 5
	numEats         = 3
	maxConcurrent   = 2
)

type Chopstick struct{ sync.Mutex }

type Host struct {
	sem chan struct{}
}

func NewHost() *Host {
	h := &Host{sem: make(chan struct{}, maxConcurrent)}
	for i := 0; i < maxConcurrent; i++ {
		h.sem <- struct{}{}
	}
	return h
}

func (h *Host) RequestPermission() { <-h.sem }
func (h *Host) Release()           { h.sem <- struct{}{} }

type Philosopher struct {
	id         int
	left, right *Chopstick
	host        *Host
}

func (p *Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numEats; i++ {
		p.host.RequestPermission()
		p.left.Lock()
		p.right.Lock()

		fmt.Printf("starting to eat %d\n", p.id)
		fmt.Printf("finishing eating %d\n", p.id)

		p.right.Unlock()
		p.left.Unlock()

		p.host.Release()
	}
}

func main() {
	host := NewHost()

	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := range chopsticks {
		chopsticks[i] = &Chopstick{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := range philosophers {
		philosophers[i] = &Philosopher{
			id:    i + 1,
			left:  chopsticks[i],
			right: chopsticks[(i+1)%numPhilosophers],
			host:  host,
		}
	}

	var wg sync.WaitGroup
	for _, p := range philosophers {
		wg.Add(1)
		go p.eat(&wg)
	}
	wg.Wait()
}