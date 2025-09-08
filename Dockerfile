# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.25 AS builder
WORKDIR /app

# mod dosyaları
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

# kaynak kod
COPY . .

# main paketi cmd altında
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/backend ./cmd

# ---- Runtime stage (küçük ve güvenli) ----
FROM gcr.io/distroless/static-debian12
COPY --from=builder /bin/backend /backend
EXPOSE 8080
ENTRYPOINT ["/backend"]
