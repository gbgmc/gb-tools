package fileops

import (
	"io"
	"os"

	"github.com/JakBaranowski/gb-tools/common"
)

// Copies file and it's content from a file at srcPath to a file at dstPath.
func Copy(srcPath string, dstPath string) {
	src, err := os.Open(srcPath)
	common.Must(err)
	defer src.Close()

	dst, err := os.Create(dstPath)
	common.Must(err)
	defer dst.Close()

	_, err = io.Copy(dst, src)
	common.Must(err)
}

// Touches, i.e. creates an empty file, at dstPath.
func Touch(dstPath string) {
	dst, err := os.Create(dstPath)
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
func ReadFile(file *os.File) []byte {
	fileStat, err := file.Stat()
	common.Must(err)

	fileByte := make([]byte, fileStat.Size())
	_, err = file.Read(fileByte)
	common.Must(err)

	return fileByte
}

// Checks if directory under path exist, and creates it if it does not.
func CreateDirIfDoesntExist(path string, perm os.FileMode) {
	if !DoesExist(path) {
		CreateDir(path, perm)
	}
}

// Checks if a directory or file exists under specified path
func DoesExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Creates directory at path with the specified permission.
func CreateDir(path string, perm os.FileMode) {
	err := os.Mkdir(path, perm)
	common.Must(err)
}

// Removes the file at path.
func RemoveFile(path string) {
	err := os.Remove(path)
	common.Must(err)
}

// Creates a copy of the srcFile in dstFile, and replaces all occurences of
// Will keep the file size unchanged if keepSize is True.
func ReplaceBytesNewFile(
	srcFile string,
	dstFile string,
	find string,
	replace string,
	keepSize bool,
) {
	byteFind := []byte(find)
	byteReplace := []byte(replace)

	byteSrc := OpenAndReadFile(srcFile)
	byteDst := ReplaceBytes(byteSrc, byteFind, byteReplace, keepSize)

	dst, err := os.Create(dstFile)
	common.Must(err)

	_, err = dst.Write(byteDst)
	common.Must(err)
}

// Replaces all occurences of stringFind in the specified file with stringReplace.
// Will keep the file size unchanged if keepSize is True.
func ReplaceBytesInFile(
	file string,
	find string,
	replace string,
	keepSize bool,
) {
	byteFind := []byte(find)
	byteReplace := []byte(replace)

	src, err := os.OpenFile(file, os.O_RDWR, 0644)
	common.Must(err)
	defer src.Close()

	byteSrc := ReadFile(src)
	byteDst := ReplaceBytes(byteSrc, byteFind, byteReplace, keepSize)

	_, err = src.WriteAt(byteDst, 0)
	common.Must(err)
}
