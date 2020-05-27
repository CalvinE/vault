package access

import (
	"encoding/json"
	"net/http"

	"calvinechols.com/vault/file"
)

// Handler is the http handler interface for access records
type Handler interface {
	AddAccess(w http.ResponseWriter, r *http.Request)
	AccessFile(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	AccessService Service
	FileService   file.Service
}

// NewAccessHandler returns a new access handler
func NewAccessHandler(accessService Service, fileService file.Service) Handler {
	return &handler{
		AccessService: accessService,
		FileService:   fileService,
	}
}

func (h *handler) AddAccess(w http.ResponseWriter, r *http.Request) {
	var access Access
	err := json.NewDecoder(r.Body).Decode(&access)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *handler) AccessFile(w http.ResponseWriter, r *http.Request) {

}
