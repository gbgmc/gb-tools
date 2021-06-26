package command

import (
	"log"

	"github.com/JakBaranowski/gb-tools/config"
)

// Mission parses the provided arguments and calls the appropriate subcommand.
// Heavily in WIP possibly abandoned. :(
func Mission(config config.Config) {
	log.Printf(
		"Please be advised that mission commands are experimental, and" +
			"the 'fix' command does not work at the moment. Use at your own risk.",
	)
	// subCommand, err := common.GetArgument(2)
	// common.Must(err)
	// switch subCommand {
	// case "fix":
	// 	missionPath, err := common.GetArgument(3)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	missionFix(config, missionPath)
	// case "read":
	// 	missionPath, err := common.GetArgument(3)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	missionRead(missionPath)
	// case "compare":
	// 	missionPath1, err := common.GetArgument(3)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	missionPath2, err := common.GetArgument(4)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	missionCompare(missionPath1, missionPath2)
	// case "print":
	// 	missionPath, err := common.GetArgument(3)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	missionPrint(missionPath, "/Script/GroundBranch.GBMissionData")
	// }
}

// func missionFix(config config.Config, missionPath string) {
// 	missionBackupPath := missionPath + ".backup_" + common.GetDateTimeString()

// 	fmt.Printf(
// 		"Creating mission backup at \"%s\".\n",
// 		missionBackupPath,
// 	)

// 	fileops.Copy(missionPath, missionBackupPath)
// 	fmt.Printf(
// 		"Attempting to fix invalid class paths in mission \"%s\".\n",
// 		missionPath,
// 	)

// 	src, err := os.OpenFile(missionPath, os.O_RDWR, 0644)
// 	common.Must(err)

// 	byteSrc := fileops.ReadFile(src)

// 	for key, value := range config.Missions.StringsToReplace {
// 		fmt.Printf("Fixing %s > %s.\n", key, value)
// 		keyByte := []byte(key)
// 		valueByte := []byte(value)
// 		count := fileops.CountOccurences(byteSrc, keyByte)
// 		log.Printf("Found %d occurences of %s.", count, key)
// 		byteSrc = fileops.ReplaceAndCount(byteSrc, keyByte, valueByte, false)
// 		log.Printf("Replaced %s with %s.", key, value)
// 	}

// 	_, err = src.WriteAt(byteSrc, 0)
// 	common.Must(err)

// 	src.Close()

// 	fmt.Printf(
// 		"Finished replacing invalid class paths in mission \"%s\".\n"+
// 			"Please remember to finis",
// 		missionPath,
// 	)
// }

// func missionRead(missionPath string) {
// 	mission.ParseFileAfterString(
// 		missionPath,
// 		"/Script/GroundBranch.GBMissionData",
// 		len("/Script/GroundBranch.GBMissionData")+1,
// 	)
// }

// func missionCompare(fileAPath string, fileBPath string) {
// 	fileABytes := fileops.OpenAndReadFile(fileAPath)
// 	fileBBytes := fileops.OpenAndReadFile(fileBPath)
// 	fmt.Printf("SizeA: %d, SizeB: %d.\n", len(fileABytes), len(fileBBytes))
// 	fileops.CompareBytes(fileABytes, fileBBytes)
// }

// func missionPrint(FilePath string, StartString string) {
// 	byteData := fileops.OpenAndReadFile(FilePath)
// 	startByte := []byte(StartString)
// 	length := len(byteData)
// 	index := bytes.Index(byteData, startByte) + len(StartString) + 1
// 	lastOk := false
// 	for ; index < length; index++ {
// 		if isAcceptedByte(byteData[index]) {
// 			fmt.Printf("%c", byteData[index])
// 			lastOk = true
// 		} else if lastOk {
// 			fmt.Print("\n")
// 			lastOk = false
// 		}
// 	}
// }

// func isAcceptedByte(byteCheck byte) bool {
// 	if (byteCheck >= 46 && byteCheck <= 57) ||
// 		(byteCheck >= 65 && byteCheck <= 90) ||
// 		(byteCheck >= 97 && byteCheck <= 122) ||
// 		byteCheck == 95 {
// 		return true
// 	}
// 	return false
// }
