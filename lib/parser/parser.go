package parser

import (
	"errors"
	"regexp"
	"strconv"
)

func ParseLine(gameID int, slc *[]Match, line string) error {
	regex, error := regexp.Compile(`\d+:\d+ (\w+): (.*)`)
	if error != nil {
		return errors.New("error on Parse Line")
	}
	subStrings := regex.FindStringSubmatch(line)
	switch subStrings[1] {
	case "InitGame":
		*slc = append((*slc), Match{
			Players: []Player{},
			Events:  []Kill{},
		})
	case "ClientConnect":
		playerID, _ := strconv.Atoi(subStrings[2])
		(*slc)[gameID].Players = append((*slc)[gameID].Players, Player{
			ID:   playerID,
			Name: "",
		})
	case "ClientUserinfoChanged":
		regex, error := regexp.Compile(`^(\d+) n\\(\w+).*$`)
		if error != nil {
			return errors.New("error on Parse Line")
		}
		playerInfo := regex.FindStringSubmatch(subStrings[2])
		userID, _ := strconv.Atoi(playerInfo[1])
		userIndex := FindUserByID((*slc)[gameID].Players, userID)
		(*slc)[gameID].Players[userIndex].Name = playerInfo[2]
	case "Kill":
		regex, error := regexp.Compile(`^(\d+) (\d+) (\d+): .*$`)
		if error != nil {
			return errors.New("error on Parse Line")
		}
		playerInfo := regex.FindStringSubmatch(subStrings[2])
		killerID, _ := strconv.Atoi(playerInfo[1])
		victimID, _ := strconv.Atoi(playerInfo[2])
		meanOfDeath, _ := strconv.Atoi(playerInfo[3])
		(*slc)[gameID].Events = append((*slc)[gameID].Events, Kill{
			KillerID:    killerID,
			VictimID:    victimID,
			MeanOfDeath: meanOfDeath,
		})
	default:
		return nil
	}
	return nil
}

func FindUserByID(players []Player, id int) int {
	index := -1
	for key, value := range players {
		if value.ID == id {
			index = key
			break
		}
	}
	return index
}

type Player struct {
	ID   int
	Name string
}

type Kill struct {
	KillerID    int
	VictimID    int
	MeanOfDeath int
}

type Match struct {
	Players []Player
	Events  []Kill
}
