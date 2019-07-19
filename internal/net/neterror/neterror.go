package neterror

import (
	"encoding/json"
	"net/http"
	"test/internal/storage"
)

func WriteError(w http.ResponseWriter, err error) {
	if _, ok := err.(storage.AccountNotFoundErr); ok {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	bytes, _ := json.Marshal(map[string]string{"err": err.Error()})
	w.Write(bytes)
}
