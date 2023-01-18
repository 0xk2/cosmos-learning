package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	cmd := args[0]
	switch cmd {
	case "create":
		cmd_create()
	}
}

func cmd_create() {
	fmt.Println("create")
}

func cmd_vote() {
	fmt.Println("vote")
}

// create a decision making process, no template for now
func create() {

}

// vote
func vote() {

}

// store the decision making process into a json file
func storeData() {

}

// load the decision making process from a json file
func loadData() {

}
