# Stage 1: Build the Go application
FROM golang:1.24.5-alpine AS builder
WORKDIR /app

COPY . .



ARG BUILD_TARGET 
RUN go build -o /app/bin/app $BUILD_TARGET

# Stage 2: Create a smaller image for running the binary
FROM alpine:latest
#FROM scratch

COPY --from=builder /app/bin/app /app

ENTRYPOINT ["/app"]