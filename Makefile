dep:
	go mod tidy
	go mod vendor

up:
	docker-compose up -d

down:
	docker-compose down

dev:
	go run cmd/service/main.go

processor:
	go run cmd/processor/main.go -history -flagger -threshold

topic:
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

