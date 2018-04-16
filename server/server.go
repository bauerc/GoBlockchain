package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bauerc/GoBlockchain/blocks"
	"github.com/bauerc/GoBlockchain/chain"
	"github.com/gorilla/mux"
)

var Blockchain chain.Blockchain

func Run() error {
	Blockchain = chain.Genesis(blocks.BPMGenesisBlock)
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", httpAddr)
	s := http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}
