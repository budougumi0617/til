.PHONY: list_import



# 再帰的にimportしているpkgを表示する
list_import: ## list import pkgs
	go list -f '{{join .Deps "\n"}}' ./... \
		| xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'
