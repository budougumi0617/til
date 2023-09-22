# ターミナルを開く

この辺を設定すると`CTRL+;`で行ったり来たりできる。

```json
[
  {
    "key": "ctrl+;",
    "command": "workbench.action.terminal.focus",
    "when": "editorTextFocus"
  },
  {
    "key": "ctrl+;",
    "command": "workbench.action.focusFirstEditorGroup",
    "when": "terminalFocus"
  }
]
```

# 新しいファイルを作る

`explorer.newFile`にショートカットを設定すると選択中のフォルダにファイルを作れるようになる。

```json
[
  {
    "key": "cmd+n",
    "command": "explorer.newFile"
  }
]
```
