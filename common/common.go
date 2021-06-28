package common

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Checks if argument under specified index exists and returns it if it does.
// If the argument does not exist returns an empty string and an error.
func GetRequiredArgument(index int, expected string) (string, error) {
	errorString := fmt.Sprintf("missing arguments. %s.", expected)
	if len(os.Args) < index+1 {
		return "", errors.New(errorString)
	}
	return os.Args[index], nil
}

// Checks if argument under specified index exists and returns it and true if it does.
// If the argument does not exist returns empty string and false.
func GetOptionalArgument(index int) (string, bool) {
	if len(os.Args) < index+1 {
		return "", false
	}
	return os.Args[index], true
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
