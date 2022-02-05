package main

import (
	"fmt"
	"sync"
	"time"
)

func getCS(count int) []*Chopstick {
	chopsticks := make([]*Chopstick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = &Chopstick{}
	}

	return chopsticks
}

type Chopstick struct {
	sync.Mutex
}

func getPhilosophers(count int) []*Philosopher {
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &Philosopher{id: i, eatingEnded: make(chan struct{})}
	}

	return philosophers
}

func setTable(philophers []*Philosopher, chopSticks []*Chopstick, count int) {
	for i, philosopher := range philophers {
		philosopher.leftCS = chopSticks[i]
		philosopher.rightCS = chopSticks[(i+1)%count]
	}
}

type Philosopher struct {
	id              int
	leftCS, rightCS *Chopstick
	eatingEnded     chan struct{}
}

func (p *Philosopher) lockChopsticks() {
	p.leftCS.Lock()
	p.rightCS.Lock()
}

func (p *Philosopher) unlockChopsticks() {
	p.leftCS.Unlock()
	p.rightCS.Unlock()
}

func (p *Philosopher) eat() {
	p.lockChopsticks()

	fmt.Println("starting to eat ", p.id)
	time.Sleep(100)
	fmt.Println("finishing eating ", p.id)

	p.unlockChopsticks()

	p.eatingEnded <- struct{}{}
}

type Host struct {
	servingAvailable chan struct{}
}

func (h *Host) acceptRequest(philosopher *Philosopher, requestID int, wg *sync.WaitGroup) {
	go func() {
		<-h.servingAvailable
		philosopher.eat()
	}()

	go func() {
		<-philosopher.eatingEnded
		h.servingAvailable <- struct{}{}
		wg.Done()
	}()
}

func main() {
	ph := getPhilosophers(5)
	cs := getCS(5)
	setTable(ph, cs, 5)

	host := &Host{servingAvailable: make(chan struct{}, 2)}

	host.servingAvailable <- struct{}{}
	host.servingAvailable <- struct{}{}

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		for _, ph := range ph {
			wg.Add(1)
			go host.acceptRequest(ph, ph.id*3+i, &wg)
		}
	}

	wg.Wait()
	fmt.Println("all done eating")
}
