package nfcecompact

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// CompactFilesByCompetence move and compact files from the path, year and month
func CompactFilesByCompetence(path string, competenceYear, competenceMonth int) {
	filesToCompact := []string{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if isNfeFile(f) {
			year, month, _ := f.ModTime().Date()
			if year == competenceYear && int(month) == competenceMonth {
				fmt.Println(f.Name(), year, int(month))
				yearPath := fmt.Sprintf("%s/%v", path, year)
				createPath(yearPath)
				monthPath := fmt.Sprintf("%s/%v/%v", path, year, int(month))
				createPath(monthPath)
				filePath := fmt.Sprintf("%s/%s", path, f.Name())
				newFilePath := fmt.Sprintf("%s/%s", monthPath, f.Name())
				filesToCompact = append(filesToCompact, newFilePath)
				moveFileToPath(filePath, newFilePath)
			}
		}
	}

	output := fmt.Sprintf("%s/%d/%d-%d.zip", path, competenceYear, competenceYear, competenceMonth)

	err = zipFiles(output, filesToCompact)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Zipped File: " + output)
}

// Check if file is a nfe file
func isNfeFile(file os.FileInfo) bool {
	if !file.IsDir() {
		return strings.Contains(file.Name(), "-nfe")
	}
	return false
}

// Create a path based on name
func createPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

// move a file to new path

func moveFileToPath(oldPath, newPath string) {
	fmt.Printf("Moving file %v to %v\n", oldPath, newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}
}

// zipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func zipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer zipfile.Close()

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Using FileInfoHeader() above only uses the basename of the file. If we want
		// to preserve the folder structure we can overwrite this with the full path.
		fileNameParts := strings.Split(file, "/")
		header.Name = fileNameParts[len(fileNameParts)-1]

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		fmt.Println("Compacting file: ", file)
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		if _, err = io.Copy(writer, zipfile); err != nil {
			return err
		}
	}
	return nil
}
