# gPic Go Module [![CI Status](https://circleci.com/gh/gopherlibs/gpic.svg?style=shield)](https://app.circleci.com/pipelines/github/gopherlibs/gpic) [![Go Report Card](https://goreportcard.com/badge/github.com/gopherlibs/gpic)](https://goreportcard.com/report/github.com/gopherlibs/gpic) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/gpic/master/LICENSE)

`gPic` is a Go (Golang) module that allows you to generate Avatar URLs easily in your own Go project. This currently support Gravatar and GitHub avatars. GitLab support will be added soon.

## Requirements

This Go module is tested with the ~~two~~ most recent minor ~~releases~~ release of Go.
Currently this is Go v1.20.

## Installation

`gPic` is a Go module and so the best way to use it is to import it into your own code and then run `go mod tidy` to get it downloaded.

```go
import(
	"github.com/gopherlibs/gpic/gpic"
)
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/gopherlibs/gpic/gpic"
)

func main() {

	img, err := gpic.NewImage("Ricardo@Feliciano.Tech")
	if err != nil {
		fmt.Println(err)
		return
	}

	imgURL, err := img.URL()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(imgURL.String())
}
```

## Development

This library is written and tested with Go v1.20+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.

## License & Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
