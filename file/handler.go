package file

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fileID, err := primitive.ObjectIDFromHex(fileIDString)
	if err != nil {
		// fileid is invalid
		log.Printf("file id provided is invalid: %v\n", fileIDString)
	}
	ownerID := owneridQuery[0]
	file, err := h.Service.GetFile(fileID)
	if err != nil {
		// error occurred while pulling file record
		log.Printf("An error occurred while pulling file record: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if file.OwnerID == ownerID {
		fileHandler, err := h.Service.GetFileHandleFromStorage(file.InternalFileName)
		defer fileHandler.Close()
		if err != nil {
			// error getting file from disk...
			log.Printf("error getting file from storage provider: %v - %+v\n", err, file)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		stat, err := fileHandler.Stat()
		if err != nil {
			// error getting stats on file from disk
			log.Printf("error gettings stats on file: %v - %+v\n", err, file)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var dataLen = stat.Size()
		fileLenStr := strconv.FormatInt(dataLen, 10)
		w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
		w.Header().Set("Content-Type", file.MimeType)
		w.Header().Set("Content-Length", fileLenStr)
		io.Copy(w, fileHandler)
		return
	}
	// person requesting file is not the owner...
	w.WriteHeader(http.StatusUnauthorized)

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
	mimeType := http.DetectContentType(fileData)
	ownerID := r.FormValue("ownerid")
	f := NewFile(mimeType, newFileName, handler.Filename, ownerID, "disk")
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
