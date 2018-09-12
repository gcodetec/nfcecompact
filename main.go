package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("listing files")
	path := flag.String("path", "./data", "o diretorio dos arquivos")
	flag.Parse()

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if isNfeFile(f) {
			year, month, _ := f.ModTime().Date()
			fmt.Println(f.Name(), year, int(month))
			yearPath := fmt.Sprintf("%s/%v", *path, year)
			createPath(yearPath)
			monthPath := fmt.Sprintf("%s/%v/%v", *path, year, int(month))
			createPath(monthPath)
			filePath := fmt.Sprintf("%s/%s", *path, f.Name())
			newFilePath := fmt.Sprintf("%s/%s", monthPath, f.Name())
			moveFileToPath(filePath, newFilePath)
		}
	}
}

func isNfeFile(file os.FileInfo) bool {
	if !file.IsDir() {
		return strings.Contains(file.Name(), "-nfe")
	}
	return false
}

func createPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}

func moveFileToPath(oldPath, newPath string) {
	fmt.Printf("Moving file %v to %v\n", oldPath, newPath)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}
}
