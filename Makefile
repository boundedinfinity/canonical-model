makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list purge build install generate test commit tag publish

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

test:
	cd $(makefile_dir)/lang/go && make test

generate:
	cd $(makefile_dir)/lang/go && make generate

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
