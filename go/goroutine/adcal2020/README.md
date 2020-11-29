
```bash
$ go test -bench . -benchmem -count 10 -memprofile mem.out
```

```bash
$ go test -bench .*Simple* -benchmem -count 2 -trace simple.trace
$ go test -bench .*Smart* -benchmem -count 2 -trace smart.trace
```

```bash
$ go tool trace --http localhost:6060 smart.trace
$ go tool trace --http localhost:6061 simple.trace
```