package nfcecompact

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var fileName = "/tmp/nfcetmp-nfe.xml"
var tmpPath = "/tmp/nfcecompact-test"
var newFileName = tmpPath + "/nfcetmp-nfe.xml"
var pathToCopy = "/tmp/nfces"

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	createPath(tmpPath)
	createPath(pathToCopy)
	createTestFile()
}

func tearDown() {
	_ = os.Remove(fileName)
	_ = os.RemoveAll(tmpPath)
	_ = os.RemoveAll(pathToCopy)
}

func createTestFile() {
	names := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, name := range names {
		content := []byte("<nfe>" + string(name) + "</nfe>")
		fileName := fmt.Sprintf("%s/arquivo-%v-nfe.xml", pathToCopy, name)
		err := ioutil.WriteFile(fileName, content, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func TestIsNfeFile(t *testing.T) {
	fileName := "some-file-nfe.xml"
	v := isNfeFile(fileName)
	if v != true {
		t.Error("Expected true, got ", v)
	}
}

func TesNotIsNfeFile(t *testing.T) {
	fileName := "some-file.xml"
	v := isNfeFile(fileName)
	if v != false {
		t.Error("Expected false, got ", v)
	}
}

func TestCreatePath(t *testing.T) {
	path := tmpPath + "/create-path"
	createPath(path)
	fileInfo, _ := os.Stat(path)
	isDir := fileInfo.IsDir()
	if !isDir {
		t.Error("Expected true, got ", isDir)
	}
	_ = os.Remove(path)
}

func TestCopyAllFilesToPath(t *testing.T) {
	totalBytes := CopyAllFilesToPath(pathToCopy, tmpPath)
	files := readDir(tmpPath)
	length := len(files)
	if length == 0 {
		t.Errorf("files not copied")
	}

	if totalBytes == 0 {
		t.Errorf("any file copied. Total bytes %v", totalBytes)
	}
}
func TestMoveFileToPath(t *testing.T) {
	moveFileToPath(fileName, newFileName)
	_, err := os.Stat(newFileName)
	if os.IsNotExist(err) {
		t.Errorf("file %v not exists", newFileName)
	}
}

func TestReadDir(t *testing.T) {
	path := tmpPath
	files := readDir(path)
	if len(files) == 0 {
		t.Errorf("files not found")
	}
}

func TestZipFiles(t *testing.T) {
	outFile := "/tmp/nfcecompact.zip"
	filesList := []string{newFileName}
	err := zipFiles(outFile, filesList)
	if err != nil {
		t.Errorf("files can not be compacted on %v - %v", outFile, err)
	}
	_ = os.Remove(outFile)
}

func TestCompactFilesByCompetence(t *testing.T) {
	now := time.Now().Local()
	CompactFilesByCompetence(tmpPath, now.Year(), int(now.Month()))
	fileDir := fmt.Sprintf("%s/%v/%v/%v", tmpPath, now.Year(), int(now.Month()), "nfcetmp-nfe.xml")
	_, err := os.Stat(fileDir)
	if err != nil {
		t.Errorf("deu zebra")
	}
}
