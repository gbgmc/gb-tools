package fileops

import (
	"archive/zip"
	"io"
	"log"
	"os"

	"github.com/JakBaranowski/gb-tools/common"
)

// CompressFiles will iterate through the provided files list add them to
// the a zip compressed file saved under filename.
func CompressFiles(filename string, files []string) {
	log.Printf("Compressing '%s'.", filename)
	zipFile, err := os.Create(filename)
	common.Must(err)
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = addFileToZip(zipWriter, file); err != nil {
			log.Fatal()
		}
	}
	log.Printf("Finished compressing '%s'.", filename)
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	log.Printf("Adding file '%s' to archive.", filename)
	fileToZip, err := os.Open(filename)
	common.Must(err)
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	common.Must(err)

	header, err := zip.FileInfoHeader(info)
	common.Must(err)

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
