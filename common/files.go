package common

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// Copies file and it's content from a file at sourcePath to a file at dstPath.
func Copy(pathToSource string, pathToDestination string) {
	src, err := os.Open(pathToSource)
	cobra.CheckErr(err)
	defer src.Close()

	dst, err := os.Create(pathToDestination)
	cobra.CheckErr(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	cobra.CheckErr(err)
}

// Opens and reads a specified file. Returns read byte array.
func OpenAndReadFile(pathToFile string) []byte {
	file, err := os.Open(pathToFile)
	cobra.CheckErr(err)
	defer file.Close()
	return ReadFile(file)
}

// Reads a provided file. Returns read byte array.
func ReadFile(file *os.File) (fileByte []byte) {
	fileStat, err := file.Stat()
	cobra.CheckErr(err)

	fileByte = make([]byte, fileStat.Size())
	_, err = file.Read(fileByte)
	cobra.CheckErr(err)

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
	cobra.CheckErr(err)
}

// Removes the file or empty directory at path. Ignores non-empty directories.
func Remove(path string) {
	err := os.Remove(path)
	var pathErr *os.PathError
	if errors.As(err, &pathErr) &&
		pathErr.Err.Error() == "The directory is not empty." {
		return
	}
	cobra.CheckErr(err)
}

// Writes body to file. If file doesn't exist it's created with permission perm
// otherwise permission doesn't change.
func WriteFile(pathToFile string, body []byte, perm os.FileMode) {
	err := os.WriteFile(pathToFile, body, perm)
	cobra.CheckErr(err)
}
