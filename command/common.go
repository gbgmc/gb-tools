package command

import (
	"errors"
	"os"
)

// Checks if argument under specified index exists and returns it if it does.
func GetArgument(index int) (string, error) {
	if len(os.Args) < index+1 {
		return "", errors.New("insufficient number of arguments given")
	}
	return os.Args[index], nil
}
