package hive

type Statistic[T interface{}] struct {
	Name     string
	getValue func(Statistics) T
}

func DirectStatistic(id, name string) *Statistic[int] {
	return &Statistic[int]{
		Name: name,
		getValue: func(s Statistics) int {
			return s.getStat(id)
		},
	}
}

func RatioStatistic(name string, n, d *Statistic[int]) *Statistic[float64] {
	return &Statistic[float64]{
		Name: name,
		getValue: func(s Statistics) float64 {
			nValue := s.GetInt(n)
			dValue := s.GetInt(d)
			if dValue == 0 {
				return float64(nValue)
			}
			return float64(nValue) / float64(dValue)
		},
	}
}

var (
	StatisticExperience = DirectStatistic("xp", "XP")
	StatisticLevel      = &Statistic[int]{
		Name: "Level",
		getValue: func(s Statistics) int {
			if s.game.MaxLevel == 0 {
				return 0
			}
			xp := s.getStat("xp")
			level := 0
			currentIncrement := 0
			for xp >= 0 {
				level++
				if level >= s.game.MaxLevel {
					currentIncrement = 0
				}
				if (level % s.game.MaxLevel) <= s.game.LevelIncrementCap {
					currentIncrement += s.game.LevelIncrement
				}
				xp -= currentIncrement
			}
			return level
		},
	}
	StatisticGamesPlayed = DirectStatistic("played", "Games Played")
	StatisticWins        = DirectStatistic("victories", "Wins")
	StatisticLosses      = &Statistic[int]{
		Name: "Losses",
		getValue: func(s Statistics) int {
			return s.GetInt(StatisticGamesPlayed) - s.GetInt(StatisticWins)
		},
	}
	StatisticWinRate        = RatioStatistic("Win Rate", StatisticWins, StatisticGamesPlayed)
	StatisticKills          = DirectStatistic("kills", "Kills")
	StatisticDeaths         = DirectStatistic("deaths", "Deaths")
	StatisticKillDeathRatio = RatioStatistic("KDR", StatisticKills, StatisticDeaths)
	StatisticKillsPerGame   = RatioStatistic("Kills/Game", StatisticKills, StatisticGamesPlayed)
)

type Statistics struct {
	game *Game
	data map[string]interface{}
}

func (s Statistics) getStat(stat string) int {
	value, ok := s.data[stat].(float64)
	if !ok {
		return 0
	}
	return int(value)
}

func (s Statistics) GetInt(stat *Statistic[int]) int {
	return stat.getValue(s)
}

func (s Statistics) GetFloat(stat *Statistic[float64]) float64 {
	return stat.getValue(s)
}

func (s Statistics) leaderboardPosition() (int, bool) {
	pos, ok := s.data["human_index"].(float64)
	if !ok || pos > 100 {
		return 0, false
	}
	return int(pos), true
}
