package logic

import (
	"fmt"

	"github.com/0xk2/cosmos-learning/tictoctoe/conversion"
	. "github.com/0xk2/cosmos-learning/tictoctoe/datatype"
)

func checkEndGame(u User, g Game, px Position, py Position) bool {
	s := conversion.ChooseMove(u)
	fmt.Println("Check with ", s)
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
	j := int(py)
	for i := px; i < g.Dy; i++ {
		if j < int(g.Dx) {
			if g.Data[j][i] == s {
				// fmt.Println("-- ldiag: ", j, ",", i)
				j++
				ldiag++
			} else {
				break
			}
		} else {
			break
		}
	}
	ldiag--
	j = int(py)
	for i := px; i >= 0; i-- {
		if j >= 0 {
			if g.Data[j][i] == s {
				// fmt.Println("-- ldiag: ", j, ",", i)
				j--
				ldiag++
			} else {
				break
			}
		} else {
			break
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
	j = int(py)
	for i := px; i < g.Dy; i++ {
		if j >= 0 {
			if g.Data[j][i] == s {
				// fmt.Println("-- rdiag: ", j, ",", i)
				j--
				rdiag++
			} else {
				break
			}
		} else {
			break
		}
	}
	j = int(py)
	rdiag--
	for i := px; i >= 0; i-- {
		if j < int(g.Dx) {
			if g.Data[j][i] == s {
				// fmt.Println("-- rdiag: ", j, ",", i)
				j++
				rdiag++
			} else {
				break
			}
		} else {
			break
		}
	}
	fmt.Println("rdiag: ", rdiag)
	if rdiag >= 5 {
		return true
	}
	return false
}
