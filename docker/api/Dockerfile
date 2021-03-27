
FROM golang:alpine AS build
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /dist/api ./pkg/ports/api/main.go

FROM alpine AS bin
WORKDIR /app
COPY --from=build /dist/api /app
COPY ./pkg/ports/api/static /app/static
ENV GIN_MODE=release
EXPOSE 8000 8001 8002 8004
ENTRYPOINT ["/app/api"]
