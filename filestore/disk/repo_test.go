package disk

import (
	"testing"
)

var (
	testFile      = "C:/Users/calvi/go/src/calvinechols.com/vault/go.mod"
	testWriteFile = "C:/Users/calvi/test/tempfile.txt"
	testData      = "This is a test string."
)

func TestCreateFile(t *testing.T) {
	fsrepo := NewFileStoreDiskRepo()
	err := fsrepo.CreateFile(testWriteFile, []byte(testData))
	if err != nil {
		t.Errorf("Error occurred while writing file %s: %v", testWriteFile, testData)
	}
}

func TestReadFile(t *testing.T) {
	fsrepo := NewFileStoreDiskRepo()
	_, err := fsrepo.ReadFile(testFile)

	if err != nil {
		t.Errorf("Error occurred reading file from %s: %v", testFile, err)
	}

	// if data != nil {
	// 	fmt.Printf("%v: %v\n", len(data), data)
	// }
}

func TestGetFileHandle(t *testing.T) {
	fsrepo := NewFileStoreDiskRepo()
	_, err := fsrepo.GetFileHandle(testFile)
	if err != nil {
		t.Errorf("Error occurred getting file handler: %v", err)
	}
	// fi, err := fh.Stat()
	// if err != nil {
	// 	t.Errorf("")
	// }
	// filesize := fi.Size()
	// fmt.Printf("%s read is %d bytes\n", testFile, filesize)
}
