# ---------- Stage 1: Build ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependencies
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello .

# ---------- Stage 2: Run ----------
FROM scratch

# Copy CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary and server certs
COPY --from=builder /app/hello /hello
COPY --from=builder /app/server.crt /server.crt
COPY --from=builder /app/server.key /server.key

# Expose HTTPS port
EXPOSE 8443

ENTRYPOINT ["/hello"]
