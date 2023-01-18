package logic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	. "github.com/0xk2/cosmos-learning/tictoctoe/datatype"
)

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

func LoadGameData(gameId int) (Game, error) {
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
