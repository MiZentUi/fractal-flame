FROM golang:1.26.3-alpine3.22 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go run -o iam-service ./cmd/iam-service/main.go

FROM alpine:3.22

WORKDIR /build

COPY --from=builder /app/iam-service ./

RUN adduser -S appgroup && \
    adduser -S appuser -G appgroup && \
    chown appuser:appgroup ./iam-service

USER appuser

EXPOSE 50051

ENTRYPOINT [ "./iam-service" ]
