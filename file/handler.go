package file

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	GetFile(w http.ResponseWriter, r *http.Request)
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

func (h *handler) GetFile(w http.ResponseWriter, r *http.Request) {
	// TODO: move some of this code to the service... business logic should be in service not handler...
	fileidQuery, ok := r.URL.Query()["fileid"]
	if !ok || len(fileidQuery[0]) < 1 {
		// bad request, no fileid provided
		log.Print("no file id provided\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO get ownerID from JWT
	owneridQuery, ok := r.URL.Query()["ownerid"]
	if !ok || len(owneridQuery[0]) < 1 {
		// bad request, no fileid provided
		log.Println("no owner id provided")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileIDString := fileidQuery[0]
	ownerID := owneridQuery[0]
	file, fileHandler, err := h.Service.RetreiveFile(fileIDString, ownerID)
	if err != nil {
		// TODO figure this stuff out... how to use errors.Is and errors.As...
		if err.Error() == "unauthorized attempt" {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer fileHandler.Close()
	if err != nil {
		// error getting file from disk...
		errMsg := fmt.Sprintf("error getting file from storage provider: %v - %+v", err, file)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	fileLenStr := strconv.FormatInt(file.FileSize, 10)
	w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
	w.Header().Set("Content-Type", file.MimeType)
	w.Header().Set("Content-Length", fileLenStr)
	io.Copy(w, fileHandler)
}

func (h *handler) PutFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(h.MaxFileSizeBytes)
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// continue processing
	fileLength := handler.Size
	fileData := make([]byte, fileLength)
	_, err = file.Read(fileData)
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newFileName, err := h.Service.SaveFileToStorage(fileData)
	if err != nil {
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO try to make map of file extensions to mime types for this, also potentially allow users to specify mime type...
	mimeType := http.DetectContentType(fileData)
	ownerID := r.FormValue("ownerid")
	f := NewFile(mimeType, newFileName, handler.Filename, ownerID, "disk", fileLength)
	newFileID, err := h.Service.AddFile(f)
	if err != nil {
		log.Fatalf("failed to insert new fili record... orphaned file saved to disk with name: %v - %v", newFileName, err)
		// handle error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// remove the internal name because its none of their business!
	f.InternalFileName = ""
	fmt.Printf("insertedFileID: %v", newFileID)
	jsonFile, err := json.Marshal(f)
	if err != nil {
		log.Fatalf("convert file to json failed: %v - %+v", err, f)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonFile)
}
