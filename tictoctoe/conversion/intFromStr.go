package conversion

import (
	"fmt"
	"strconv"
)

func IntFromStr(param string) int {
	result, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("Error Int is incorrect")
		return -1
	}
	return result
}
