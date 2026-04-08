FROM golang:1.26-alpine AS builder
WORKDIR /build
COPY go.work ./
COPY services/gateway/go.mod services/gateway/go.sum ./services/gateway/
COPY services/achievement/go.mod services/achievement/go.sum ./services/achievement/
RUN cd services/gateway && go mod download
COPY services/ ./services/
RUN cd services/gateway && CGO_ENABLED=0 go build -o /app .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
