package common

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Checks if argument under specified index exists and returns it if it does.
func GetArgument(index int) (string, error) {
	if len(os.Args) < index+1 {
		return "", errors.New("insufficient number of arguments given")
	}
	return os.Args[index], nil
}

// Parses current date time seconds and returns a YYYYMDHMS formatted string.
func GetDateTimeString() string {
	t := time.Now()
	return fmt.Sprintf(
		"%d%d%d%d%d%d",
		t.Year(),
		int(t.Month()),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}
