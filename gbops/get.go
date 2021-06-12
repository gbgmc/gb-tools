package gbops

import (
	"bytes"
	"log"
	"math"
)

func SplitAtStart(SrcByte []byte) (Prefix []byte, Base []byte) {
	StartString := []byte("/Script/GroundBranch.GBMissionData")
	StartIndex := bytes.Index(SrcByte, StartString) + len(StartString) + 1
	Prefix, Base = SrcByte[:StartIndex-1], SrcByte[StartIndex:]
	return
}

// Gets property name, type, value and it's ending index.
// VarNameSize 00 00 00 VarName 00 VarTypeSize 00 00 00 VarType 00
func GetProperty(
	SrcByte []byte,
	StartIndex int,
) (
	VariableName []byte,
	VariableType []byte,
	VariableValue []byte,
	VariableValueType []byte,
	EndingIndex int,
) {
	VariableValueType = nil
	VariableName, EndingIndex = GetBasic(SrcByte, StartIndex)
	if string(VariableName) == "None" {
		return
	}
	if getSize(VariableName) == 0 {
		return nil, nil, nil, nil, -1
	}
	VariableType, EndingIndex = GetBasic(SrcByte, EndingIndex)
	switch string(VariableType) {
	case "StrProperty":
		VariableValue, EndingIndex = GetStringValue(SrcByte, EndingIndex)
		log.Printf("%s %s = '%s'", VariableName, VariableType, VariableValue)
	case "NameProperty":
		VariableValue, EndingIndex = GetNameValue(SrcByte, EndingIndex)
		log.Printf("%s %s = '%s'", VariableName, VariableType, VariableValue)
	case "TextProperty":
		VariableValue, EndingIndex = GetTextValue(SrcByte, EndingIndex)
		log.Printf("%s %s = '%s'", VariableName, VariableType, VariableValue)
	case "ByteProperty":
		VariableValue[0], EndingIndex = GetByteValue(SrcByte, EndingIndex)
		log.Printf("%s %s = %d", VariableName, VariableType, VariableValue)
	case "BoolProperty":
		VariableValue[0], EndingIndex = GetBoolValue(SrcByte, EndingIndex)
		log.Printf("%s %s = %b", VariableName, VariableType, VariableValue)
	case "ArrayProperty":
		VariableValue, VariableValueType, EndingIndex = GetArrayValue(SrcByte, EndingIndex)
		log.Printf(
			"%s %s [%s] = %v",
			VariableName,
			VariableType,
			VariableValueType,
			VariableValue,
		)
	case "StructProperty":
		VariableValue, VariableValueType, EndingIndex = GetStructValue(SrcByte, EndingIndex)
		log.Printf(
			"%s %s [%s] = %v",
			VariableName,
			VariableType,
			VariableValueType,
			VariableValue,
		)
	default:
		println("Unknown type!")
	}
	return
}

// Parses basic data, i.e. variable name or variable type declaration
func GetBasic(SrcByte []byte, LengthIndex int) (Value []byte, ValueEnd int) {
	ValueLength := int(SrcByte[LengthIndex])
	if ValueLength == 0 {
		return []byte{0}, LengthIndex
	}
	ValueStart := LengthIndex + 4
	ValueEnd = ValueStart + ValueLength
	Value = SrcByte[ValueStart : ValueEnd-1]
	// log.Printf("'%s' @%d-%d [%d]", Value, ValueStart, ValueEnd, ValueLength)
	return Value, ValueEnd
}

// Parses StrProperty and returns the Value in byte array as well as the ending byte index
// `ValueSize+4 00 00 00 00 00 00 00 00 ValueSize 00 00 00 Value 00`
func GetStringValue(SrcByte []byte, FirstLengthIndex int) (Value []byte, ValueEnd int) {
	return getGenericStringValue(SrcByte, FirstLengthIndex, 9, 4)
}

// Parses NameProperty and returns the Value in byte array as well as the ending byte index
// `ValueSize-4 00 00 00 00 00 00 00 00 ValueSize 00 00 00 Value 00`
func GetNameValue(SrcByte []byte, FirstLengthIndex int) (Value []byte, ValueEnd int) {
	return getGenericStringValue(SrcByte, FirstLengthIndex, 9, 4)
}

// Parses TextProperty and returns the Value in byte array as well as the ending byte index
// `ValueSize+13 00 00 00 00 00 00 00 00 02 00 00 00 FF 01 00 00 00 ValueSize 00 00 00 Value 00`
func GetTextValue(SrcByte []byte, FirstLengthIndex int) (Value []byte, ValueEnd int) {
	return getGenericStringValue(SrcByte, FirstLengthIndex, 18, 13)
}

