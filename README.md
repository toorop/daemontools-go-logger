## daemontools-go-logger

Simple logging package for go services supervised by DJB deamontools

## Installing

### Using *go get*

    $ go get github.com/Toorop/daemontools-go-logger

You can use `go get -u -a` to update all installed packages.

## Example

```go
package main

import "github.com/Toorop/daemontools-go-logger"

var log = new(logger.Logger)

func main() {
	log.SetLevel("debug")
	log.SetTimeStamp(true)
	log.Debug("toto")
}

```

## Documentation

For docs, see http://godoc.org/github.com/Toorop/daemontools-go-logger or run:

    $ go doc github.com/Toorop/daemontools-go-logger
