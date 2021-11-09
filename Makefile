run:
	docker-compose up -d

stop:
	docker-compose down

local:
	go run ./cmd/api/main.go

build:
	docker build -t urlshortner .

test:
	go test ./*/*