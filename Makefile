generate:
	protoc -I=. -I=${GOPATH}/src  --gogoslick_out=./pkg/application/message user.proto

swagger:
	cd ./pkg/infrastructure/ports/api
	swag init --parseDependency=true

migrate:
	docker build --tag migration:1.0 -f docker/migrate/Dockerfile .
	docker run --rm  --name migration --network host migration:1.0

clean:
	cd ./db
	flyway clean

prod:
	./scripts/deploy.sh pro

dev:
	./scripts/deploy.sh dev

local:
	docker-compose -f docker-compose.yml up -d
