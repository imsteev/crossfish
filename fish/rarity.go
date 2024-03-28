package fish

var (
	Common    string = "common"
	Uncommon  string = "uncommon"
	Rare      string = "rare"
	UltraRare string = "ultrarare"

	chances = map[string]float64{
		Common:    0.64,
		Uncommon:  0.30,
		Rare:      0.05,
		UltraRare: 0.01,
	}
)
