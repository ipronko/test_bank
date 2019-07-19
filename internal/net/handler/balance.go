package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"test/internal/net/neterror"
	"test/internal/storage"
)

func NewBalanceHandler(storage *storage.Storage, maxBody int64) *BalanceHandler {
	return &BalanceHandler{
		Storage:     storage,
		MaxBodySize: maxBody,
	}
}

type BalanceHandler struct {
	Storage     *storage.Storage
	MaxBodySize int64
}

func (h BalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	reader := io.LimitReader(r.Body, h.MaxBodySize)

	jDecoder := json.NewDecoder(reader)
	in := BalanceIn{}

	err := jDecoder.Decode(&in)
	if err != nil {
		neterror.WriteError(w, err)
		return
	}

	balance, err := h.Storage.GetBalance(in.Id)
	if err != nil {
		neterror.WriteError(w, err)
		return
	}

	bytes, err := json.MarshalIndent(BalanceOut{Id: in.Id, Balance: balance}, "", "\t")
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

type BalanceIn struct {
	Id int64 `json:"id"`
}

type BalanceOut struct {
	Id      int64   `json:"id"`
	Balance float64 `json:"balance"`
}
