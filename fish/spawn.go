package fish

import (
	"math/rand"
	"time"
)

var (
	defaultFish = PaleChub // 🙃

	chances = map[string]float64{
		Common:    0.64,
		Uncommon:  0.30,
		Rare:      0.05,
		UltraRare: 0.01,
	}
)

// Spawn selects a fish given rarity chances.
func Spawn() Fish {
	f := rand.Float64()

	for _, rarity := range []string{UltraRare, Rare, Uncommon, Common} {
		// no dice
		if f >= chances[rarity] {
			continue
		}

		group, ok := FishGroups[rarity]
		if !ok || len(group) == 0 {
			return defaultFish
		}

		// Choose a random fish in group
		return group[rand.Intn(len(group))]
	}

	return defaultFish
}

type Spawner struct {
	C    <-chan Fish
	tick *time.Ticker
}

// Call Stop to release resources.
func (s *Spawner) Stop() {
	s.tick.Stop()
}

func SpawnEvery(dur time.Duration) *Spawner {
	spawner := make(chan Fish)
	tick := time.NewTicker(dur)
	go func() {
		for range tick.C {
			spawner <- Spawn()
		}
	}()
	return &Spawner{
		C:    spawner,
		tick: tick,
	}
}
