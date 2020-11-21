
FROM golang:alpine AS build
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /dist/server ./pkg/infrastructure/ports/server/main.go

FROM alpine AS bin
WORKDIR /app
COPY --from=build /dist/server /app
COPY env /app/env
EXPOSE 9001
ENTRYPOINT ["/app/server"]

