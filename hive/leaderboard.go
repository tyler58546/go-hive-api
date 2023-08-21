package hive

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Leaderboard []map[string]interface{}

func (l Leaderboard) PlayerPosition(p *Player) (int, bool) {
	for _, entry := range l {
		username := entry["username"].(string)
		if username == p.UsernameCc {
			return int(entry["human_index"].(float64)), true
		}
	}
	return 0, false
}

func getLeaderboard(url string) (Leaderboard, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)

	var lb Leaderboard
	err = json.Unmarshal(data, &lb)
	if err != nil {
		return nil, err
	}
	return lb, nil
}

func AllTimeLeaderboard(game *Game) (Leaderboard, error) {
	return getLeaderboard(fmt.Sprintf(AllTimeLeaderboardUrl, game.Id))
}

func MonthlyLeaderboard(game *Game) (Leaderboard, error) {
	return getLeaderboard(fmt.Sprintf(MonthlyLeaderboardUrl, game.Id))
}
