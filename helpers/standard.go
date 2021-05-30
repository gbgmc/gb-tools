package helpers

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Copies file and it's content from a file at srcPath to a file at dstPath.
func Copy(srcPath string, dstPath string) {
	fmt.Printf("Attempting to copy file from \"%s\" to \"%s\".\n", srcPath, dstPath)

	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Copied file from \"%s\" to \"%s\".\n", srcPath, dstPath)
}

// Touches, i.e. creates an empty file, at dstPath
func Touch(dstPath string) {
	fmt.Printf("Attempting to touch file \"%s\".\n", dstPath)
	dst, err := os.Create(dstPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	fmt.Printf("Successfully touched file \"%s\".\n", dstPath)
}

func OpenAndReadFile(pathToFile string) []byte {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return readFile(file)
}

func readFile(file *os.File) []byte {
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileByte := make([]byte, fileStat.Size())
	_, err = file.Read(fileByte)
	if err != nil {
		log.Fatal(err)
	}

	return fileByte
}
