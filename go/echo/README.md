# Sample of echo

## Why install v3?
github.com/labstack/echo というモジュールの最新版がv3だから。  
https://github.com/labstack/echo/tree/v3.3.10 とかをみるとわかるのですが、v3まではgo.modがないです。（つまり `github.com/labstack/echo/v2` とか `github.com/labstack/echo/v3`  というmoduleがないということでもあります
echoはv4になったときはじめてgo.modが置かれ、go.modが置かれたとき`github.com/labstack/echo/v4`というモジュール名になった（`go.mod`へmoudle名が書かれるようになった）。  
なので、`v4`は`github.com/labstack/echo` moduleに存在しない。 `github.com/labstack/echo`というmoduleの最新版を取ってこようとすると、latestがv3になる。  
なので、`github.com/labstack/echo` module名でv4をinstallしようとしたとき失敗する。

```bash
$ go get github.com/labstack/echo@v4.0.0
go: github.com/labstack/echo@v4.0.0: invalid version: go.mod has post-v4 module path "github.com/labstack/echo/v4" at revision v4.0.0
```