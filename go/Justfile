m   := "updates"

list:
    just --list

purge:
	find . -name '*.enum.go' -type f -delete
	find . -name '*.gen.go' -type f -delete

generate:
	@just purge
	go generate ./...

test:
	go test ./...

build:
	@just purge
	@just generate
	go build .

commit:
	git add . || true
	git commit -m "{{ m }}" || true

push: commit
	git push origin master

tag tag:
	git tag -a {{ tag }} -m "{{ tag }}"
	git push origin {{ tag }}

publish:
	just commit m={{ m }}
	just tag tag={{ m }}
