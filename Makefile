go_tidy:
	docker compose run --rm app go mod tidy

go_fmt:
	docker compose run --rm app go fmt ./...

go_vet:
	docker compose run --rm app go vet ./...
