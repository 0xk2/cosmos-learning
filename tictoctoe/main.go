package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type State string
type User string
type Position int8

const (
	EMPTY State = "empty"
	O     State = "o"
	X     State = "x"
)

const (
	NO_USER User = "no_user"
	O_User  User = "o"
	X_User  User = "x"
)

type Game struct {
	Id     int       `json:"id"`
	Dx     Position  `json:"dx"`
	Dy     Position  `json:"dy"`
	Data   [][]State `json:"data"`
	Next   User      `json:"nextUser"`
	Winner User      `json:"winner"`
}

func main() {
	args := os.Args[1:]
	cmd := args[0]
	switch cmd {
	case "init":
		cmd_init(args[1])
		break
	case "move":
		cmd_move(args[1], args[2], args[3], args[4])
		break
	case "print":
		cmd_display(args[1])
		break
	}
}

func cmd_init(pGameId string) {
	gameId := intFromCmd(pGameId)
	if gameId == -1 {
		return
	} else {
		initGame(gameId, 9, 9)
	}
}

func cmd_move(pGameId string, pUser string, pPosX string, pPosY string) {
	gameId := intFromCmd(pGameId)
	posX := intFromCmd(pPosX)
	posY := intFromCmd(pPosY)
	if gameId == -1 || pUser == "" || posX == -1 || posY == -1 {
		return
	}
	game, err := loadGameData(gameId)
	if err != nil {
		return
	}
	move(User(pUser), Position(posX), Position(posY), game)
}

func cmd_display(pGameId string) {
	gameId := intFromCmd(pGameId)
	if gameId == -1 {
		return
	}
	display(gameId)
}

func intFromCmd(param string) int {
	result, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("Error Int is incorrect")
		return -1
	}
	return result
}

func chooseMove(u User) State {
	var result State
	switch u {
	case O_User:
		result = O
	case X_User:
		result = X
	}
	return result
}

// declare the board
func initGame(gameId int, dx Position, dy Position) {
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

func storeGameData(g Game, gameId int) error {
	content, err := json.Marshal(g)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(strconv.Itoa(gameId)+".json", content, 0644)
	if err != nil {
		fmt.Println("Error in writing Game Data")
	}
	return err
}

func loadGameData(gameId int) (Game, error) {
	content, err := ioutil.ReadFile(strconv.Itoa(gameId) + ".json")
	if err != nil {
		fmt.Println("Error when opening file: ", err)
	}
	// Now let's unmarshall the data into `game`
	var game Game
	err = json.Unmarshal(content, &game)
	if err != nil {
		fmt.Println("Error during Unmarshal(): ", err)
	}
	return game, err
}

func move(u User, px Position, py Position, g Game) {
	isValid := validate(u, px, py, g)
	if !isValid {
		return
	}
	g.Data[py][px] = chooseMove(u)
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

func display(gameId int) {
	g, err := loadGameData(gameId)
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

func checkEndGame(u User, g Game, px Position, py Position) bool {
	s := chooseMove(u)
	// hrz
	// - to the left
	// - to the right
	hrz := 0
	for i := px; i < g.Dy; i++ {
		if g.Data[py][i] == s {
			hrz++
		} else {
			break
		}
	}
	hrz--
	for i := px; i >= 0; i-- {
		if g.Data[py][i] == s {
			hrz++
		} else {
			break
		}
	}
	fmt.Println("hrz: ", hrz)
	if hrz >= 5 {
		return true
	}
	// vrt
	// - to the left
	// - to the right
	vrt := 0
	for i := py; i < g.Dx; i++ {
		if g.Data[i][px] == s {
			vrt++
		} else {
			break
		}
	}
	vrt--
	for i := py; i >= 0; i-- {
		if g.Data[i][px] == s {
			vrt++
		} else {
			break
		}
	}
	fmt.Println("vrt: ", vrt)
	if vrt >= 5 {
		return true
	}
	// ldiag
	// - to the left
	// - to the right
	ldiag := 0
	for i := px; i < g.Dy; i++ {
		for j := py; j < g.Dx; j++ {
			if g.Data[j][i] == s {
				ldiag++
			} else {
				break
			}
		}
	}
	ldiag--
	for i := px; i >= 0; i-- {
		for j := py; j >= 0; j-- {
			if g.Data[j][i] == s {
				ldiag++
			} else {
				break
			}
		}
	}
	fmt.Println("ldiag: ", ldiag)
	if ldiag >= 5 {
		return true
	}
	// rdiag
	// - to the left
	// - to the right
	rdiag := 0
	for i := px; i < g.Dy; i++ {
		for j := py; j >= 0; j-- {
			if g.Data[j][i] == s {
				rdiag++
			} else {
				break
			}
		}
	}
	rdiag--
	for i := px; i >= 0; i-- {
		for j := py; j < g.Dx; j++ {
			if g.Data[j][i] == s {
				rdiag++
			} else {
				break
			}
		}
	}
	fmt.Println("rdiag: ", rdiag)
	if rdiag >= 5 {
		return true
	}
	return false
}
