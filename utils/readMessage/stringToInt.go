package readMessage

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) (res int, err error) {
	res, err = strconv.Atoi(s)
	fmt.Println("aid***:", s)
	if err != nil {
		fmt.Println("StringToInt() utils.readMessage.strconv.Atoi err=", err)
		return res, err
	}
	return res, nil
}
