# Coinbit

## Architecture requirement:
- Use [Goka](https://github.com/lovoo/goka) to build the service above.
- Use [protobuf](https://developers.google.com/protocol-buffers/docs/gotutorial) when encoding/decoding payload to/from Kafka broker.
- Use Goka's Local storage mechanism, if a database is required.

## Prerequisites
- Golang 1.19
- Docker is required for simple setup env using docker compose
- Makefile to simplify command execution

## Development 

### Setup development environment

- Run `docker compose up -d` or run `make up` to setup environment
- Install dependencies
```cmd
go mod tidy
go mod vendor -- idk harus ga
```
Or simply run `make dep` command.
- Make sure you create the required Kafka topics (deposit, flag-wallet,  balance-table, threshold-table, and flagger-table)
```
docker-compose exec broker kafka-topics --create --bootstrap-server \
localhost:9092 --replication-factor 1 --partitions 1 --topic deposit

 docker-compose exec broker kafka-topics --create --bootstrap-server \
localhost:9092 --replication-factor 1 --partitions 1 --topic flag-wallet

 docker-compose exec broker kafka-topics --create --bootstrap-server \
localhost:9092 --replication-factor 1 --partitions 1 --topic balance-table

 docker-compose exec broker kafka-topics --create --bootstrap-server \
localhost:9092 --replication-factor 1 --partitions 1 --topic threshold-table

 docker-compose exec broker kafka-topics --create --bootstrap-server \
localhost:9092 --replication-factor 1 --partitions 1 --topic flagger-table
```
Or simply run `make topic` command.

### How to run
1. Create a wallet, using command:
```
go run cmd/flag-wallet/main.go -wallet 4n6UziJObz7zWzPY
```
2. Run service (handlers, emitter, and view) with 
```
make dev
```
3. Run processor in other terminal with
```
make processor
```

Internally the Go Program will start three Goka processors.

## Directory structure

```
├── cmd
│   └── flag-wallet
│       ├── main.go
│   └── processor
│       ├── main.go
│   └── service
│       ├── main.go
├── deposit
│   └── deposit.go
│   └── deposit_list.go
├── flagger
│   └── history.go
├── history
│   └── deposit.go
│   └── deposit_list.go
├── proto
│   └── counter.proto
│   └── deposit.proto
│   └── flag.proto
│   └── pb
│       ├── counter.pb.go
│       ├── deposit.pb.go
│       ├── flag.pb.go
├── service
│   └── handler.go
│   └── response.go
│   └── service.go
│   └── type.go
├── threshold
│   └── codec.go
│   └── threshold.go
├── docker-compose.yml
├── Makefile
├── README.md
```

## API Reference

#### Create a deposit

```http
  POST /deposit
```
Request body:
| Field           | Type     | Description                     |
| :-------------- | :------- | :------------------------------ |
| `wallet_id`     | `string` | **Required**. wallet identifier |
| `amount`        | `float64`| **Required**. deposit amount    |

Examples:
```bash
curl -X POST "localhost:8080/deposit" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"wallet_id":"4n6UziJObz7zWzPY","amount":6000000}'
```

#### Get wallet detail

```http
  GET /wallet/{id}
```

Examples:
```bash
curl -X GET "localhost:8080/wallet/4n6UziJObz7zWzPY" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \

Result
{
  "wallet_id": "4n6UziJObz7zWzPY",
  "balance": 12000,
  "above_threshold": true
}
```

## Authors

- [@hutamy](https://www.github.com/hutamy)