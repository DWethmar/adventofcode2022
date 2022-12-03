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

//X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
type PlayerStrategy string

const (
	PlayerStrategyLose PlayerStrategy = "X"
	PlayerStrategyTie  PlayerStrategy = "Y"
	PlayerStrategyWin  PlayerStrategy = "Z"
)

func MapPlayerStrategy(s string) PlayerStrategy {
	switch s {
	case "X":
		return PlayerStrategyLose
	case "Y":
		return PlayerStrategyTie
	case "Z":
		return PlayerStrategyWin
	default:
		panic("invalid player move")
	}
}

func SelectPlayerMove(opponentMove OpponentMove, strategy PlayerStrategy) PlayerMove {
	switch opponentMove {
	case OpponentMoveRock:
		switch strategy {
		case PlayerStrategyLose:
			return PlayerMoveScissors
		case PlayerStrategyTie:
			return PlayerMoveRock
		case PlayerStrategyWin:
			return PlayerMovePaper
		}
	case OpponentMovePaper:
		switch strategy {
		case PlayerStrategyLose:
			return PlayerMoveRock
		case PlayerStrategyTie:
			return PlayerMovePaper
		case PlayerStrategyWin:
			return PlayerMoveScissors
		}
	case OpponentMoveScissors:
		switch strategy {
		case PlayerStrategyLose:
			return PlayerMovePaper
		case PlayerStrategyTie:
			return PlayerMoveScissors
		case PlayerStrategyWin:
			return PlayerMoveRock
		}
	}
	panic("invalid round")
}

type PlayerMove string

const (
	PlayerMoveRock     PlayerMove = "X"
	PlayerMovePaper    PlayerMove = "Y"
	PlayerMoveScissors PlayerMove = "Z"
)

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

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
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
