package main

import (
	"coinbit/flagger"
	"coinbit/proto/pb"
	"flag"
	"fmt"
	"os"

	"github.com/lovoo/goka"
)

var (
	wallet = flag.String("wallet", "", "wallet_id identifier to be flagged as above-threshold")
	remove = flag.Bool("remove", false, "remove flag of wallet_id")
	broker = flag.String("broker", "localhost:9092", "bootstrap kafka broker")
)

func main() {
	flag.Parse()
	if *wallet == "" {
		fmt.Println("cannot remove flag of wallet_id ''")
		os.Exit(1)
	}

	emmiter, err := goka.NewEmitter(
		[]string{*broker},
		flagger.Stream,
		new(flagger.FlagEventCodec),
	)
	if err != nil {
		panic(err)
	}
	defer emmiter.Finish()

	err = emmiter.EmitSync(
		*wallet,
		&pb.FlagEvent{FlagRemoved: *remove},
	)
	if err != nil {
		panic(err)
	}
}
