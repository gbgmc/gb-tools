package fileops

import (
	"bytes"
	"fmt"
	"strconv"
)

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
