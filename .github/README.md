<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/codemicro/initGit">
    <img src="https://raw.githubusercontent.com/codemicro/initGit/master/.github/img.png" alt="Logo" height="200px">
  </a>

  <h3 align="center">initGit</h3>

  <p align="center">
    Streamline your repository creation process
    <br />
    <a href="https://github.com/codemicro/initGit/issues">Report bug</a>
    ·
    <a href="https://github.com/codemicro/initGit/issues">Request feature</a>
  </p>
</p>

<h2 style="display: inline-block">Table of Contents</h2>
<ol>
  <li>
    <a href="#about-the-project">About The Project</a>
    <ul>
      <li><a href="#built-with">Built With</a></li>
    </ul>
  </li>
  <li>
    <a href="#getting-started">Getting Started</a>
    <ul>
      <li><a href="#prerequisites">Prerequisites</a></li>
      <li><a href="#installation">Installation</a></li>
    </ul>
  </li>
  <li><a href="#usage">Usage</a></li>
  <li><a href="#contributing">Contributing</a></li>
  <li><a href="#license">License</a></li>
</ol>

<!-- ABOUT THE PROJECT -->
## About The Project

`initGit` automates a few tiny things that I have to do every time I make a new repository.

  * Adds a `LICENCE`
  * Adds a `.gitignore`
  * Creates a project layout using predefined templates
  * Runs `git init` and commits these things for you
  * Adds a remote (or doesn't) of your choosing

initGit includes a few templates. A template defines files and directories to create and commands to run. The included ones can be used as examples, and are:

* [Golang with Go modules](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/go.json)
* [Python with Poetry](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/python.json)
* [README only](https://github.com/codemicro/initGit/blob/master/internal/data/dataFiles/templates/readme.json)

If you add your own template, feel free to submit a PR to have it included in the project.

### Built With

* [go-bindata](https://github.com/go-bindata/go-bindata/)

<!-- GETTING STARTED -->
## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

* Go >= 1.15 (1.15 being the latest version tested on)
* Python 3 (for building)

### Installation

Building from source requires a version of Python 3 to be installed. If you're building on anything other than Windows, 
you may need to change the `go:generate` declaration in `internal/data/data.go` to point to the correct version of
Python installed on your system.

1. Clone the repo
   ```sh
   git clone https://github.com/codemicro/initGit.git
   ```
2. Build the project
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

<!-- USAGE EXAMPLES -->
## Usage

Run the `initGit` executable (no arguments required) and follow the instructions.

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the project
2. Create a Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to your branch and open a PR (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->
## License

[![ISC License][license-shield]][license-url]

Distributed under the ISC License. See `LICENSE` for more information.

[license-shield]: https://img.shields.io/github/license/codemicro/initGit.svg?style=for-the-badge
[license-url]: https://github.com/codemicro/initGit/blob/master/LICENSE.txt