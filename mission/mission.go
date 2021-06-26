package mission

import (
	"bytes"
	"log"

	"github.com/JakBaranowski/gb-tools/fileops"
)

type Mission struct {
	properties []*Property
}

func NewMission() Mission {
	return Mission{
		properties: make([]*Property, 0),
	}
}

func (m *Mission) AddProperty(property *Property) {
	m.properties = append(m.properties, property)
	property.Id = len(m.properties) - 1
}

func ParseFileFromByte(filePath string, start []byte, skip int) (nission Mission) {
	byteData := fileops.OpenAndReadFile(filePath)
	index := bytes.Index(byteData, start) + skip
	return ParseBytesFromIndex(byteData, index)
}

func ParseFileAfterString(filePath string, start string, skip int) (nission Mission) {
	return ParseFileFromByte(filePath, []byte(start), skip)
}

func ParseFileFromIndex(filePath string, index int) (nission Mission) {
	byteData := fileops.OpenAndReadFile(filePath)
	return ParseBytesFromIndex(byteData, index)
}

func ParseBytesFromIndex(byteSrc []byte, index int) (mission Mission) {
	mission = NewMission()
	parseBytes(byteSrc, index, &mission, -1)
	for _, value := range mission.properties {
		log.Print(value.ToString())
	}
	return
}

func parseBytes(byteSrc []byte, index int, mission *Mission, parentId int) {
	for index < len(byteSrc)-1 && index >= 0 {
		var NewProperty Property
		NewProperty, index = GetProperty(byteSrc, index)
		mission.AddProperty(&NewProperty)
		if parentId >= 0 {
			NewProperty.AddParent(parentId)
			mission.properties[parentId].AddChildren(NewProperty.Id)
		}
	}
}

func SplitAtValue(SrcByte []byte, Value string) (Prefix []byte, Base []byte) {
	StartString := []byte(Value)
	StartIndex := bytes.Index(SrcByte, StartString) + len(StartString) + 1
	Prefix, Base = SrcByte[:StartIndex-1], SrcByte[StartIndex:]
	return
}
