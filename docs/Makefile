makefile_dir	:= $(abspath $(shell pwd))
m				?= "updates"

.PHONY: list purge build install generate test commit tag publish docs

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

dev:
	npm install && npm run dev
