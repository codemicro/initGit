# initGit

![ISC License](https://img.shields.io/github/license/codemicro/initGit) ![Lines of code](https://img.shields.io/tokei/lines/github/codemicro/initGit)

`initGit` automates a few tiny things that I have to do every time I make a new repository.

  * Adds a `LICENCE`
  * Adds a `.gitignore`
  * Creates a project layout using predefined templates
  * Runs `git init` and commits these things for you
  * Adds a remote (or doesn't) of your choosing

Alternatively, it'll do only a select of those options, if you so desire.

initGit includes a few templates. A template defines files and directories to create and commands to run. The included ones can be used as examples, and are:

* [Golang with Go modules](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/go.json)
* [Python with Poetry](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/python.json)
* [README only](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/readme.json)

If you add your own template, feel free to submit a PR to have it included in the project.

### Prerequisites

* Go >= 1.15 (1.15 being the latest version tested on)
* Python 3 (for building)

### Installation

If you're building on anything other than Windows, you may need to change the `go:generate` declaration in `internal/data/data.go`
to point to the correct version of Python installed on your system.

1. Clone the repo
   ```sh
   git clone https://github.com/codemicro/initGit.git
   ```
2. Build
   ```sh
   cd initGit
   go generate github.com/codemicro/initGit/...
   go build github.com/codemicro/initGit/cmd/initGit
   ```
3. Run!
   ```sh
   sudo chmod +x initGit
   ./initGit
   ```

## License

Distributed under the MIT License. See `LICENCE` for more information.
