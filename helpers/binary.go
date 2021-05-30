package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

// Creates a copy of the srcFile in dstFile, and replaces all occurences of
// stringFind in the dstFile with stringReplace.
func CreateNewBinaryFileWithReplacedString(
	srcFile string,
	dstFile string,
	findToReplace map[string]string,
) {
	byteSrc := OpenAndReadFile(srcFile)
	byteDst := replaceAllMultipleValues(byteSrc, findToReplace)

	dst, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}

	bytesWritten, err := dst.Write(byteDst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Written %d bytes to %s.", bytesWritten, dstFile)
}

// Replaces all occurences of stringFind in the specified file with stringReplace.
func ReplaceStringInBinaryFile(
	file string,
	findToReplace map[string]string,
) {
	src, err := os.OpenFile(file, os.O_RDWR, 0777)
	if err != nil {
		log.Fatal(err)
	}

	byteSrc := readFile(src)
	byteDst := replaceAllMultipleValues(byteSrc, findToReplace)

	bytesWritten, err := src.WriteAt(byteDst, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Written %d bytes to %s.", bytesWritten, file)
}

func replaceAllMultipleValues(
	byteSrc []byte,
	findToReplace map[string]string,
) []byte {
	byteDst := make([]byte, len(byteSrc))
	for key, value := range findToReplace {
		byteFind := []byte(key)
		byteReplace := []byte(value)
		byteFind, byteReplace = equalizeByteSize(byteFind, byteReplace)
		byteDst = bytes.ReplaceAll(byteSrc, byteFind, byteReplace)
	}
	return byteDst
}

func equalizeByteSize(
	a []byte,
	b []byte,
) (
	[]byte,
	[]byte,
) {
	diff := len(a) - len(b)
	if diff > 0 {
		extra := make([]byte, diff)
		b = append(b, extra...)
	} else if diff < 0 {
		extra := make([]byte, -diff)
		a = append(a, extra...)
	}
	return a, b
}
