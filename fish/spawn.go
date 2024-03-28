package fish

import (
	"math/rand"
	"time"
)

var (
	defaultFish = PaleChub
)

type Spawner struct {
	C    <-chan Fish
	tick *time.Ticker
}

// Stop stops the spawner and releases resources.
func (s *Spawner) Stop() {
	s.tick.Stop()
}

// SpawnEvery starts a timer that will spawn one [Fish] every [dur] intervals.
// Callers are responsible for calling [Spawner].Stop() to cleanup resources.
func SpawnEvery(dur time.Duration) *Spawner {
	spawner := make(chan Fish)
	tick := time.NewTicker(dur)
	go func() {
		for range tick.C {
			spawner <- spawn()
		}
	}()
	return &Spawner{
		C:    spawner,
		tick: tick,
	}
}

// spawn selects a fish given rarity.
func spawn() Fish {
	f := rand.Float64()

	for _, rarity := range []string{UltraRare, Rare, Uncommon, Common} {
		// no dice
		if f >= chances[rarity] {
			continue
		}

		group := FishGroups[rarity]
		if len(group) == 0 {
			continue
		}

		// Choose a random fish in group
		return group[rand.Intn(len(group))]
	}

	return defaultFish
}
