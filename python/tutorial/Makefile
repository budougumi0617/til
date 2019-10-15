.PHONY: help install pcs pds flk

.DEFAULT_GOAL=help

install: ## Install Dependencies by Pipenv
	@pipenv install

pcs:  ## Execute pycodestyle
	@pipenv run pycodestyle *.py

pds:  ## Execute pydocstyle
	@pipenv run pydocstyle *.py

flk: ## Execute pyflakes
	@pipenv run pyflakes *.py

lint: pcs pds flk ## Execute all linter

test: ## Execute pytest
	@PYTHONPATH=./ pipenv run pytest


help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
