# Testing
---

- simple `go test`
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ go test
```
```
PASS
ok  	github.com/ipatser/vcs_ipatser	0.798s
```

- `go test` with additional information
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ go test -v
```
```
=== RUN   TestMoviesInitial
--- PASS: TestMoviesInitial (0.06s)
=== RUN   TestMoviesAdd
--- PASS: TestMoviesAdd (0.32s)
=== RUN   TestMoviesDelByID
--- PASS: TestMoviesDelByID (0.20s)
=== RUN   TestMoviesDelAll
--- PASS: TestMoviesDelAll (0.13s)
PASS
ok  	github.com/ipatser/vcs_ipatser	0.720s
```

- to get code coverage of the testcases in a file
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ go test -coverprofile=coverage.out
```
```
PASS
coverage: 0.0% of statements
```

- content of the `coverage.out` file
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ cat coverage.out 
```
```
mode: set
github.com/ipatser/vcs_ipatser/main.go:13.13,39.2 10 0
```

- to display code coverage using html
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ go tool cover -html=coverage.out
```

![vcs_ipatser.coverage.png](/static/vcs_ipatser.coverage.png)

- to get benchmark
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ go test -bench=.
```
```
goos: linux
goarch: amd64
pkg: github.com/ipatser/vcs_ipatser
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMoviesAdd-2      	1000000000	         0.2775 ns/op
BenchmarkMoviesDelAll-2   	1000000000	         0.06414 ns/op
PASS
ok  	github.com/ipatser/vcs_ipatser	5.729s
```

- to run only specific benchmark function
```bash
~/workspace/src/github.com/ipatser/vcs_ipatser$ go test -bench=MoviesDelAll
```
```
goos: linux
goarch: amd64
pkg: github.com/ipatser/vcs_ipatser
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMoviesDelAll-2   	1000000000	         0.06850 ns/op
PASS
ok  	github.com/ipatser/vcs_ipatser	5.450s
```
