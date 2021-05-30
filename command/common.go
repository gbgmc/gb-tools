package command

import (
	"errors"
	"os"
)

func CheckForArgument(index int) (string, error) {
	if len(os.Args) < index+1 {
		return "", errors.New("insufficient number of arguments given")
	}
	return os.Args[index], nil
}
