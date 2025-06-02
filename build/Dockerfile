FROM golang:1.24 as builder

WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
# We only need cmd and internal for the build
COPY cmd cmd/
COPY internal internal/

# Build the application
# Output the binary to /app/gocn-server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/gocn-server ./cmd/gocn-server/main.go

# --- Final Stage ---
FROM alpine:latest

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/gocn-server /usr/local/bin/gocn-server

# Set ownership for the copied binary and workdir
RUN chown appuser:appgroup /usr/local/bin/gocn-server

# Switch to the non-root user
USER appuser

# Expose port if your application listens on one (e.g., for a web server)
# EXPOSE 8080

# Run the application
ENTRYPOINT ["gocn-server"] 