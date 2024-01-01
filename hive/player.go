package hive

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Player struct {
	Username                string   `json:"username"`
	UsernameCc              string   `json:"username_cc"`
	Uuid                    string   `json:"uuid"`
	Rank                    string   `json:"rank"`
	Xuid                    int      `json:"xuid"`
	FirstPlayed             int      `json:"first_played"`
	DailyLoginStreak        int      `json:"daily_login_streak"`
	LongestDailyLoginStreak int      `json:"longest_daily_login_streak"`
	HubTitlesUnlocked       []string `json:"hub_titles_unlocked"`
	AvatarsUnlocked         []Avatar `json:"avatar_unlocked"`
	CostumesUnlocked        []string `json:"costume_unlocked"`
	FriendCount             int      `json:"friend_count"`
	EquippedHubTitle        string   `json:"equipped_hub_title"`
	EquippedAvatar          Avatar   `json:"equipped_avatar"`
	QuestCount              int      `json:"quest_count"`
	PaidRanks               []string `json:"paid_ranks"`
	Pets                    []string `json:"pets"`
	Mounts                  []string `json:"mounts"`
	Hats                    []string `json:"hats"`
	allTimeStatistics       map[string]Statistics
	monthlyStatistics       map[string]Statistics
	Handler                 PlayerHandler
}

func GetPlayer(player string) (*Player, error) {
	p := &Player{
		UsernameCc: player,
		Username:   player,
		Handler:    nopPlayerHandler{},
	}
	err := p.Update()
	if err != nil {
		return nil, err
	}
	return p, nil
}

const (
	AllTimeStatsUrl       = "https://api.playhive.com/v0/game/all/all/%s"
	MonthlyStatsUrl       = "https://api.playhive.com/v0/game/monthly/player/all/%s"
	AllTimeLeaderboardUrl = "https://api.playhive.com/v0/game/all/%s"
	MonthlyLeaderboardUrl = "https://api.playhive.com/v0/game/monthly/%s"
)

var ErrorInvalidPlayer = errors.New("invalid player")

func (p *Player) getStats(url string) (map[string]Statistics, error) {
	resp, err := http.Get(fmt.Sprintf(url, p.UsernameCc))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrorInvalidPlayer
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	// Main
	decoded := &struct {
		Main *Player `json:"main"`
	}{p}
	err = json.Unmarshal(data, decoded)
	if err != nil {
		return nil, err
	}

	// Game Stats
	rawStats := map[string]any{}
	err = json.Unmarshal(data, &rawStats)
	if err != nil {
		return nil, err
	}
	delete(rawStats, "main")
	stats := map[string]Statistics{}
	for gameId, rawData := range rawStats {
		data, ok := rawData.(map[string]any)
		if !ok {
			continue
		}
		stats[gameId] = Statistics{
			game: games[gameId],
			data: data,
		}
	}
	return stats, nil
}

func (p *Player) Update() error {
	allTimeStats, err := p.getStats(AllTimeStatsUrl)
	if err != nil {
		return err
	}

	oldStats := p.allTimeStatistics
	p.allTimeStatistics = allTimeStats

	monthlyStats, err := p.getStats(MonthlyStatsUrl)
	if err != nil {
		return nil
	}
	p.monthlyStatistics = monthlyStats

	var currentGame *Game = nil

	for game, newGameStats := range allTimeStats {
		oldGameStats, ok := oldStats[game]
		if !ok || newGameStats.GetInt(StatisticGamesPlayed) > oldGameStats.GetInt(StatisticGamesPlayed) {
			currentGame = games[game]
			break
		}
	}

	if currentGame != nil {
		p.Handler.HandleStatsUpdated(currentGame)
	}

	return nil
}

func (p *Player) AllTimeStatistics(game *Game) Statistics {
	return p.allTimeStatistics[game.Id]
}

func (p *Player) MonthlyStatistics(game *Game) Statistics {
	return p.monthlyStatistics[game.Id]
}

func (p *Player) AllTimeLeaderboardPosition(game *Game) (int, bool) {
	lb, err := AllTimeLeaderboard(game)
	if err != nil {
		return 0, false
	}
	return lb.PlayerPosition(p)
}

func (p *Player) MonthlyLeaderboardPosition(game *Game) (int, bool) {
	return p.MonthlyStatistics(game).leaderboardPosition()
}

type PlayerHandler interface {
	HandleStatsUpdated(currentGame *Game)
}

type nopPlayerHandler struct{}

func (nopPlayerHandler) HandleStatsUpdated(_ *Game) {}

type Avatar struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
