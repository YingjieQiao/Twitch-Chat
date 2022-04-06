package main

import (
	"fmt"
	"sync"
)

type VectorClock struct {
	VClock    map[int]int // vector clock type
	machineID int
	mu        sync.RWMutex
}

func CreateVectorClock(machineID int) *VectorClock {
	return &VectorClock{
		VClock:    map[int]int{machineID: 0},
		machineID: machineID,
		mu:        sync.RWMutex{},
	}
}

func (selfVectorClock *VectorClock) ToString() string {
	if selfVectorClock == nil {
		return "<nil>"
	}
	return fmt.Sprintf("&{%v %v}", selfVectorClock.machineID, selfVectorClock.VClock)
}

func (selfVectorClock *VectorClock) MergeClock(otherClock map[int]int) bool {
	// If any potential causality violation, return true
	// Else return false
	result := false
	count := 0
	selfClock := selfVectorClock.VClock

	selfVectorClock.mu.Lock()
	for k, v := range otherClock {
		if val, ok := selfClock[k]; ok {
			// If the local clock is latter, causality violation is flagged
			if val > v {
				result = true
				// If the local clock is equal, count++
			} else if val == v {
				count++
				// Else, update the local clock and continue
			} else {
				selfClock[k] = v
			}
		} else {
			selfClock[k] = v
		}
	}
	// Release read lock
	selfVectorClock.mu.Unlock()
	// advance clock now
	selfVectorClock.Advance()
	// Check whether all elements are equal
	if count == len(selfClock) {
		result = true
	}
	return result
}

func (selfVectorClock *VectorClock) Advance() {
	selfVectorClock.mu.Lock()
	selfVectorClock.VClock[selfVectorClock.machineID]++
	selfVectorClock.mu.Unlock()
}

func (selfVectorClock *VectorClock) GetVectorClock() map[int]int {
	selfVectorClock.mu.RLock()
	c := make(map[int]int)
	for k, v := range selfVectorClock.VClock {
		c[k] = v
	}
	selfVectorClock.mu.RUnlock()
	return c
}
