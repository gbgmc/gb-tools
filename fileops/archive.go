package fileops

import (
	"archive/zip"
	"io"
	"log"
	"os"

	"github.com/JakBaranowski/gb-tools/common"
)

func CompressFiles(filename string, files []string) {
	zipFile, err := os.Create(filename)
	common.Must(err)
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err = CompressFile(zipWriter, file); err != nil {
			log.Fatal()
		}
	}
}

func CompressFile(zipWriter *zip.Writer, filename string) error {
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
