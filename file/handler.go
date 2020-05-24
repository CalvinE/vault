package file

import (
	"fmt"
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
	fmt.Printf("%v", newFileName)
}
