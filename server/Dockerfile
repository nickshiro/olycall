FROM golang:1.24.2-alpine3.20 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags "-s -w" -o /bin/appbin ./cmd/server

FROM gcr.io/distroless/static:nonroot

USER nonroot:nonroot

WORKDIR /app

COPY --from=builder /bin/appbin ./appbin
COPY --from=builder /src/migrations/domain ./migrations

ENTRYPOINT ["./appbin"]
