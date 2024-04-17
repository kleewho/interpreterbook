set fallback
default_name := "interpreterbook"

build name=default_name:
    go build -o build/{{name}}

run *ARGS: (build default_name)
    build/./{{default_name}} {{ARGS}}

debug *ARGS:
    dlv debug main.go -- {{ARGS}}

record: build
    vhs repl.tape