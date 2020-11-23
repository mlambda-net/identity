.PHONY:  build test lint swagger

generate:
	protoc -I=. -I=${GOPATH}/src  --gogoslick_out=./pkg/application/message user.proto

swagger:
	cd ./pkg/infrastructure/ports/api && swag init --parseDependency=true

migrate:
	docker build --tag migration:1.0 -f docker/migrate/Dockerfile .
	docker run --rm  --name migration --network host migration:1.0

build:
	go build -o ./dist/server ./pkg/infrastructure/ports/server/main.go
	go build -o ./dist/api  ./pkg/infrastructure/ports/api/main.go
lint:
	golangci-lint run

test:
	go test ./... -v

clean:
	cd ./db
	flyway clean

prod:
	./scripts/deploy.sh prod

dev:
	./scripts/deploy.sh dev

local:
	docker-compose -f docker-compose.yml up -d
