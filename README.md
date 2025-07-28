# HTML Hello World Docker 🐳

A simple static HTML page served by a Go web server, containerized with Docker and ready for Docker Hub deployment.

## 🚀 Features

- **Go Web Server**: Lightweight HTTP server written in Go
- **Static HTML**: Beautiful, responsive HTML page with modern CSS
- **Docker Ready**: Optimized multi-stage Dockerfile for production
- **Health Checks**: Built-in health check endpoints
- **Logging**: Request logging middleware
- **Security**: Non-root user execution in container

## 📋 Prerequisites

- Docker installed on your system
- Go 1.21+ (for local development)
- Git (for version control)

## 🏃‍♂️ Quick Start

### Running with Docker

1. **Build the Docker image:**
   ```bash
   docker build -t html-hello-world .
   ```

2. **Run the container:**
   ```bash
   docker run -p 8080:8080 html-hello-world
   ```

3. **Open your browser and visit:**
   - Main page: http://localhost:8080
   - Health check: http://localhost:8080/health
   - App info: http://localhost:8080/info

### Running locally (without Docker)

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

## 🐳 Docker Hub Deployment

### Building and pushing to Docker Hub

1. **Tag your image:**
   ```bash
   docker tag html-hello-world your-dockerhub-username/html-hello-world:latest
   ```

2. **Push to Docker Hub:**
   ```bash
   docker push your-dockerhub-username/html-hello-world:latest
   ```

3. **Run from Docker Hub:**
   ```bash
   docker run -p 8080:8080 your-dockerhub-username/html-hello-world:latest
   ```

## 📁 Project Structure

```
.
├── Dockerfile          # Multi-stage Docker configuration
├── main.go            # Go web server
├── index.html         # Static HTML page
├── go.mod            # Go module definition
└── README.md         # This file
```

## 🛠️ Development

### Environment Variables

- `PORT`: Server port (default: 8080)

### Endpoints

- `/` - Main HTML page
- `/health` - Health check endpoint (JSON response)
- `/info` - Application information (JSON response)

## 🔧 Docker Configuration

The Dockerfile uses a multi-stage build process:

1. **Build Stage**: Uses `golang:1.21-alpine` to compile the Go application
2. **Production Stage**: Uses `alpine:latest` for a minimal runtime environment

### Security Features

- Runs as non-root user
- Minimal Alpine Linux base image
- Health checks included
- Only necessary files copied to final image

## 📊 Health Monitoring

The application includes built-in health checks:

- **Docker Health Check**: Automatically checks `/health` endpoint every 30 seconds
- **Manual Health Check**: Visit `/health` for application status
- **Application Info**: Visit `/info` for detailed application information

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test with Docker
5. Submit a pull request

## 📝 License

This project is open source and available under the MIT License.

## 🎯 Use Cases

- Learning Docker containerization
- Go web server examples
- Static site deployment
- Docker Hub publishing practice
- Kubernetes deployment templates

---

**Made with ❤️ using Go, Docker, and modern web technologies** 