package main

import (
	"context"
	"fmt"
	"time"
)

type PersonScheduler struct {
	ctx context.Context
}

func (p *PersonScheduler) SchedulePerson() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case <-time.After(time.Second):
			fmt.Println("scheduling person...")
		}
	}
}

type BalanceScheduler struct {
	ctx context.Context
}

func (b *BalanceScheduler) ScheduleBalance() {
	for {
		select {
		case <-b.ctx.Done():
			return
		case <-time.After(time.Second):
			fmt.Println("scheduling balance...")
		}
	}
}

type Scheduler struct {
	*PersonScheduler
	*BalanceScheduler
}

func (s Scheduler) StartScheduling() {
	go s.PersonScheduler.SchedulePerson()
	go s.BalanceScheduler.ScheduleBalance()
}

func NewScheduler(ctx context.Context) *Scheduler {
	return &Scheduler{PersonScheduler: &PersonScheduler{ctx}, BalanceScheduler: &BalanceScheduler{ctx}}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s := NewScheduler(ctx)
	s.StartScheduling()
	<-ctx.Done()
}
