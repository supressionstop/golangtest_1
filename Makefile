.PHONY: run
run:
	docker-compose up processor postgres kiddy-provider -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: test
test:
	go test -v -cover -race ./internal/...

# Helpers

mock:
	mockgen -source=./internal/usecase/interfaces.go -destination=./internal/usecase/mock_test.go -package=usecase_test

bin-deps:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@latest

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'migrate_name'
.PHONY: migrate-create

test-integration:
	docker-compose up integration-test

rebuild:
	docker-compose down
	docker volume rm softpro6_pg-data
	docker-compose build
	docker-compose up
