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

func (s *Scoreboard) writeScore(index, score int) error {
	if err := output.WriteFile("debug/score/players/"+strconv.Itoa(index)+"/score.txt", strconv.Itoa(score)); err != nil {
		return fmt.Errorf("failed to write player score: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetScore(playerIndex, score int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Players[playerIndex].Score = score
	return s.writeScore(playerIndex, score)
}

func (s *Scoreboard) writePlayerName(index int, name string) error {
	if err := output.WriteFile("debug/score/players/"+strconv.Itoa(index)+"/name.txt", name); err != nil {
		return fmt.Errorf("failed to write player name: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetPlayerName(playerIndex int, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Players[playerIndex].Name = name
	return s.writePlayerName(playerIndex, name)
}

func (s *Scoreboard) writeTeamName(playerIndex int, teamName string) error {
	if err := output.WriteFile("debug/score/players/"+strconv.Itoa(playerIndex)+"/teamName.txt", teamName); err != nil {
		return fmt.Errorf("failed to write player team name: %w", err)
	}
	return nil
}

func (s *Scoreboard) SetTeamName(playerIndex int, teamName string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Players[playerIndex].TeamName = teamName
	return s.writeTeamName(playerIndex, teamName)
}
