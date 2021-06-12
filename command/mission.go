package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JakBaranowski/gb-tools/config"
	"github.com/JakBaranowski/gb-tools/fileops"
	"github.com/JakBaranowski/gb-tools/gbops"
)

func Mission(config config.Config) {
	subCommand, err := GetArgument(2)
	if err != nil {
		log.Fatal(err)
	}
	switch subCommand {
	case "fix":
		MissionFix(config)
	case "read":
		MissionRead()
	case "compare":
		MissionCompare()
	}
}

func MissionFix(config config.Config) {
	missionPath, err := GetArgument(3)
	if err != nil {
		log.Fatal(err)
	}

	missionBackupPath := missionPath + ".backup_" + getTime()

	fmt.Printf(
		"Creating mission backup at \"%s\".\n",
		missionBackupPath,
	)

	fileops.Copy(missionPath, missionBackupPath)
	fmt.Printf(
		"Attempting to fix invalid class paths in mission \"%s\".\n",
		missionPath,
	)

	src, err := os.OpenFile(missionPath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	bytes := fileops.ReadFile(src)

	for key, value := range config.Missions.StringsToReplace {
		fmt.Printf("Fixing %s > %s.\n", key, value)
		keyByte := []byte(key)
		valueByte := []byte(value)
		bytes = fileops.ReplaceBytes(bytes, keyByte, valueByte, true)
		bytes = fileops.ReplaceByteAtOffsetFromMatch(
			bytes,
			valueByte,
			-4,
			byte(len(valueByte)+1),
		)
	}

	_, err = src.WriteAt(bytes, 0)
	if err != nil {
		log.Fatal(err)
	}

	src.Close()

	fmt.Printf(
		"Finished fixing invalid class paths in mission \"%s\".\n",
		missionPath,
	)
}

func getTime() string {
	t := time.Now()
	return fmt.Sprintf(
		"%d%d%d%d%d%d",
		t.Year(),
		int(t.Month()),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}

func MissionRead() {
	missionPath, err := GetArgument(3)
	if err != nil {
		log.Fatal(err)
	}

	byteData := fileops.OpenAndReadFile(missionPath)
	StartString := []byte("/Script/GroundBranch.GBMissionData")
	index := bytes.Index(byteData, StartString) + len(StartString) + 1
	for index < len(byteData)-1 && index > 0 {
		_, _, _, _, index = gbops.GetProperty(byteData, index)
	}
}

func MissionCompare() {
	fileAPath, err := GetArgument(3)
	if err != nil {
		log.Fatal(err)
	}
	fileBPath, err := GetArgument(4)
	if err != nil {
		log.Fatal(err)
	}
	fileABytes := fileops.OpenAndReadFile(fileAPath)
	fileBBytes := fileops.OpenAndReadFile(fileBPath)
	fmt.Printf("SizeA: %d, SizeB: %d.\n", len(fileABytes), len(fileBBytes))
	fileops.CompareBytes(fileABytes, fileBBytes)
}
