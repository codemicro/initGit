# initGit

`initGit` automates a few tiny things that I have to do every time I make a new repository.

  * Adds a `LICENCE`
  * Adds a `.gitignore`
  * Creates a `README.md`
  * Runs `git init` and commits these things for you
  * Adds a remote (or doesn't) of your choosing

### Build

```
go get -u github.com/codemicro/initGit
go build github.com/codemicro/initGit
```

You might want to stick the resultant binary on your `PATH`, too.

### TODO (eventually)

* Interface with the GitHub API to create new repositories