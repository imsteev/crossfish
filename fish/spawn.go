package fish

import "math/rand"

type Spawner struct {
}

// Spawn selects a fish given rarity chances.
func (s *Spawner) Spawn() Fish {
	f := rand.Float64()

	for _, rarity := range []string{UltraRare, Rare, Uncommon, Common} {
		// no dice
		if f >= Chance[rarity] {
			continue
		}

		group, ok := FishGroups[rarity]
		if !ok || len(group) == 0 {
			return PaleChub
		}

		// Choose a random fish in group
		rand.Shuffle(len(group), func(i, j int) {
			tmp := group[i]
			group[i] = group[j]
			group[j] = tmp
		})

		return group[0]
	}

	return PaleChub
}
