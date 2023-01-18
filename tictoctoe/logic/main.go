package logic

import (
	"fmt"

	"github.com/0xk2/cosmos-learning/tictoctoe/conversion"
	. "github.com/0xk2/cosmos-learning/tictoctoe/datatype"
)

// declare the board
func InitGame(gameId int, dx Position, dy Position) {
	// clear the default game data
	g := Game{}
	g.Id = gameId
	g.Dx = dx
	g.Dy = dy
	g.Data = make([][]State, dx, dy)
	var i, j Position
	for i = 0; i < dx; i++ {
		g.Data[i] = make([]State, dy)
		for j = 0; j < dy; j++ {
			g.Data[i][j] = EMPTY
		}
	}
	g.Next = O_User
	g.Winner = NO_USER
	storeGameData(g, gameId)
}

func Move(u User, px Position, py Position, g Game) {
	isValid := validate(u, px, py, g)
	if !isValid {
		return
	}
	g.Data[py][px] = conversion.ChooseMove(u)
	switch u {
	case O_User:
		g.Next = X_User
		break
	case X_User:
		g.Next = O_User
		break
	}
	if checkEndGame(u, g, px, py) == true {
		g.Winner = u
		fmt.Println("Winner is ", u)
	}
	storeGameData(g, g.Id)
}

func validate(u User, px Position, py Position, g Game) bool {
	result := true
	if g.Winner != NO_USER {
		fmt.Println("Game is already finished")
		result = false
	}
	if g.Data[py][px] != EMPTY {
		fmt.Printf("There is %s at [%d,%d]\n", g.Data[py][px], px, py)
		result = false
	}
	if g.Next != u {
		fmt.Println("Invalid user make the move")
		result = false
	}
	return result
}

func Display(gameId int) {
	g, err := LoadGameData(gameId)
	if err != nil {
		return
	}
	if g.Winner != NO_USER {
		fmt.Printf("Winner is %s\n", g.Winner)
	}
	fmt.Printf("Next move is %s\n", g.Next)
	for i := 0; i < int(g.Dx); i++ {
		for j := 0; j < int(g.Dy); j++ {
			switch g.Data[i][j] {
			case EMPTY:
				fmt.Print(" - ")
				break
			case O:
				fmt.Print(" o ")
			case X:
				fmt.Print(" x ")
			}
		}
		fmt.Println("")
	}
}
