
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


# ref
- flagの説明
    - https://golang.org/cmd/go/#hdr-Testing_flags
- `go tool trace`の説明
    - https://qiita.com/ma91n/items/caa82d84c8d3ac07cfe0
- benchstat
    - https://pkg.go.dev/golang.org/x/perf/cmd/benchstat
- worker poolの参考実装
    - https://qiita.com/tutuz/items/798361dc78412da2bba6
- 1x指定の方法
    - https://qiita.com/syossan27/items/148e33dd9da4ee3dc89b
- フィボナッチ
    - https://medium.com/eureka-engineering/go%E8%A8%80%E8%AA%9E%E3%81%AE%E3%83%99%E3%83%B3%E3%83%81%E3%83%9E%E3%83%BC%E3%82%AF%E3%81%A7%E3%83%91%E3%83%95%E3%82%A9%E3%83%BC%E3%83%9E%E3%83%B3%E3%82%B9%E6%B8%AC%E5%AE%9A-6c72832a7cb5
- 同時実行数の制御（やや語弊あり）
    - https://qiita.com/ReSTARTR/items/ee943512243aedb3aa25
