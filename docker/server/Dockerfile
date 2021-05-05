
FROM golang:alpine AS build
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /build

RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.4/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /dist/server ./pkg/ports/server/main.go

FROM alpine AS bin
COPY --from=build /bin/grpc_health_probe ./grpc_health_probe

WORKDIR /app
COPY --from=build /dist/server /app
COPY env /app/env

EXPOSE 8090
ENTRYPOINT ["/app/server"]

