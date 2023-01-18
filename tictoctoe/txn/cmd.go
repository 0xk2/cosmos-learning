package txn

import (
	"github.com/0xk2/cosmos-learning/tictoctoe/conversion"
	. "github.com/0xk2/cosmos-learning/tictoctoe/datatype"
	"github.com/0xk2/cosmos-learning/tictoctoe/logic"
)

func Init(pGameId string) {
	gameId := conversion.IntFromStr(pGameId)
	if gameId == -1 {
		return
	} else {
		logic.InitGame(gameId, 9, 9)
	}
}

func Move(pGameId string, pUser string, pPosX string, pPosY string) {
	gameId := conversion.IntFromStr(pGameId)
	posX := conversion.IntFromStr(pPosX)
	posY := conversion.IntFromStr(pPosY)
	if gameId == -1 || pUser == "" || posX == -1 || posY == -1 {
		return
	}
	game, err := logic.LoadGameData(gameId)
	if err != nil {
		return
	}
	logic.Move(User(pUser), Position(posX), Position(posY), game)
}

func Display(pGameId string) {
	gameId := conversion.IntFromStr(pGameId)
	if gameId == -1 {
		return
	}
	logic.Display(gameId)
}
