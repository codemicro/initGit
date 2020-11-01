# initGit

`initGit` automates a few tiny things that I have to do every time I make a new repository.

1: Adds a `LICENCE`
2: Adds a `.gitignore`
3: Creates a `README.md`
4: Runs `git init` and commits these things for you
5: Adds a remote (or doesn't) of your choosing

### Build

```
go get -u github.com/codemicro/initGit
go build github.com/codemicro/initGit
```

You might want to stick the resultant binary on your `PATH`, too.

### TODO (eventually)

* Interface with the GitHub API to create new repositories