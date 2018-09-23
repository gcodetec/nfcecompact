package nfcecompact

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var fileName = "/tmp/nfcetmp-nfe.xml"
var tmpPath = "/tmp/nfcecompact-test"
var newFileName = tmpPath + "/nfcetmp-nfe.xml"

func TestMain(m *testing.M) {
	setUp()
	retCode := m.Run()
	tearDown()
	os.Exit(retCode)
}

func setUp() {
	os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	createPath(tmpPath)
}

func tearDown() {
	_ = os.Remove(fileName)
	_ = os.RemoveAll(tmpPath)
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

func TestMoveFileToPath(t *testing.T) {
	moveFileToPath(fileName, newFileName)
	_, err := os.Stat(newFileName)
	if os.IsNotExist(err) {
		t.Errorf("file %v not exists", newFileName)
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
