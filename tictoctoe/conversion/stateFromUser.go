package conversion

import . "github.com/0xk2/cosmos-learning/tictoctoe/datatype"

func ChooseMove(u User) State {
	var result State
	switch u {
	case O_User:
		result = O
	case X_User:
		result = X
	}
	return result
}
