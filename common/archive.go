package common

import (
	"archive/zip"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Compress will iterate through the provided files list add them to
// the a zip compressed file saved under filename.
func Compress(filename string, files []string, bodies map[string][]byte) {
	log.Printf("Compressing '%s'.", filename)

	zipFile, err := os.Create(filename)
	cobra.CheckErr(err)
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		addFileToZip(zipWriter, file)
	}

	for name, body := range bodies {
		addBodiesToZip(zipWriter, body, name)
	}

	log.Printf("Finished compressing '%s'.", filename)
}

// addFileToZip will add compressed files to the zip specified by filename.
func addFileToZip(zipWriter *zip.Writer, filename string) {
	log.Printf("Adding file '%s' to archive.", filename)
	fileToZip, err := os.Open(filename)
	cobra.CheckErr(err)
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	cobra.CheckErr(err)

	header, err := zip.FileInfoHeader(info)
	cobra.CheckErr(err)

	header.Name = filename
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	cobra.CheckErr(err)
	_, err = io.Copy(writer, fileToZip)
	cobra.CheckErr(err)
}

// addBytesToZip creates a file from the provided body and writes it to a zip file.
func addBodiesToZip(zipWriter *zip.Writer, body []byte, filename string) {
	log.Printf("Adding file bytes as file %s to archive.", filename)

	zipFile, err := zipWriter.Create(filename)
	cobra.CheckErr(err)

	_, err = zipFile.Write(body)
	cobra.CheckErr(err)
}
