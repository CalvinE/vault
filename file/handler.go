package file

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler interface {
	PutFile(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	MaxFileSizeBytes int64
	Service          Service
}

func NewFileHandler(fileService Service, maxFileSizeBytes int64) Handler {
	// defaults to 10GB?
	return &handler{
		MaxFileSizeBytes: maxFileSizeBytes,
		Service:          fileService,
	}
}

func (h *handler) PutFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(h.MaxFileSizeBytes)
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		// handle error
		return
	}
	// continue processing
	fileLength := handler.Size
	fileData := make([]byte, fileLength)
	_, err = file.Read(fileData)
	if err != nil {
		// handle error
		return
	}
	newFileName, err := h.Service.SaveFileToStorage(fileData)
	if err != nil {
		// handle error
		return
	}
	mimeType := http.DetectContentType(fileData)
	ownerID := r.FormValue("ownerid")
	f := NewFile(mimeType, newFileName, handler.Filename, ownerID, "disk")
	newFileID, err := h.Service.AddFile(f)
	if err != nil {
		log.Fatalf("failed to insert new fiel record... orphaned file saved to disk with name: %v - %v", newFileName, err)
		// handle errpr
	}
	// remove the internal name because its none of their business!
	f.InternalFileName = ""
	fmt.Printf("insertedFileID: %v", newFileID)
	jsonFile, err := json.Marshal(f)
	if err != nil {
		log.Fatalf("convert file to json failed: %v", err)
	}
	w.Write(jsonFile)
}
