mock:
	mockgen -source=./internal/usecase/interfaces.go -destination=./internal/usecase/mock_test.go -package=usecase_test

test:
	go test -v -cover -race ./internal/...