// Generic function for parsing string like properties (StrProperty, NameProperty, TextProperty)
func getGenericStringValue(
	SrcByte []byte,
	FirstLengthIndex int,
	LengthSkip int,
	LengthDiff int,
) (Value []byte, ValueEnd int) {
	BigValueLength := int(SrcByte[FirstLengthIndex])
	SmallValueLengthIndex := FirstLengthIndex + LengthSkip
	SmallValueLength := int(SrcByte[SmallValueLengthIndex])
	if SmallValueLength == 0 {
		return
	}
	if SmallValueLength+LengthDiff != BigValueLength {
		log.Printf(
			"Invalid generic string: Small length (%d + %d) not equal to big length (%d)",
			SmallValueLength,
			LengthDiff,
			BigValueLength,
		)
	}
	ValueStart := SmallValueLengthIndex + 4
	ValueEnd = ValueStart + SmallValueLength - 1
	Value = SrcByte[ValueStart:ValueEnd]
	ValueEnd = ValueEnd + 1
	// log.Printf("'%s' @%d-%d [%d]", Value, ValueStart, ValueEnd, SmallValueLength)
	return
}

// Parses and gets ByteProperty Value, as well as it's ending index ValueEnd.
// `ValueSize-4 00 00 00 00 00 00 00 ValueSize 00 00 00 N o n e 00 00 Value`
func GetByteValue(
	SrcByte []byte,
	FirstLengthIndex int,
) (Value byte, ValueEnd int) {
	return getGenericByteValue(SrcByte, FirstLengthIndex, 18)
}

// Parses and gets BoolProperty Value, as well as it's ending index ValueEnd.
// `00 00 00 00 00 00 00 00 00 Value`
func GetBoolValue(
	SrcByte []byte,
	FirstLengthIndex int,
) (Value byte, ValueEnd int) {
	return getGenericByteValue(SrcByte, FirstLengthIndex, 8)
}

func getGenericByteValue(
	SrcByte []byte,
	FirstLengthIndex int,
	ValueSkip int,
) (Value byte, ValueEnd int) {
	ValueIndex := FirstLengthIndex + ValueSkip
	Value = SrcByte[ValueIndex]
	ValueEnd = ValueIndex + 2
	// log.Printf("'%d' @%d", Value, ValueIndex)
	return
}

// Parses and gets ArrayProperty Value, as well as it's ending index ValueEnd.
// `ArraySizeBytes[4] 00 00 00 00 EntryTypeSize 00 00 00 EntryType 00 00 ArrayEntryAmount 00 00 00 Value`
func GetArrayValue(
	SrcByte []byte,
	LengthIndex int,
) (Value []byte, EntryType []byte, ValueEnd int) {
	ValueLenght := getSize(SrcByte[LengthIndex : LengthIndex+4])

	EntryTypeLengthIndex := LengthIndex + 8
	EntryTypeLength := int(SrcByte[EntryTypeLengthIndex])
	EntryTypeStart := LengthIndex + 12
	EntryTypeEnd := EntryTypeStart + EntryTypeLength - 1
	EntryType = SrcByte[EntryTypeStart:EntryTypeEnd]

	ValueStart := EntryTypeEnd + 2
	ValueEnd = ValueStart + ValueLenght
	Value = SrcByte[ValueStart:ValueEnd]

	// log.Printf(
	// 	"Array %d@%d-%d [%s@%d-%d] = %v",
	// 	EntriesLenght,
	// 	LengthIndex,
	// 	LengthIndex+4,
	// 	EntryType,
	// 	EntriesStart,
	// 	EntriesEnd,
	// 	EntriesByte,
	// )
	return
}

// func ParseArrayValue(ValueBytes []byte) (Values [][]byte) {
// 	EntriesAmount := int(ValueBytes[0])
// 	Values = make([][]byte, EntriesAmount)
// 	ValueLengthIndex := 4
// 	for i := 0; i < EntriesAmount; i++ {
// 		ValueLength := int(ValueBytes[ValueLengthIndex])
// 		ValueStart := ValueLengthIndex + 4
// 		ValueEnd := ValueStart + ValueLength - 1
// 		Values[i] = ValueBytes[ValueStart:ValueEnd]
// 		ValueLengthIndex = ValueEnd + 1
// 		if ValueLengthIndex >= len(ValueBytes) {
// 			break
// 		}
// 	}
// 	return
// }

// Parses and gets StructProperty Value, as well as it's ending index ValueEnd.
// `StructSizeBytes[4] 00 00 00 00 StructTypeNameSize 00 00 00 StructTypeName 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 Value`
func GetStructValue(
	SrcByte []byte,
	LengthIndex int,
) (
	Value []byte,
	StructType []byte,
	ValueEnd int,
) {
	ValueLenght := getSize(SrcByte[LengthIndex : LengthIndex+4])

	StructTypeLengthIndex := LengthIndex + 8
	StructTypeLength := int(SrcByte[StructTypeLengthIndex])
	StructTypeStart := StructTypeLengthIndex + 4
	StructTypeEnd := StructTypeStart + StructTypeLength - 1
	StructType = SrcByte[StructTypeStart:StructTypeEnd]

	ValueStart := StructTypeEnd + 18
	ValueEnd = ValueStart + ValueLenght
	Value = SrcByte[ValueStart:ValueEnd]
	return
}

// Parses a multiple byte long byte sized. They are written in right to left
// format.
func getSize(byteSize []byte) int {
	length := len(byteSize)
	sum := 0
	for i := range byteSize {
		sum += int(byteSize[length-i-1]) * int(math.Pow(256, float64(length-i-1)))
	}
	return sum
}
