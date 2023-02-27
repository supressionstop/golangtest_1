.PHONY: run
run:
	docker-compose up processor postgres kiddy-provider -d

.PHONY: stop
stop:
	docker-compose stop processor postgres kiddy-provider

.PHONY: test
test: test-unit test-integration


.PHONY: lint
lint:
	golangci-lint run


# Helpers

.PHONY: mock
mock: ### update test mock files from interfaces
	mockgen -source=./internal/usecase/interfaces.go -destination=./internal/usecase/mock_test.go -package=usecase_test

.PHONY: migrate-create
bin-deps: ### installs helper binaries
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@latest
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.1

.PHONY: migrate-create
migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'

.PHONY: test-unit
test-unit: ### only unit tests
	go test -v -cover -race ./internal/...

.PHONY: test-integration
test-integration: ### up integration env and run tests
	docker-compose -f docker-compose.test.yaml up -d --build
	docker-compose -f docker-compose.test.yaml run test-integration

.PHONY: swagger
swagger: ### update files in api/swagger/
	swag init -g ./internal/controller/http/v1/router.go -o ./api/swagger

.PHONY: rebuild
rebuild: ### only unit tests
	docker-compose stop
	docker-compose build
	docker-compose up processor postgres kiddy-provider -d

.PHONY: proto
proto:
	protoc --go_out=internal/controller/grpc/v1 --go-grpc_out=internal/controller/grpc/v1 ./api/protobuf/processor.proto
