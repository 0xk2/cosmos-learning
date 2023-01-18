package datatype

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
