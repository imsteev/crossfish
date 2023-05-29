package fish

type Fish struct {
	Name   string
	Rarity string
}

var (
	Bitterling        = Fish{"Bitterling", Common}
	PaleChub          = Fish{"Pale Chub", Common}
	CrucianCarp       = Fish{"Crucian Carp", Common}
	Dace              = Fish{"Dace", Common}
	Carp              = Fish{"Carp", Common}
	Koi               = Fish{"Koi", Common}
	Goldfish          = Fish{"Goldfish", Common}
	PopEyedGoldfish   = Fish{"Pop-Eyed Goldfish", Common}
	RanchuGoldfish    = Fish{"Ranchu Goldfish", Common}
	Killifish         = Fish{"Killifish", Common}
	Crawfish          = Fish{"Crawfish", Common}
	SoftShelledTurtle = Fish{"Soft-Shelled Turtle", Common}
	SnappingTurtle    = Fish{"Snapping Turtle", Common}
	Tadpole           = Fish{"Tadpole", Common}
	Frog              = Fish{"Frog", Common}
	FreshwaterGoby    = Fish{"Freshwater Goby", Common}
	Loach             = Fish{"Loach", Common}
	Catfish           = Fish{"Catfish", Common}
	GiantSnakehead    = Fish{"Giant Snakehead", Common}
	Bluegill          = Fish{"Bluegill", Common}
	YellowPerch       = Fish{"Yellow Perch", Common}
	BlackBass         = Fish{"Black Bass", Common}
	Tilapia           = Fish{"Tilapia", Common}
	Pike              = Fish{"Pike", Common}
	PondSmelt         = Fish{"Pond Smelt", Common}
	Sweetfish         = Fish{"Sweetfish", Common}
	CherrySalmon      = Fish{"Cherry Salmon", Common}
	Char              = Fish{"Char", Common}
	GoldenTrout       = Fish{"Golden Trout", Common}
	Stringfish        = Fish{"Stringfish", Common}
	Salmon            = Fish{"Salmon", Common}
	KingSalmon        = Fish{"King Salmon", Common}
	MittenCrab        = Fish{"Mitten Crab", Common}
	Guppy             = Fish{"Guppy", Common}
	NibbleFish        = Fish{"Nibble Fish", Common}
	Angelfish         = Fish{"Angelfish", Common}
	Betta             = Fish{"Betta", Common}
	NeonTetra         = Fish{"Neon Tetra", Common}
	Rainbowfish       = Fish{"Rainbowfish", Common}
	Piranha           = Fish{"Piranha", Common}
	Arowana           = Fish{"Arowana", Common}
	Dorado            = Fish{"Dorado", Common}
	Gar               = Fish{"Gar", Common}
	Arapaima          = Fish{"Arapaima", Common}
	SaddledBichir     = Fish{"Saddled Bichir", Common}
	Sturgeon          = Fish{"Sturgeon", Common}
	SeaButterfly      = Fish{"Sea Butterfly", Common}
	SeaHorse          = Fish{"Sea Horse", Common}
	ClownFish         = Fish{"Clown Fish", Common}
	Surgeonfish       = Fish{"Surgeonfish", Common}
	ButterflyFish     = Fish{"Butterfly Fish", Common}
	Napoleonfish      = Fish{"Napoleanfish", Common}
	ZebraTurkeyfish   = Fish{"Zebra Turkeyfish", Common}
	Blowfish          = Fish{"Blowfish", Common}
	PufferFish        = Fish{"PufferFish", Common}
	Anchovy           = Fish{"Anchovy", Common}
	HorseMackerel     = Fish{"Horse Mackerel", Common}
	BarredKnifejaw    = Fish{"Barred Knifejaw", Common}
	SeaBass           = Fish{"Sea Bass", Common}
	RedSnapper        = Fish{"Red Snapper", Common}
	Dab               = Fish{"Dab", Common}
	OliveFlounder     = Fish{"Olive Founder", Common}
	Squid             = Fish{"Squid", Common}
	MorayEel          = Fish{"Moray Eel", Common}
	RibbonEel         = Fish{"Ribbon Eel", Common}
	Tuna              = Fish{"Tuna", Common}
	BlueMarlin        = Fish{"Blue Marlin", Common}
	GiantTrevally     = Fish{"Giant Trevally", Common}
	MahiMahi          = Fish{"Mahi Mahi", Common}
	OceanSunfish      = Fish{"Ocean Sunfish", Common}
	Ray               = Fish{"Ray", Common}
	SawShark          = Fish{"Saw Shark", Common}
	HammerheadShark   = Fish{"Hammerhead Shark", Common}
	GreatWhiteShark   = Fish{"Great White Shark", Common}
	WhaleShark        = Fish{"Whale Shark", Common}
	Suckerfish        = Fish{"Suckerfish", Common}
	FootballFish      = Fish{"FootballFish", Common}
	Oarfish           = Fish{"Oarfish", Common}
	Barreleye         = Fish{"Barreleye", Common}
	Coelacanth        = Fish{"Coelacanth", UltraRare}

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
