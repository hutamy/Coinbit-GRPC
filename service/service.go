package service

import (
	d "coinbit/deposit"
	"coinbit/flagger"
	"context"
	"log"
	"net/http"

	"coinbit/history"

	"github.com/gorilla/mux"
	"github.com/lovoo/goka"
)

func Run(brokers []string, stream goka.Stream) {
	view, err := goka.NewView(brokers, history.Table, new(d.DepositListCodec))
	if err != nil {
		panic(err)
	}
	go view.Run(context.Background())

	flagView, err := goka.NewView(brokers, flagger.Table, new(flagger.FlagValueCodec))
	if err != nil {
		panic(err)
	}
	go flagView.Run(context.Background())

	emitter, err := goka.NewEmitter(brokers, stream, new(d.DepositCodec))
	if err != nil {
		panic(err)
	}
	defer emitter.Finish()

	router := mux.NewRouter()
	router.HandleFunc("/deposit", deposit(emitter, stream)).Methods("POST")
	router.HandleFunc("/wallet/{id}", walletDetail(view, flagView)).Methods("GET")

	log.Printf("Listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
