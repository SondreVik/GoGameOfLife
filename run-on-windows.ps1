# PowerShell script to run the Game of Life on Windows

Write-Host "Starting Conway's Game of Life..." -ForegroundColor Green

# Check if Go is installed
if (!(Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Error: Go is not installed or not in PATH" -ForegroundColor Red
    exit 1
}

# Run the Go application directly (native Windows execution)
go run .
