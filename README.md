# im-distributed-**cache**
### Another in-memory distributed cache. Why? Because...why not?

**Another in-memory distributed cache built for learning purposes** but feel free to use it on your projects if you think it can be useful.

The aim of it is to explore: 
- Go concurrency model
- Inter-communication between multiple servers (nodes)
- Add/remove server (node) from a cluster and re-balance the load
- Hashing mechanism to distribute the data between the nodes
- Benchmarking and performance optimization strategies

### Usage

*Work in progress*

*See [example/main.go](example/main.go) for usage*

### Testing

```bash
make test
```

### Benchmarking

```bash
make bench
```

#### Results
```bash
❯ make bench
go test -bench=. ./...
goos: darwin
goarch: amd64
pkg: github.com/claudiunicolaa/im-distributed-cache
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkCache_Get-8     6630663               189.7 ns/op
PASS
ok      github.com/claudiunicolaa/im-distributed-cache  6.631s
```
