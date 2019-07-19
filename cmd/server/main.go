package main

import (
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"test/internal/config"
	"test/internal/net/handler"
	"test/internal/storage"
)

const MaxBodySize = 1024 * 1024

func main() {
	cfg := config.New()
	stor := storage.NewStorage()

	createHandler := handler.NewCreateHandler(stor, MaxBodySize)
	balanceHandler := handler.NewBalanceHandler(stor, MaxBodySize)
	transferHandler := handler.NewTransferHandler(stor, MaxBodySize)

	router := mux.NewRouter()
	router.Handle("/", createHandler).Methods(http.MethodPost)
	router.Handle("/", balanceHandler).Methods(http.MethodGet)
	router.Handle("/", transferHandler).Methods(http.MethodPatch)

	err := http.ListenAndServe(cfg.GetString(config.ServeAddr.String()), router)
	if err != nil {
		log.Fatal("failed serve", zap.Error(err))
	}
}
