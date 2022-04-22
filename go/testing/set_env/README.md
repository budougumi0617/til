# Setenv
Setenvは平行に動いている別のパッケージのテストには影響しない。
Cleanupするので別のテストケースにも影響しない。サブテストの中身を親に影響しない。
```bash
$ go test -v -count=1 ./...
=== RUN   TestSetenv
start TestSetenv
=== RUN   TestSetenv/sub
TEST_ENV is "sub test" in sub
TEST_ENV is "test" in parent
finish TestSetenv
--- PASS: TestSetenv (5.00s)
    --- PASS: TestSetenv/sub (0.00s)
=== RUN   TestWatch
TEST_ENV is empty
--- PASS: TestWatch (0.00s)
PASS
ok      github.com/budougumi0617/til/go/testing/env/setenv      6.641s
=== RUN   TestWatch
start TestWatch
TEST_ENV is empty
TEST_ENV is empty
TEST_ENV is empty
TEST_ENV is empty
TEST_ENV is empty
finish TestWatch
--- PASS: TestWatch (5.01s)
PASS
ok      github.com/budougumi0617/til/go/testing/env/watch       5.825s

```