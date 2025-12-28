# syntax=docker/dockerfile:1
# --- Builder Stage ---
# Use the native platform for building (not cross-compiling)
FROM --platform=$TARGETPLATFORM golang:alpine AS builder

# Install necessary build dependencies for Ebiten (requires CGO and X11)
RUN apk add --no-cache gcc musl-dev libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev mesa-dev

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application with CGO enabled for Ebiten (native build)
RUN CGO_ENABLED=1 go build -o app .

# --- Runner Stage ---
# Use alpine instead of scratch since we need C libraries for Ebiten
FROM alpine:latest

# Install runtime dependencies for X11 and OpenGL
RUN apk add --no-cache libx11 libxrandr libxcursor libxinerama libxi mesa-gl

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app /app

# Set the entry point command to run the application
CMD ["/app"]
