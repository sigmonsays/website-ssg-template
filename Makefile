default: build

dev: ## build with drafts
	go generate ./site
	go install ./website
	website -dev -gen
	rsync -ar ./overlay/ ./docs/

build: ## build without drafts
	go generate ./site
	go install ./website
	website -dev -gen
	rsync -ar ./overlay/ ./docs/

serve: ## serve HTTP server
	website -dev
