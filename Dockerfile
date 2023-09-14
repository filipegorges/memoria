# Start with the official Go image.
FROM golang:1.17 AS builder

# Install Tesseract and its dependencies.
RUN apt-get update && \
    apt-get install -y \
    tesseract-ocr \
    libtesseract-dev \
    libleptonica-dev

# Set up the Go workspace.
WORKDIR /app

# Copy go.mod and go.sum to download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire directory contents (assuming this Dockerfile is at the root of your Go project).
COPY . .

# Build the Go application.
RUN go build -o myapp

# Start a new, smaller base image for the final stage.
FROM debian:buster-slim

# Copy Tesseract data files and libraries from the builder image.
COPY --from=builder /usr/share/tesseract-ocr/4.00/tessdata /usr/share/tesseract-ocr/4.00/tessdata
COPY --from=builder /usr/lib/x86_64-linux-gnu/libtesseract.so.4 /usr/lib/x86_64-linux-gnu/
COPY --from=builder /usr/lib/x86_64-linux-gnu/liblept.so.5 /usr/lib/x86_64-linux-gnu/

# Install runtime dependencies.
RUN apt-get update && \
    apt-get install -y \
    libtesseract4 \
    libleptonica3 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy the Go binary from the builder stage.
COPY --from=builder /app/myapp /app/myapp

# Run the Go application.
CMD ["/app/myapp"]
