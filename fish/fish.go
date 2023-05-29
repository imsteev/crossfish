// fish contains data types describing fish and groupings based off rarity.
package fish

type Fish struct {
	Name   string
	Rarity string
}

var (
	// Ultra rare
	Coelacanth = Fish{"Coelacanth", UltraRare}

	// Rare
	Sturgeon        = Fish{"Sturgeon", Rare}
	GoldenTrout     = Fish{"Golden Trout", Rare}
	Stringfish      = Fish{"Stringfish", Rare}
	Napoleonfish    = Fish{"Napoleanfish", Rare}
	OceanSunfish    = Fish{"Ocean Sunfish", Rare}
	HammerheadShark = Fish{"Hammerhead Shark", Rare}
	GreatWhiteShark = Fish{"Great White Shark", Rare}
	WhaleShark      = Fish{"Whale Shark", Rare}
	Oarfish         = Fish{"Oarfish", Rare}
	Barreleye       = Fish{"Barreleye", Rare}
	Suckerfish      = Fish{"Suckerfish", Rare}
	SawShark        = Fish{"Saw Shark", Rare}

	// Uncommon
	RibbonEel         = Fish{"Ribbon Eel", Uncommon}
	Tuna              = Fish{"Tuna", Uncommon}
	BlueMarlin        = Fish{"Blue Marlin", Uncommon}
	GiantTrevally     = Fish{"Giant Trevally", Uncommon}
	MahiMahi          = Fish{"Mahi Mahi", Uncommon}
	FootballFish      = Fish{"FootballFish", Uncommon}
	Koi               = Fish{"Koi", Uncommon}
	Goldfish          = Fish{"Goldfish", Uncommon}
	PopEyedGoldfish   = Fish{"Pop-Eyed Goldfish", Uncommon}
	RanchuGoldfish    = Fish{"Ranchu Goldfish", Uncommon}
	Killifish         = Fish{"Killifish", Uncommon}
	Crawfish          = Fish{"Crawfish", Uncommon}
	SoftShelledTurtle = Fish{"Soft-Shelled Turtle", Uncommon}
	SnappingTurtle    = Fish{"Snapping Turtle", Uncommon}
	Sweetfish         = Fish{"Sweetfish", Uncommon}
	CherrySalmon      = Fish{"Cherry Salmon", Uncommon}
	Char              = Fish{"Char", Uncommon}
	Salmon            = Fish{"Salmon", Uncommon}
	KingSalmon        = Fish{"King Salmon", Uncommon}
	MittenCrab        = Fish{"Mitten Crab", Uncommon}
	Piranha           = Fish{"Piranha", Uncommon}
	Arowana           = Fish{"Arowana", Uncommon}
	Dorado            = Fish{"Dorado", Uncommon}
	Gar               = Fish{"Gar", Uncommon}
	Arapaima          = Fish{"Arapaima", Uncommon}
	SaddledBichir     = Fish{"Saddled Bichir", Uncommon}
	GiantSnakehead    = Fish{"Giant Snakehead", Uncommon}
	Pike              = Fish{"Pike", Uncommon}

	// Common
	Bitterling      = Fish{"Bitterling", Common}
	PaleChub        = Fish{"Pale Chub", Common}
	CrucianCarp     = Fish{"Crucian Carp", Common}
	Dace            = Fish{"Dace", Common}
	Carp            = Fish{"Carp", Common}
	Tadpole         = Fish{"Tadpole", Common}
	Frog            = Fish{"Frog", Common}
	FreshwaterGoby  = Fish{"Freshwater Goby", Common}
	Loach           = Fish{"Loach", Common}
	Catfish         = Fish{"Catfish", Common}
	Bluegill        = Fish{"Bluegill", Common}
	YellowPerch     = Fish{"Yellow Perch", Common}
	BlackBass       = Fish{"Black Bass", Common}
	Tilapia         = Fish{"Tilapia", Common}
	PondSmelt       = Fish{"Pond Smelt", Common}
	Guppy           = Fish{"Guppy", Common}
	NibbleFish      = Fish{"Nibble Fish", Common}
	Angelfish       = Fish{"Angelfish", Common}
	Betta           = Fish{"Betta", Common}
	NeonTetra       = Fish{"Neon Tetra", Common}
	Rainbowfish     = Fish{"Rainbowfish", Common}
	SeaButterfly    = Fish{"Sea Butterfly", Common}
	SeaHorse        = Fish{"Sea Horse", Common}
	ClownFish       = Fish{"Clown Fish", Common}
	Surgeonfish     = Fish{"Surgeonfish", Common}
	ButterflyFish   = Fish{"Butterfly Fish", Common}
	ZebraTurkeyfish = Fish{"Zebra Turkeyfish", Common}
	Blowfish        = Fish{"Blowfish", Common}
	PufferFish      = Fish{"PufferFish", Common}
	Anchovy         = Fish{"Anchovy", Common}
	HorseMackerel   = Fish{"Horse Mackerel", Common}
	BarredKnifejaw  = Fish{"Barred Knifejaw", Common}
	SeaBass         = Fish{"Sea Bass", Common}
	RedSnapper      = Fish{"Red Snapper", Common}
	Dab             = Fish{"Dab", Common}
	OliveFlounder   = Fish{"Olive Founder", Common}
	Squid           = Fish{"Squid", Common}
	MorayEel        = Fish{"Moray Eel", Common}
	Ray             = Fish{"Ray", Common}

	FishGroups = map[string][]Fish{
		UltraRare: {
			Coelacanth,
		},
		Rare: {
			Sturgeon,
			GoldenTrout,
			Stringfish,
			Napoleonfish,
			OceanSunfish,
			SawShark,
			HammerheadShark,
			GreatWhiteShark,
			WhaleShark,
			Suckerfish,
			Oarfish,
			Barreleye,
		},
		Uncommon: {
			GiantSnakehead,
			Koi,
			Goldfish,
			PopEyedGoldfish,
			RanchuGoldfish,
			Killifish,
			Crawfish,
			SoftShelledTurtle,
			SnappingTurtle,
			Pike,
			Sweetfish,
			CherrySalmon,
			Char,
			Salmon,
			KingSalmon,
			MittenCrab,
			Piranha,
			Arowana,
			Dorado,
			Gar,
			Arapaima,
			SaddledBichir,
			RibbonEel,
			Tuna,
			BlueMarlin,
			GiantTrevally,
			MahiMahi,
			FootballFish,
		},
		Common: {
			Bitterling,
			PaleChub,
			CrucianCarp,
			Dace,
			Carp,
			Tadpole,
			Frog,
			FreshwaterGoby,
			Loach,
			Catfish,
			Bluegill,
			YellowPerch,
			BlackBass,
			Tilapia,
			PondSmelt,
			Guppy,
			NibbleFish,
			Angelfish,
			Betta,
			NeonTetra,
			Rainbowfish,
			SeaButterfly,
			SeaHorse,
			ClownFish,
			Surgeonfish,
			ButterflyFish,
			ZebraTurkeyfish,
			Blowfish,
			PufferFish,
			Anchovy,
			HorseMackerel,
			BarredKnifejaw,
			SeaBass,
			RedSnapper,
			Dab,
			OliveFlounder,
			Squid,
			MorayEel,
			Ray,
		},
	}
)
