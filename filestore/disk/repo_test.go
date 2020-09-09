package disk

import (
	"fmt"
	"os"
	"testing"
)

var (
	testFile      = "go.mod"
	testWriteFile = "tempfile.txt"
	testData      = "This is a test string."
)

func TestCreateFile(t *testing.T) {
	fsrepo := NewFileStoreDiskRepo("C:/Users/calvi/test/")
	err := fsrepo.CreateFile(testWriteFile, []byte(testData))
	if err != nil {
		t.Errorf("Error occurred while writing file %s: %v", testWriteFile, testData)
	}
}

func TestReadFile(t *testing.T) {
	workingDir, _ := os.Getwd()
	t.Log("workingDir=", workingDir)
	targetDir := fmt.Sprintf("%s%s..%s..", workingDir, string(os.PathSeparator), string(os.PathSeparator))
	t.Log("targetDir=", targetDir)
	fsrepo := NewFileStoreDiskRepo(targetDir) //("C:/Users/calvi/go/src/calvinechols.com/vault/")
	_, err := fsrepo.ReadFile(testFile)
	if err != nil {
		t.Errorf("Error occurred reading file from %s: %v", testFile, err)
	}
}

func TestGetFileHandle(t *testing.T) {
	workingDir, _ := os.Getwd()
	t.Log("workingDir=", workingDir)
	targetDir := fmt.Sprintf("%s%s..%s..", workingDir, string(os.PathSeparator), string(os.PathSeparator))
	t.Log("targetDir=", targetDir)
	fsrepo := NewFileStoreDiskRepo(targetDir) //("C:/Users/calvi/go/src/calvinechols.com/vault/")
	_, err := fsrepo.GetFileHandle(testFile)
	if err != nil {
		t.Errorf("Error occurred getting file handler: %v", err)
	}
}
