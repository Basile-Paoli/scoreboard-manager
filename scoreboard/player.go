package scoreboard

import (
	"fmt"
	"scoreboard-manager/output"
	"strconv"
)

type Player struct {
	Name     string
	TeamName string
	Score    int
}

func (s *Scoreboard) SetScore(playerIndex, score int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if playerIndex < 0 || playerIndex >= len(s.Players) {
		return fmt.Errorf("player index out of range")
	}
	s.Players[playerIndex].Score = score

	return s.writeScore(playerIndex, score)
}

func (s *Scoreboard) writeScore(index, score int) error {
	err := output.WriteFile("debug/score/players/"+strconv.Itoa(index)+"/score.txt", strconv.Itoa(score))
	if err != nil {
		return fmt.Errorf("failed to write player score: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetPlayerName(playerIndex int, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if playerIndex < 0 || playerIndex >= len(s.Players) {
		return fmt.Errorf("player index out of range")
	}
	s.Players[playerIndex].Name = name

	return s.writePlayerName(playerIndex, name)
}

func (s *Scoreboard) writePlayerName(index int, name string) error {
	err := output.WriteFile("debug/score/players/"+strconv.Itoa(index)+"/name.txt", name)
	if err != nil {
		return fmt.Errorf("failed to write player name: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetTeamName(playerIndex int, teamName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if playerIndex < 0 || playerIndex >= len(s.Players) {
		return fmt.Errorf("player index out of range")
	}
	s.Players[playerIndex].TeamName = teamName

	return s.writeTeamName(playerIndex, teamName)
}

func (s *Scoreboard) writeTeamName(playerIndex int, teamName string) error {
	err := output.WriteFile("debug/score/players/"+strconv.Itoa(playerIndex)+"/teamName.txt", teamName)
	if err != nil {
		return fmt.Errorf("failed to write player team name: %w", err)
	}
	return nil
}
