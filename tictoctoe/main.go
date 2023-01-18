package main

import (
	"os"

	"github.com/0xk2/cosmos-learning/tictoctoe/txn"
)

func main() {
	args := os.Args[1:]
	cmd := args[0]
	switch cmd {
	case "init":
		txn.Init(args[1])
		break
	case "move":
		txn.Move(args[1], args[2], args[3], args[4])
		break
	case "print":
		txn.Display(args[1])
		break
	}
}
