package fileops

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"strconv"
)

// Replaces all occurences of byteFind in byteSrc with byteReplace. If keepSize
// is true will not change the len(byteSrc)
func ReplaceBytes(
	byteSrc []byte,
	byteFind []byte,
	byteReplace []byte,
	keepSize bool,
) []byte {
	if keepSize {
		byteFind, byteReplace = equalizeByteSize(byteFind, byteReplace)
	}
	return bytes.ReplaceAll(byteSrc, byteFind, byteReplace)
}

// Replaces byteFind with byteReplace in byteSrc. If keepSize is true will not
// chance len(byteSrc). Returns the count of occurences of byteFind in original
// file and the changed byte array.
func ReplaceAndCount(
	byteSrc []byte,
	byteFind []byte,
	byteReplace []byte,
	keepSize bool,
) (int, []byte) {
	if keepSize {
		byteFind, byteReplace = equalizeByteSize(byteFind, byteReplace)
	}
	byteIndex := 0
	replaces := 0
	for {
		byteIndex = bytes.Index(byteSrc, byteFind)
		if byteIndex >= 0 && byteIndex < len(byteSrc) {
			replaces++
			byteSrc = bytes.Replace(byteSrc, byteFind, byteReplace, 1)
			byteSrc = ReplaceByteAtOffsetFromMatch(
				byteSrc,
				byteReplace,
				-4,
				byte(len(byteReplace)+1),
			)
		} else {
			break
		}
	}
	return replaces, byteSrc
}

// Finds the byteMatch in byteSrc and replaces the byte at the offset from the first
// index of the found byteMatch with replace.
func ReplaceByteAtOffsetFromMatch(
	byteSrc []byte,
	byteMatch []byte,
	offset int,
	replace byte,
) []byte {
	index := bytes.Index(byteSrc, byteMatch)
	if index < 0 {
		return byteSrc
	}
	for {
		byteSrc[index+offset] = replace
		indexPlusLen := index + len(byteMatch)
		baseIndex := bytes.Index(byteSrc[indexPlusLen:], byteMatch)
		if indexPlusLen > len(byteSrc) || baseIndex < 0 {
			break
		}
		index = indexPlusLen + baseIndex
	}
	return byteSrc
}

// Takes the provided byte arrays, finds the shorter array and adds elements to
// the shorter array to match the length of the longer array.
// This can be used to make sure that the file size will be not change when
// replacing byte array - if the replacement is shorter - and that the bytes we
// want to replace don't overlap other non zero bytes - if the find is shorter.
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

// Compares two byte arrays and prints out the result.
func CompareBytes(
	a []byte,
	b []byte,
) bool {
	if len(a) != len(b) {
		fmt.Printf("Arrays are not of equal length a[%d], b[%d].\n", len(a), len(b))
		return false
	}

	result := true
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Printf(
				"Index %s a: str %c, int %d; b: str %c, int %d;\n",
				strconv.FormatInt(int64(i), 16),
				a[i],
				a[i],
				b[i],
				b[i],
			)
			result = false
		}
	}
	return result
}

// Parses a multiple byte long byte sized. They are written in right to left
// format.
func GetSize(byteSize []byte) (intSize int) {
	sum := 0
	for i := range byteSize {
		sum += int(byteSize[i]) * int(math.Pow(256, float64(i)))
	}
	return sum
}

// Parses a multiple byte long byte sized. They are written in right to left
// format.
func SetSize(intSize int) (byteSize []byte) {
	byteSize = make([]byte, 4)
	rest := 0
	for i := 3; i >= 0; i-- {
		denominator := int(math.Pow(256, float64(i)))
		result := intSize / denominator
		log.Printf("Result: %d", result)
		byteSize[3] = byte(result)
		intSize = intSize % denominator
		log.Printf("Rest: %d", rest)
	}
	log.Printf("%d", byteSize)
	return
}

// Returns the count of occurences of byteFind in byteSrc.
func CountOccurences(byteSrc []byte, byteFind []byte) (count int) {
	count = 0
	byteIndex := 0
	for byteIndex < len(byteSrc) {
		baseIndex := bytes.Index(byteSrc[byteIndex:], byteFind)
		if baseIndex >= 0 {
			count++
		} else {
			break
		}
		byteIndex = byteIndex + baseIndex + len(byteFind)
	}
	return
}
