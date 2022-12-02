package main

type OpponentMove string

const (
	OpponentMoveRock     OpponentMove = "A"
	OpponentMovePaper    OpponentMove = "B"
	OpponentMoveScissors OpponentMove = "C"
)

func MapOpponentMove(s string) OpponentMove {
	switch s {
	case "A":
		return OpponentMoveRock
	case "B":
		return OpponentMovePaper
	case "C":
		return OpponentMoveScissors
	default:
		panic("invalid opponent move")
	}
}

type PlayerMove string

const (
	PlayerMoveRock     PlayerMove = "X"
	PlayerMovePaper    PlayerMove = "Y"
	PlayerMoveScissors PlayerMove = "Z"
)

func MapPlayerMove(s string) PlayerMove {
	switch s {
	case "X":
		return PlayerMoveRock
	case "Y":
		return PlayerMovePaper
	case "Z":
		return PlayerMoveScissors
	default:
		panic("invalid player move")
	}
}

type Round struct {
	OpponentMove OpponentMove
	PlayerMove   PlayerMove
}

type RoundOutcome int

const (
	RoundOutcomeTie RoundOutcome = iota
	RoundOutcomePlayerWin
	RoundOutcomeOpponentWin
)

func (r Round) Outcome() RoundOutcome {
	switch r.OpponentMove {
	case OpponentMoveRock:
		switch r.PlayerMove {
		case PlayerMoveRock:
			return RoundOutcomeTie
		case PlayerMovePaper:
			return RoundOutcomePlayerWin
		case PlayerMoveScissors:
			return RoundOutcomeOpponentWin
		}
	case OpponentMovePaper:
		switch r.PlayerMove {
		case PlayerMoveRock:
			return RoundOutcomeOpponentWin
		case PlayerMovePaper:
			return RoundOutcomeTie
		case PlayerMoveScissors:
			return RoundOutcomePlayerWin
		}
	case OpponentMoveScissors:
		switch r.PlayerMove {
		case PlayerMoveRock:
			return RoundOutcomePlayerWin
		case PlayerMovePaper:
			return RoundOutcomeOpponentWin
		case PlayerMoveScissors:
			return RoundOutcomeTie
		}
	}
	panic("invalid round")
}

// Score:
// 0 if you lost, 3 if the round was a draw, and 6 if you won
// and 1 for Rock, 2 for Paper, and 3 for Scissors
func (r Round) Score() int {
	score := 0
	switch r.Outcome() {
	case RoundOutcomeTie:
		score += 3
	case RoundOutcomePlayerWin:
		score += 6
	case RoundOutcomeOpponentWin:
		score += 0
	}

	switch r.PlayerMove {
	case PlayerMoveRock:
		score += 1
	case PlayerMovePaper:
		score += 2
	case PlayerMoveScissors:
		score += 3
	}

	return score
}
