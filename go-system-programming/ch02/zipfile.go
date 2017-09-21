package main

import (
	"archive/zip"
	"io"
	"os"
)

func main() {
	oldFile, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()

	file, err := os.Create("zipped.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("old.txt")
	io.Copy(writer, oldFile)

}
