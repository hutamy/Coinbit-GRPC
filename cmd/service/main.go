package main

import (
	"coinbit/deposit"
	"flag"

	"coinbit/service"
)

var (
	broker = flag.String("broker", "localhost:9092", "bootstrap Kafka broker")
)

func main() {
	flag.Parse()
	service.Run([]string{*broker}, deposit.DepositStream)
}
