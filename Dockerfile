# ---------- STAGE 1: Build ----------
FROM golang:1.23-alpine AS builder

# Install git and build tools
RUN apk add --no-cache git

# I don't think this is right
WORKDIR /app 

# Copy go.mod and go.sum first 
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build Binary (a lot to unpack here. See what all this means)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server

# ---------- STAGE 2: Run ----------
FROM alpine:latest

WORKDIR /app

# Copy only binary + any static files (no source code)
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080

# Default command
CMD ["./main"]