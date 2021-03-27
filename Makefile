.PHONY:  build test lint swagger

generate:
	protoc -I=. -I=${GOPATH}/src  --gogoslick_out=./pkg/application/message user.proto
	protoc -I=. -I=${GOPATH}/src  --gogoslick_out=./pkg/application/message app.proto
	protoc -I=. -I=${GOPATH}/src  --gogoslick_out=./pkg/application/message rol.proto

swagger:
	cd ./pkg/ports/api && swag init --parseDependency=true

migrate:
	docker-compose -f docker-compose.yml up -d --build migrate

build:
	go build -o ./dist/server ./pkg/ports/server/main.go
	go build -o ./dist/api  ./pkg/ports/api/main.go
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
