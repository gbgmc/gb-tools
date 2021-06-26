package mission

import (
	"fmt"
	"log"

	"github.com/JakBaranowski/gb-tools/fileops"
)

type Property struct {
	Id           int
	ParentId     int
	ChildrenId   []int
	Name         string
	Type         string
	DetailedType string
	Value        []byte
}

func NewPropertyFromBytes(name []byte, propertyType []byte, detailedType []byte, value []byte) Property {
	nameStr := string(name)
	propertyTypeStr := string(propertyType)
	var detailedTypeStr string
	if detailedType == nil {
		detailedTypeStr = ""
	} else {
		detailedTypeStr = string(detailedType)
	}
	return Property{
		Id:           -1,
		ParentId:     -1,
		Name:         nameStr,
		Type:         propertyTypeStr,
		DetailedType: detailedTypeStr,
		Value:        value,
	}
}

func (p *Property) AddChildren(childrenId int) {
	p.ChildrenId = append(p.ChildrenId, childrenId)
}

func (p *Property) AddParent(parentId int) {
	p.ParentId = parentId
}

func (p *Property) ToString() string {
	return fmt.Sprintf(
		"Name: %s; Type: %s; DetailedType %s; Value %d; ID: %d; ParentID: %d; ChildrenId: %v.\n",
		p.Name,
		p.Type,
		p.DetailedType,
		p.Value,
		p.Id,
		p.ParentId,
		p.ChildrenId,
	)
}

// Gets property name, type, value and it's ending index.
// VarNameSize 00 00 00 VarName 00 VarTypeSize 00 00 00 VarType 00
func GetProperty(
	SrcByte []byte,
	StartIndex int,
) (Property, int) {
	var (
		PropertyName         string
		PropertyType         string
		PropertyDetailedType string
		PropertyValue        []byte
		index                int
	)
	PropertyName, index = GetData(SrcByte, StartIndex)
	if string(PropertyName) == "None" || fileops.GetSize([]byte(PropertyName)) == 0 {
		return Property{}, -1
	}
	PropertyType, index = GetData(SrcByte, index)
	switch PropertyType {
	case "StrProperty":
		PropertyValue, index = GetStringValue(SrcByte, index)
		// log.Printf("%s %s = '%s'", VariableName, VariableType, VariableValue)
	case "NameProperty":
		PropertyValue, index = GetNameValue(SrcByte, index)
		// log.Printf("%s %s = '%s'", VariableName, VariableType, VariableValue)
	case "TextProperty":
		PropertyValue, index = GetTextValue(SrcByte, index)
		// log.Printf("%s %s = '%s'", VariableName, VariableType, VariableValue)
	case "BoolProperty":
		PropertyValue[0], index = GetBoolValue(SrcByte, index)
		// log.Printf("%s %s = %b", VariableName, VariableType, VariableValue)
	case "ByteProperty":
		PropertyValue, index = GetByteValue(SrcByte, index)
		// log.Printf("%s %s = %d", VariableName, VariableType, VariableValue)
	case "ArrayProperty":
		PropertyValue, PropertyDetailedType, index = GetArrayValue(SrcByte, index)
		// log.Printf(
		// 	"%s %s [%s] = %v",
		// 	VariableName,
		// 	VariableType,
		// 	VariableValueType,
		// 	VariableValue,
		// )
	case "StructProperty":
		PropertyValue, PropertyDetailedType, index = GetStructValue(SrcByte, index)
		// log.Printf(
		// 	"%s %s [%s] = %v",
		// 	VariableName,
		// 	VariableType,
		// 	VariableValueType,
		// 	VariableValue,
		// )
	default:
		log.Printf("Unknown type %s %s!", PropertyType, PropertyName)
	}
	return Property{
		Name:         PropertyName,
		Type:         PropertyType,
		DetailedType: PropertyDetailedType,
		Value:        PropertyValue,
	}, index
}

