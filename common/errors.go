package common

import "log"

// Checks if error was reported and logs it as fatal if it was.
func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
