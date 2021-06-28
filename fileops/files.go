package fileops

import (
	"io"
	"os"

	"github.com/JakBaranowski/gb-tools/common"
)

// Copies file and it's content from a file at sourcePath to a file at dstPath.
func Copy(pathToSource string, pathToDestination string) {
	src, err := os.Open(pathToSource)
	common.Must(err)
	defer src.Close()

	dst, err := os.Create(pathToDestination)
	common.Must(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	common.Must(err)
}

// Touches, i.e. creates an empty file, at destinationPath.
func Touch(pathToDestination string) {
	dst, err := os.Create(pathToDestination)
	common.Must(err)
	defer dst.Close()
}

// Opens and reads a specified file. Returns read byte array.
func OpenAndReadFile(pathToFile string) []byte {
	file, err := os.Open(pathToFile)
	common.Must(err)
	defer file.Close()
	return ReadFile(file)
}

// Reads a provided file. Returns read byte array.
func ReadFile(file *os.File) (fileByte []byte) {
	fileStat, err := file.Stat()
	common.Must(err)
	fileByte = make([]byte, fileStat.Size())
	_, err = file.Read(fileByte)
	common.Must(err)
	return fileByte
}

// Checks if directory under pathToDir exist, and creates it if it does not.
func CreateDirIfDoesntExist(pathToDir string, perm os.FileMode) {
	if !DoesExist(pathToDir) {
		CreateDir(pathToDir, perm)
	}
}

// Checks if a directory or file exists under specified pathToCheck.
func DoesExist(pathToCheck string) bool {
	_, err := os.Stat(pathToCheck)
	return !os.IsNotExist(err)
}

// Creates directory at pathToDir with the specified permission.
func CreateDir(pathToDir string, perm os.FileMode) {
	err := os.Mkdir(pathToDir, perm)
	common.Must(err)
}

// Removes the file at pathToFile.
func RemoveFile(pathToFile string) {
	err := os.Remove(pathToFile)
	common.Must(err)
}
