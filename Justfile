m				:= "updates"

list:
    just --list

git-push:
    git add . || true
    git commit -m "{{ m }}" || true
    git push origin master
