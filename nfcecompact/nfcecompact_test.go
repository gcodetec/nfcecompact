package nfcecompact

import (
	"fmt"
	"os"
	"testing"
)

func TestIsNfeFile(t *testing.T) {
	fileName := "some-file-nfe.xml"
	v := isNfeFile(fileName)
	if v != true {
		t.Error("Expected true, got ", v)
	}
}

func TesNottIsNfeFile(t *testing.T) {
	fileName := "some-file.xml"
	v := isNfeFile(fileName)
	if v != false {
		t.Error("Expected false, got ", v)
	}
}

func TestCreatePath(t *testing.T) {
	path := "/tmp/nfececompact-test"
	createPath(path)
	fileInfo, _ := os.Stat(path)
	isDir := fileInfo.IsDir()
	if !isDir {
		t.Error("Expected true, got ", isDir)
	}
	err := os.Remove(path)
	if err != nil {
		fmt.Println("...")
	}
}
