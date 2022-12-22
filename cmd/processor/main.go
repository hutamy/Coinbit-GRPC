package main

import (
	"coinbit/flagger"
	"coinbit/history"
	"coinbit/threshold"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

var (
	brokers      = []string{"localhost:9092"}
	runHistory   = flag.Bool("history", false, "run history processor")
	runFlagger   = flag.Bool("flagger", false, "run flagger processor")
	runThreshold = flag.Bool("threshold", false, "run threshold processor")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	if *runHistory {
		log.Println("starting history")
		grp.Go(history.Run(ctx, brokers))
	}

	if *runFlagger {
		log.Println("starting flagger")
		grp.Go(flagger.Run(ctx, brokers))
	}

	if *runThreshold {
		log.Println("starting threshold")
		grp.Go(threshold.Run(ctx, brokers))
	}

	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-waiter:
	case <-ctx.Done():
	}
	cancel()
	err := grp.Wait()
	if err != nil {
		log.Println(err)
	}
	log.Println("done")
}
