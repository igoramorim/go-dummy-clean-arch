## What is it?
This is a dummy project used to practice and understand Clean Architecture with Go.

## TODOs
- Wrap errors adding context
- Test / understand router.Use()
- Add GraphQL
- Unit tests

## How to run

1. Start docker

2. Start Kafka and MySQL:
```
docker-compose up -d
```

3. Create the topic:
```
./scripts/kafka_create_topic.sh
```

After this you can choose how to run the app

If you want to serve and rest API, run build config **API** in **IntelliJ** or
```
go run cmd/api/main.go
```

---

If you want to run it via CLI, use config **CLI** or
```
go run cmd/cli/main.go
```

---

If you want to run it as a consumer, use config **Consumer** or
```
go run cmd/queue/main.go
```

and produce some messages running the build config **Producer** in IntelliJ or
```
go run cmd/queue/producer/kafka_producer.go
```