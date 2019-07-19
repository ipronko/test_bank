package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"test/internal/net/neterror"
	"test/internal/storage"
)

func NewCreateHandler(storage *storage.Storage, maxBody int64) *CreateHandler {
	return &CreateHandler{
		Storage:     storage,
		MaxBodySize: maxBody,
	}
}

type CreateHandler struct {
	Storage     *storage.Storage
	MaxBodySize int64
}

func (h CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reader := io.LimitReader(r.Body, h.MaxBodySize)

	jDecoder := json.NewDecoder(reader)
	in := CreateIn{}

	err := jDecoder.Decode(&in)
	if err != nil {
		return
	}

	id := h.Storage.CreateAccount(in.Balance)

	bytes, err := json.MarshalIndent(CreateOut{Id: id}, "", "\t")
	if err != nil {
		neterror.WriteError(w, err)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		neterror.WriteError(w, err)
		return
	}
}

type CreateIn struct {
	Balance float64 `json:"balance"`
}

type CreateOut struct {
	Id int64 `json:"id"`
}
