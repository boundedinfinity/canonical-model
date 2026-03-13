m   := "updates"

list:
    just --list

git-push:
    git add . || true
    git commit -m "{{ m }}" || true
    git push origin master


test:
	cd example/go && just test

generate:
	cd example/go && just generate

commit:
	git add . || true
	git commit -m "{{ m }}" || true

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

push: commit
	just commit m={{ m }}

docs-dev:
	cd $(justfile_dir)/docs && just dev
