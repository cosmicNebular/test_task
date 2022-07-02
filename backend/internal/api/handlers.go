package api

import (
	"encoding/json"
	"net/http"
	"test/internal/pkg/database/entity"
)

func (c *Controller) getPublicKey(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	key, err := c.KeyService.GetKey(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if key == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(key)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) savePublicKey(w http.ResponseWriter, r *http.Request) {
	k := new(entity.Key)
	err := json.NewDecoder(r.Body).Decode(k)
	for _, cookie := range r.Cookies() {
		if cookie.Name == "token" {
			k.Id = cookie.Value
		}
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.KeyService.SaveKey(*k)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *Controller) getFiles(w http.ResponseWriter, r *http.Request) {
	files, err := c.FileService.GetAllFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if files == nil {
		files = make([]entity.File, 0)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) getFile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("fileId")
	f, err := c.FileService.GetFile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if f == nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(f)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Controller) saveFile(w http.ResponseWriter, r *http.Request) {
	f := new(entity.File)
	err := json.NewDecoder(r.Body).Decode(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.FileService.SaveFile(*f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
