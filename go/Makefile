.PHONY: list_import why

# 再帰的にimportしているpkgを表示する
list_import: ## list import pkgs
	go list -f '{{join .Deps "\n"}}' ./... \
		| xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'

# pk1, pk2への依存がどの経路で発生しているのか確認する
why:
	go mod why -m pkg1 pkg2
