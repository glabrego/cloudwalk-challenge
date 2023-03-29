package report

import (
	"fmt"
	"github.com/glabrego/cloudwalk-challenge/lib/parser"
)

type MatchReport struct {
	TotalKills   int            `json:"total_kills"`
	Players      []string       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means"`
}

var _meansOfDeath []string = []string{
	"MOD_UNKNOWN",
	"MOD_SHOTGUN",
	"MOD_GAUNTLET",
	"MOD_MACHINEGUN",
	"MOD_GRENADE",
	"MOD_GRENADE_SPLASH",
	"MOD_ROCKET",
	"MOD_ROCKET_SPLASH",
	"MOD_PLASMA",
	"MOD_PLASMA_SPLASH",
	"MOD_RAILGUN",
	"MOD_LIGHTNING",
	"MOD_BFG",
	"MOD_BFG_SPLASH",
	"MOD_WATER",
	"MOD_SLIME",
	"MOD_LAVA",
	"MOD_CRUSH",
	"MOD_TELEFRAG",
	"MOD_FALLING",
	"MOD_SUICIDE",
	"MOD_TARGET_LASER",
	"MOD_TRIGGER_HURT",
	"MOD_NAIL",
	"MOD_CHAINGUN",
	"MOD_PROXIMITY_MINE",
	"MOD_KAMIKAZE",
	"MOD_JUICED",
	"MOD_GRAPPLE",
}

const _worldID int = 1022

func ReportMatches(matches []parser.Match) map[string]MatchReport {
	matchesReport := map[string]MatchReport{}
	if len(matches) == 0 {
		return matchesReport
	}
	for key, value := range matches {
		players := []string{}
		for _, playersValue := range value.Players {
			playerIndex := parser.FindUserByID(value.Players, playersValue.ID)
			players = append(players, value.Players[playerIndex].Name)
		}
		if len(value.Events) == 0 {
			matchesReport[fmt.Sprintf("game_%d", key+1)] = MatchReport{
				TotalKills: 0,
				Players:    players,
				Kills:      map[string]int{},
			}
			continue
		}
		matchesReport[fmt.Sprintf("game_%d", key+1)] = MatchReport{
			TotalKills: len(value.Events),
			Players:    players,
			Kills:      map[string]int{},
		}

		report := matchesReport[fmt.Sprintf("game_%d", key+1)]
		report.KillsByMeans = map[string]int{}
		matchesReport[fmt.Sprintf("game_%d", key+1)] = report

		for _, eventValue := range value.Events {
			if val, ok := matchesReport[fmt.Sprintf("game_%d", key+1)].KillsByMeans[_meansOfDeath[eventValue.MeanOfDeath]]; ok {
				matchesReport[fmt.Sprintf("game_%d", key+1)].KillsByMeans[_meansOfDeath[eventValue.MeanOfDeath]] = val + 1
			} else {
				matchesReport[fmt.Sprintf("game_%d", key+1)].KillsByMeans[_meansOfDeath[eventValue.MeanOfDeath]] = 1
			}
			if eventValue.KillerID != _worldID {
				killerIndex := parser.FindUserByID(value.Players, eventValue.KillerID)
				if val, ok := matchesReport[fmt.Sprintf("game_%d", key+1)].Kills[value.Players[killerIndex].Name]; ok {
					matchesReport[fmt.Sprintf("game_%d", key+1)].Kills[value.Players[killerIndex].Name] = val + 1
				} else {
					matchesReport[fmt.Sprintf("game_%d", key+1)].Kills[value.Players[killerIndex].Name] = 1
				}
			}
		}
	}

	return matchesReport
}
