# Goforge

Forge production-ready Go applications with security, observability, and best practices built-in.

## 🚀 Installation

### Option 1: Go Install (Recommended)

```bash
go install github.com/viveksharma/goforge@latest
```

### Option 2: Download Binary

Download pre-built binaries from the [Releases](https://github.com/viveksharma/goforge/releases) page.

### Option 3: Build from Source

```bash
git clone https://github.com/viveksharma/goforge.git
cd goforge
go build -o goforge ./cmd/goforge
sudo mv goforge /usr/local/bin/
```

## 📖 Usage

Create a new project:

```bash
goforge create my-awesome-api
cd my-awesome-api
make up
```

Visit `http://localhost:8080/health/ready` to verify the API is running.

## ✨ What You Get

Every generated project includes:

### 🔒 Security First

- **Security Headers**: HSTS, CSP, X-Frame-Options, X-Content-Type-Options
- **Input Validation**: Path traversal protection, request validation
- **No Sensitive Logging**: Credentials never appear in logs
- **Panic Recovery**: Graceful error handling
- **Rate Limiting Ready**: Redis-backed rate limiting structure
- **Secure Defaults**: Non-root Docker user, read-only filesystem
- **Timeouts**: Request/connection timeouts prevent DoS

### 📊 Observability

- **Structured Logging**: JSON logs with correlation IDs (zap)
- **Health Checks**: Kubernetes-ready `/health/live` and `/health/ready`
- **Request Tracing**: Unique request IDs for log correlation
- **Error Tracking**: Contextual error logging with stack traces

### 🏗️ Production Ready

- **Fiber Web Framework**: High-performance HTTP server
- **PostgreSQL**: Production-grade database with connection pooling
- **Redis**: Caching layer with connection management
- **Docker Compose**: Zero-config local development
- **Graceful Shutdown**: Connection draining on SIGTERM/SIGINT
- **Environment Management**: Type-safe `.env` configuration

### 🔧 Developer Experience

- **Make Commands**: Common tasks via Makefile
- **Hot Reload Ready**: Easy integration with Air or CompileDaemon
- **Clean Architecture**: Separation of concerns (handler/service/repository pattern ready)
- **Dockerfile Included**: Multi-stage build with security best practices
- **Comprehensive README**: Documentation generated with every project

## 📦 Project Structure

```
your-project/
├── cmd/api/              # Application entry point
├── internal/
│   ├── config/           # Environment configuration
│   ├── handler/          # HTTP request handlers
│   ├── middleware/       # HTTP middleware (security, logging, recovery)
│   └── server/           # Server setup and routing
├── pkg/
│   ├── logger/           # Structured logging
│   ├── database/         # PostgreSQL client
│   └── cache/            # Redis client
├── deployments/
│   └── Dockerfile        # Multi-stage production build
├── docker-compose.yml    # Local development stack
├── Makefile              # Common commands
└── .env.example          # Environment variables template
```

## 🔒 Security Features

### Built-in Protections

1. **Path Traversal Prevention**: Project name validation prevents `../` attacks
2. **SQL Injection Protection**: Parameterized queries enforced
3. **XSS Protection**: Security headers set by default
4. **Clickjacking Protection**: X-Frame-Options: DENY
5. **MIME Sniffing Protection**: X-Content-Type-Options: nosniff
6. **Request Timeout Protection**: Read/write timeouts configured
7. **Body Size Limits**: 4MB default limit
8. **Non-Root Docker User**: Containers run as user 1000
9. **Read-Only Filesystem**: Docker containers use read-only root
10. **Secrets Management**: `.env` files never committed

### Security Best Practices

All generated code follows:

- OWASP Top 10 protection
- Principle of least privilege
- Defense in depth
- Secure defaults
- Input validation
- Output encoding (where applicable)

## 📚 Generated Project Commands

```bash
make up          # Start all services (API, PostgreSQL, Redis)
make down        # Stop all services
make logs        # View logs
make build       # Build the Go binary
make run         # Run without Docker
make test        # Run tests
make fmt         # Format code
make lint        # Run linter
make clean       # Clean up
```

## 🛠️ Customization

Generated projects are fully customizable:

1. Add your business logic in `internal/handler/`
2. Create services in `internal/service/`
3. Add repositories in `internal/repository/`
4. Update routes in `internal/server/server.go`
5. Modify environment variables in `.env`

## 🔐 Environment Variables

Every project includes these environment variables:

| Variable       | Description                           | Default       |
| -------------- | ------------------------------------- | ------------- |
| `APP_ENV`      | Environment (development/production)  | `development` |
| `APP_PORT`     | HTTP server port                      | `8080`        |
| `DATABASE_URL` | PostgreSQL connection string          | Required      |
| `REDIS_URL`    | Redis connection string               | Required      |
| `LOG_LEVEL`    | Logging level (debug/info/warn/error) | `info`        |

## 🚢 Deployment

Generated projects are deployment-ready for:

- **Kubernetes**: Health check endpoints configured
- **Docker**: Multi-stage Dockerfile included
- **Cloud Run**: Listens on PORT environment variable
- **AWS ECS/Fargate**: 12-factor app compliant
- **Any platform**: Standard REST API

## 🧪 Testing

Every generated project is test-ready:

```bash
cd my-project
make test
```

Add your tests in:

- `internal/handler/*_test.go`
- `internal/service/*_test.go`
- `pkg/*_test.go`

## 📖 Examples

### Create a basic API

```bash
goforge create my-api
```

### What happens:

1. ✅ Creates project directory
2. ✅ Generates all project files
3. ✅ Sets up Fiber server
4. ✅ Configures PostgreSQL and Redis
5. ✅ Adds health check endpoints
6. ✅ Includes Docker Compose
7. ✅ Creates comprehensive README

### Start developing immediately

```bash
cd my-api
make up
curl http://localhost:8080/health/ready
```

## 🤝 Contributing

Contributions welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Add tests for new features
4. Ensure all tests pass
5. Submit a pull request

## 📝 License

MIT License - feel free to use this for any project.

## 🐛 Issues

Found a bug or have a feature request? Open an issue on GitHub.

## ⭐ Show Your Support

If this tool helped you, give it a star on GitHub!

---

**Made with ❤️ for the Go community**
