package hive

type Game struct {
	Id                string
	Name              string
	MaxLevel          int
	LevelIncrement    int
	LevelIncrementCap int
}

var (
	DeathRun       = &Game{Id: "dr", Name: "DeathRun", MaxLevel: 75, LevelIncrement: 200, LevelIncrementCap: 42}
	TreasureWars   = &Game{Id: "wars", Name: "Treasure Wars", MaxLevel: 100, LevelIncrement: 150, LevelIncrementCap: 52}
	MurderMystery  = &Game{Id: "murder", Name: "Murder Mystery", MaxLevel: 100, LevelIncrement: 100, LevelIncrementCap: 82}
	SurvivalGames  = &Game{Id: "sg", Name: "Survival Games", MaxLevel: 30, LevelIncrement: 150, LevelIncrementCap: 30}
	SkyWars        = &Game{Id: "sky", Name: "SkyWars", MaxLevel: 75, LevelIncrement: 150, LevelIncrementCap: 52}
	CaptureTheFlag = &Game{Id: "ctf", Name: "Capture the Flag", MaxLevel: 20, LevelIncrement: 150, LevelIncrementCap: 20}
	BlockDrop      = &Game{Id: "drop", Name: "Block Drop", MaxLevel: 25, LevelIncrement: 150, LevelIncrementCap: 22}
	GroundWars     = &Game{Id: "ground", Name: "Ground Wars", MaxLevel: 20, LevelIncrement: 150, LevelIncrementCap: 20}
	JustBuild      = &Game{Id: "build", Name: "Just Build", MaxLevel: 20, LevelIncrement: 100, LevelIncrementCap: 100}
	BlockParty     = &Game{Id: "party", Name: "Block Party", MaxLevel: 25, LevelIncrement: 150, LevelIncrementCap: 25}
	TheBridge      = &Game{Id: "bridge", Name: "The Bridge"}
	Gravity        = &Game{Id: "grav", Name: "Gravity", MaxLevel: 25, LevelIncrement: 150, LevelIncrementCap: 25}
)

var games = map[string]*Game{
	"dr":     DeathRun,
	"wars":   TreasureWars,
	"murder": MurderMystery,
	"sg":     SurvivalGames,
	"sky":    SkyWars,
	"ctf":    CaptureTheFlag,
	"drop":   BlockDrop,
	"ground": GroundWars,
	"build":  JustBuild,
	"party":  BlockParty,
	"bridge": TheBridge,
	"grav":   Gravity,
}
