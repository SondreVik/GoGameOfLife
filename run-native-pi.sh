#!/bin/bash

# Script to run Conway's Game of Life natively on Raspberry Pi

echo "Starting Conway's Game of Life..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed."
    echo "Install Go with: sudo apt install golang"
    exit 1
fi

# Install required development libraries for Ebiten
echo "Checking dependencies..."
sudo apt-get update
sudo apt-get install -y libc6-dev libglu1-mesa-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config

# Download Go dependencies
echo "Downloading dependencies..."
go mod download

# Build the application
echo "Building application..."
go build -tags=ebitengl -o game-of-life .

# Run the application with legacy OpenGL (Pi compatibility)
echo "Running Game of Life..."
echo "Using legacy OpenGL for Raspberry Pi compatibility..."
./game-of-life

echo "Game closed."
