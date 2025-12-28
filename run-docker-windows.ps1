# PowerShell script to run the Game of Life Docker container on Windows using WSL2

Write-Host "Starting Conway's Game of Life in Docker..." -ForegroundColor Green

# Check if Docker is running
$dockerRunning = docker info 2>&1 | Select-String "Server Version"
if (!$dockerRunning) {
    Write-Host "Error: Docker is not running. Please start Docker Desktop." -ForegroundColor Red
    exit 1
}

# Check if WSL is available
$wslCheck = wsl --list --quiet 2>&1
if ($LASTEXITCODE -ne 0) {
    Write-Host "Error: WSL is not installed. Please install WSL2." -ForegroundColor Red
    exit 1
}

# Pull the latest image
Write-Host "Pulling latest image..." -ForegroundColor Cyan
docker pull sondreevik/go-game-of-life:latest

# Run the container through WSL to get proper display support
Write-Host "Running Game of Life through WSL..." -ForegroundColor Cyan
wsl -e bash -c 'xhost +local:docker 2>/dev/null; docker run --rm -e DISPLAY=:0 -v /tmp/.X11-unix:/tmp/.X11-unix:rw -v /mnt/wslg:/mnt/wslg --user $(id -u):$(id -g) sondreevik/go-game-of-life:latest; xhost -local:docker 2>/dev/null'

Write-Host "Game closed." -ForegroundColor Green
