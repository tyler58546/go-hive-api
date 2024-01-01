package hive

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Leaderboard []Statistics

func (l Leaderboard) PlayerPosition(p *Player) (int, bool) {
	for _, entry := range l {
		username := entry.data["username"].(string)
		if username == p.UsernameCc {
			return int(entry.data["human_index"].(float64)), true
		}
	}
	return 0, false
}

func getLeaderboard(url string, game *Game) (Leaderboard, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	var rawLb []map[string]any
	err = json.Unmarshal(data, &rawLb)
	if err != nil {
		return nil, err
	}

	lb := make(Leaderboard, len(rawLb))
	for i, data := range rawLb {
		lb[i] = Statistics{game, data}
	}

	return lb, nil
}

func AllTimeLeaderboard(game *Game) (Leaderboard, error) {
	return getLeaderboard(fmt.Sprintf(AllTimeLeaderboardUrl, game.Id), game)
}

func MonthlyLeaderboard(game *Game) (Leaderboard, error) {
	return getLeaderboard(fmt.Sprintf(MonthlyLeaderboardUrl, game.Id), game)
}
