FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -a -ldflags="-w -s" \
    -o kagevault \
    cmd/api/main.go

FROM debian:bookworm-slim AS runtime

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/kagevault .

EXPOSE 3000
CMD ["./kagevault"]

FROM scratch AS binary
COPY --from=builder /app/kagevault /
