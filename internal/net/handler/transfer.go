package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"test/internal/net/neterror"
	"test/internal/storage"
)

func NewTransferHandler(storage *storage.Storage, maxBody int64) *TransferHandler {
	return &TransferHandler{
		Storage:     storage,
		MaxBodySize: maxBody,
	}
}

type TransferHandler struct {
	Storage     *storage.Storage
	MaxBodySize int64
}

func (h TransferHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reader := io.LimitReader(r.Body, h.MaxBodySize)

	jDecoder := json.NewDecoder(reader)
	in := TransferIn{}

	err := jDecoder.Decode(&in)
	if err != nil {
		neterror.WriteError(w, err)
		return
	}

	err = h.Storage.Transfer(in.FromID, in.ToID, in.Sum)
	if err != nil {
		neterror.WriteError(w, err)
		return
	}
}

type TransferIn struct {
	FromID int64   `json:"from"`
	ToID   int64   `json:"to"`
	Sum    float64 `json:"sum"`
}
