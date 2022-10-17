# gremlin-go-example

GoでGremlinを実行するサンプルです

## 使い方

### devcontainer

VS Codeのdevcontainerを起動するとgremlin-serverが実行されます。
そのまま`go run main.go`を実行するとGoのプログラムを試せます。
gremlin-consoleもインストールされており、`gremlin-console`コマンドで実行できます。

### nix-shell

`nix-shell shell.nix`を実行すると起動できます。
`gremlin-server`、`gremlin-console`、VSCodeでGoを使用するのに必要な環境が含まれています。
Goのプログラムを実行する場合は`gremlin-server`コマンドで先にgremlin-serverを立ち上げてください。
