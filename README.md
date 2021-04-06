# Gravatar Go Module [![CI Status](https://circleci.com/gh/gopherlibs/gravatar.svg?style=shield)](https://app.circleci.com/pipelines/github/gopherlibs/gravatar) [![Go Report Card](https://goreportcard.com/badge/github.com/gopherlibs/gravatar)](https://goreportcard.com/report/github.com/gopherlibs/gravatar) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/gravatar/master/LICENSE)

`gravatar` is a Go (Golang) module that allows you to generate Gravatar URLs easily in your own Go project.


## Requirements

This Go module is tested with the ~~two~~ most recent minor ~~releases~~ release of Go.
Currently this is Go v1.16.


## Installation

`gravatar` is a Go module and so the best way to use it is to import it into your own code and then run `go mod tidy` to get it downloaded.

```go
import(
	"github.com/gopherlibs/gravatar/gravatar"
)
```


## Usage


```go
package main

import (
	"fmt"

	"github.com/gopherlibs/gravatar/gravatar"
)

func main() {

	img, err := gravatar.NewImage("Ricardo@Feliciano.Tech")
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

This library is written and tested with Go v1.16+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## License & Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
