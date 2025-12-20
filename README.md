# Go Training Program: Path to HTTP Forward Proxy

**Target**: Build expertise for implementing a high-performance HTTP forward proxy
**Time Investment**: ~10 hours across 3 steps
**Prerequisites**: Completed Go Tour, experienced in Python/React

---

## Step 1: Idiomatic Go Fundamentals (3.5 hours)

### Materials
- [Effective Go](https://go.dev/doc/effective_go) - Skim sections: Formatting, Commentary, Names, Control structures, Functions, Data, Methods, Interfaces, Concurrency (30 min)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) - Quick read of common pitfalls (15 min)

### Exercise 1.1: Error Handling & Custom Errors (45 min)
Build a simple configuration parser that demonstrates idiomatic error handling.

**Task**: Create a package that reads a JSON config file and validates required fields.
- Implement custom error types using `errors.New()` and `fmt.Errorf()` with `%w`
- Use error wrapping and `errors.Is()` / `errors.As()`
- Return sentinel errors for specific cases (e.g., `ErrMissingField`)
- Write a CLI tool that uses this package and handles errors gracefully

**Focus**: Error wrapping, sentinel errors, error context

---

### Exercise 1.2: Interfaces & Composition (60 min)
Create a flexible logging system using interfaces and composition (relevant for proxy logging).

**Task**: Build a logging package with multiple outputs
- Define a `Logger` interface with methods: `Info()`, `Error()`, `Debug()`
- Implement 3 concrete types: `ConsoleLogger`, `FileLogger`, `MultiLogger`
- `MultiLogger` should compose multiple loggers
- Add log levels using constants/enums (iota)
- Demonstrate interface satisfaction implicitly (no "implements" keyword)

**Focus**: Small interfaces, composition over inheritance, implicit satisfaction

---

### Exercise 1.3: Goroutines & Channels (90 min)
Build a concurrent worker pool (essential pattern for proxy request handling).

**Task**: Implement a generic worker pool that processes jobs concurrently
- Create a `WorkerPool` struct with configurable worker count
- Use channels for job distribution and result collection
- Implement graceful shutdown using context
- Add a simple job: simulate HTTP request processing with random delays
- Demonstrate buffered vs unbuffered channels
- Handle panics in workers (recover pattern)

**Focus**: Channel patterns, goroutine lifecycle, context for cancellation, select statements

---

### Exercise 1.4: Struct Design & Methods (30 min)
Practice idiomatic struct design patterns.

**Task**: Design a `ConnectionPool` struct for managing connections
- Use pointer vs value receivers appropriately
- Implement constructor pattern (`NewConnectionPool()`)
- Add methods: `Get()`, `Put()`, `Close()`
- Use sync.Mutex for thread-safety
- Demonstrate zero values working sensibly where possible

**Focus**: Constructor pattern, pointer vs value receivers, zero values, embedding

---

## Step 2: Testing & Tooling (2.5 hours)

### Materials
- [Go Testing Package](https://pkg.go.dev/testing) - Read the overview (10 min)
- [Table-driven tests](https://go.dev/wiki/TableDrivenTests) - Pattern reference (10 min)
- [Go tooling guide](https://go.dev/doc/toolchain) - Skim (10 min)

### Exercise 2.1: Unit Testing Best Practices (60 min)
Write comprehensive tests for your Exercise 1.2 logging package.

**Task**: Add tests demonstrating multiple patterns
- Table-driven tests for different log levels
- Test helpers using `t.Helper()`
- Subtests using `t.Run()`
- Use `testdata/` directory for fixtures
- Mock the `Logger` interface for testing consumers
- Benchmark test for logger performance (`Benchmark*`)

**Focus**: Table-driven tests, subtests, test helpers, benchmarks

---

### Exercise 2.2: Tooling Setup (45 min)
Set up a professional Go project structure.

**Task**: Configure your project with industry-standard tools
- Initialize with `go mod init`
- Add `.gitignore` for Go
- Configure `golangci-lint` with recommended linters
  - Install: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
  - Create `.golangci.yml` config enabling: errcheck, govet, staticcheck, unused
- Set up `gofmt` and `goimports`
- Create a `Makefile` with targets: `test`, `lint`, `fmt`, `build`
- Add GitHub Actions / GitLab CI config for automated checks (optional but recommended)

**Resources**:
- [golangci-lint](https://golangci-lint.run/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

**Focus**: Project structure, automation, code quality tools

---

### Exercise 2.3: Integration Testing (35 min)
Write integration tests using httptest.

**Task**: Create a simple HTTP server and test it
- Build a basic HTTP server with 2-3 endpoints
- Use `httptest.NewServer()` for integration tests
- Test request/response cycles
- Use `net/http/httptest.ResponseRecorder` for handler testing
- Add test coverage reporting: `go test -cover`

**Focus**: httptest package, handler testing, test coverage

---

## Step 3: Essential Libraries (4 hours)

### Exercise 3.1: HTTP Server Fundamentals (75 min)
Deep dive into `net/http` and routing patterns.

**Materials**:
- [net/http docs](https://pkg.go.dev/net/http) - Focus on Handler, ServeMux, Client (20 min)

**Task**: Build a feature-rich HTTP server
- Create a server with multiple routes using `http.ServeMux`
- Implement middleware pattern (logging, auth, CORS)
- Add request timeout handling
- Implement graceful shutdown using signals
- Parse query params, path params, and JSON bodies
- Make HTTP client requests with timeouts and retries

**Focus**: Handler interface, middleware chaining, context timeouts, graceful shutdown

---

### Exercise 3.2: Advanced HTTP with Chi Router (60 min)
Use `chi` router for production-grade routing (lightweight, idiomatically Go).

**Materials**:
- [chi router](https://github.com/go-chi/chi) - Read README and examples (15 min)

**Task**: Refactor Exercise 3.1 using chi
- Install: `go get -u github.com/go-chi/chi/v5`
- Use chi for route parameters: `/users/{id}`
- Implement route groups and sub-routers
- Add chi's built-in middleware (Logger, Recoverer, Timeout)
- Create custom middleware compatible with chi

**Why chi**: Standard library compatible, no dependencies, fast, idiomatic

---

### Exercise 3.3: Context & Cancellation (45 min)
Master context package (critical for proxy request lifecycle).

**Task**: Build request processing with proper context usage
- Implement a request handler that makes downstream HTTP calls
- Use `context.WithTimeout()` for request deadlines
- Propagate context through function calls
- Use `context.WithValue()` for request-scoped data (request ID)
- Handle context cancellation mid-flight
- Demonstrate context cancellation propagating to goroutines

**Focus**: Context propagation, timeouts, cancellation, values

---

### Exercise 3.4: Database Integration with pgx (60 min)
Learn connection pooling and query patterns with PostgreSQL.

**Materials**:
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver (15 min)

**Task**: Create a simple API backed by PostgreSQL
- Install: `go get github.com/jackc/pgx/v5`
- Set up connection pool with `pgxpool`
- Implement CRUD operations using `Query()`, `QueryRow()`, `Exec()`
- Use prepared statements
- Handle SQL NULL values properly
- Add context timeouts to queries
- Use Docker Compose to run Postgres locally

**Why pgx**: Best-in-class PostgreSQL driver, excellent performance, idiomatic Go

---

### Exercise 3.5: Key Utilities (40 min)
Quick exposure to essential utility libraries.

**Task**: Small exercises with each library
1. **Structured Logging with slog** (stdlib)
   - Replace your custom logger with `log/slog`
   - Configure JSON output
   - Add structured fields to logs

2. **Configuration with viper**
   - Install: `go get github.com/spf13/viper`
   - Read config from file and environment variables
   - Demonstrate precedence

3. **Goroutine Management with errgroup**
   - Use `golang.org/x/sync/errgroup`
   - Run parallel tasks with error handling
   - Replace your worker pool with errgroup

**Resources**:
- [slog package](https://pkg.go.dev/log/slog)
- [viper](https://github.com/spf13/viper)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)

---

## Next Steps: Toward the Proxy

After completing these exercises, you'll be ready for:
1. **HTTP/2 deep dive**: Study `golang.org/x/net/http2`
2. **TLS/SSL**: Learn `crypto/tls` package
3. **WebSocket proxying**: `github.com/gorilla/websocket`
4. **Reverse proxy patterns**: Study `net/http/httputil.ReverseProxy`
5. **Performance optimization**: Profiling with pprof, benchmarking

**Recommended Next Project**: Build a simple reverse proxy before tackling the forward proxy. It's an easier starting point that covers 70% of the same concepts.

---

## Validation Checklist

After completing this program, you should be able to:
- [ ] Write idiomatic Go with proper error handling
- [ ] Design interfaces and use composition effectively
- [ ] Manage goroutines and channels safely
- [ ] Write table-driven tests with good coverage
- [ ] Set up professional Go project tooling
- [ ] Build production-grade HTTP servers with middleware
- [ ] Use context for cancellation and timeouts
- [ ] Work with databases using connection pooling
- [ ] Understand standard library networking primitives

**Estimated Time**: 10 hours (3.5 + 2.5 + 4)
