makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list purge build install generate test commit tag publish

purge:
	find . -name '*.enum.go' -type f -delete
	find . -name '*.gen.go' -type f -delete

generate:
	@make purge
	go generate ./...

test:
	go test ./...

build:
	@make purge
	@make generate
	go build .

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

commit:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

publish:
	make commit m=$(m)
	make tag tag=$(m)
