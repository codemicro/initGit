# initGit

`initGit` automates a few tiny things that I have to do every time I make a new repository.

  * Adds a `LICENCE`
  * Adds a `.gitignore`
  * Creates a project layout using predefined templates
  * Runs `git init` and commits these things for you
  * Adds a remote (or doesn't) of your choosing

## Build

Building from source requires a version of Python 3 to be installed. If you're building on anything other than Windows, 
you may need to change the `go:generate` declaration in `internal/data/data.go` to point to the correct version of
Python installed on your system.

```
go generate github.com/codemicro/initGit/...
go build github.com/codemicro/initGit/cmd/initGit
```

## Included templates

* [Golang with Go modules](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/go.json)
* [Python with Poetry](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/python.json)
