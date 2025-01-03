# hello_world

## Study

- `cargo run`
- `clippy`
  - Cursor 用のは setting.json に設定してある

### neovim 用の Rust の設定

Lazyvim の Extra は rust-analyzer 以外もインストールするので明示的に rust-analyzer をインストールする。

```lua
return {
  -- tools
  {
    "williamboman/mason.nvim",
    opts = function(_, opts)
      vim.list_extend(opts.ensure_installed, {
        -- ...
        "rust_analyzer",
      })
    end,
  },
  {
    "neovim/nvim-lspconfig",
    ---@class PluginLspOpts
    opts = {
      ---@type lspconfig.options
      servers = {
        rust_analyzer = {
          settings = {
            ["rust-analyzer"] = {
              -- checkOnSave で clippy を自動実行
              checkOnSave = {
                command = "clippy", -- ここを 'clippy' にすると保存時に自動実行
              },
            },
          },
        },
      },
    },
  },
}
```