// Parses basic data, i.e. variable name or variable type declaration
func GetBasic(SrcByte []byte, LengthIndex int) (Value []byte, ValueEnd int) {
	ValueLength := int(SrcByte[LengthIndex])
	if ValueLength == 0 {
		return nil, LengthIndex
	}
	ValueStart := LengthIndex + 4
	ValueEnd = ValueStart + ValueLength
	Value = SrcByte[ValueStart : ValueEnd-1]
	// log.Printf("'%s' @%d-%d [%d]", Value, ValueStart, ValueEnd, ValueLength)
	return Value, ValueEnd
}

// Parses basic data, i.e. variable name or variable type declaration
func GetData(SrcByte []byte, LengthIndex int) (Value string, ValueEnd int) {
	ValueByte, ValueEnd := GetBasic(SrcByte, LengthIndex)
	Value = string(ValueByte)
	return
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
			"Invalid generic string: %d does not equal %d.",
			SmallValueLength+LengthDiff,
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
) (Value []byte, ValueEnd int) {
	SmallLenght := int(SrcByte[FirstLengthIndex])
	BigLengthIndex := FirstLengthIndex + 8
	BigLength := int(SrcByte[BigLengthIndex])
	if SmallLenght+4 != BigLength {
		log.Printf(
			"Invalid byte: Small length %d does not equal big length %d",
			SmallLenght+4,
			BigLength,
		)
	}

	ValueStart := BigLengthIndex + 4
	ValueEnd = ValueStart + BigLength
	Value = SrcByte[ValueStart : ValueEnd-1]

	if string(Value) == "None" {
		Value[0] = SrcByte[ValueEnd+3]
	}

	return
}

// Gets BoolProperty Value, as well as it's ending index ValueEnd.
// `00 00 00 00 00 00 00 00 00 Value`
func GetBoolValue(
	SrcByte []byte,
	FirstLengthIndex int,
) (Value byte, ValueEnd int) {
	ValueIndex := FirstLengthIndex + 8
	Value = SrcByte[ValueIndex]
	ValueEnd = ValueIndex + 2
	// log.Printf("'%d' @%d", Value, ValueIndex)
	return
}

// Gets StructProperty Value, as well as it's ending index ValueEnd.
func getComplexValue(
	SrcByte []byte,
	LengthIndex int,
	InnerTypeSkip int,
	ValueSkip int,
) (
	Value []byte,
	InnerType string,
	ValueEnd int,
) {
	ValueLenght := fileops.GetSize(SrcByte[LengthIndex : LengthIndex+4])

	InnerTypeLengthIndex := LengthIndex + 8
	InnerTypeLength := int(SrcByte[InnerTypeLengthIndex])
	InnerTypeStart := InnerTypeLengthIndex + InnerTypeSkip
	InnerTypeEnd := InnerTypeStart + InnerTypeLength - 1
	InnerType = string(SrcByte[InnerTypeStart:InnerTypeEnd])

	ValueStart := InnerTypeEnd + ValueSkip
	ValueEnd = ValueStart + ValueLenght
	Value = SrcByte[ValueStart:ValueEnd]

	return
}

// Parses and gets StructProperty Value, as well as it's ending index ValueEnd.
// `StructSizeBytes[4] 00 00 00 00 StructTypeNameSize 00 00 00 StructTypeName 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 Value`
func GetStructValue(
	SrcByte []byte,
	LengthIndex int,
) (
	Value []byte,
	StructType string,
	ValueEnd int,
) {
	Value, StructType, ValueEnd = getComplexValue(SrcByte, LengthIndex, 4, 18)
	return
}

// Parses and gets ArrayProperty Value, as well as it's ending index ValueEnd.
// `ArraySizeBytes[4] 00 00 00 00 EntryTypeSize 00 00 00 EntryType 00 00 ArrayEntryAmount 00 00 00 Value`
func GetArrayValue(
	SrcByte []byte,
	LengthIndex int,
) (
	Value []byte,
	EntryType string,
	ValueEnd int,
) {
	Value, EntryType, ValueEnd = getComplexValue(SrcByte, LengthIndex, 4, 2)
	return
}

// func GetStructArrayValue(ArrayValue []byte) []byte {

// }
