# compressible
[![Build Status](https://travis-ci.org/go-http-utils/compressible.svg?branch=master)](https://travis-ci.org/go-http-utils/compressible)
[![Coverage Status](https://coveralls.io/repos/github/go-http-utils/compressible/badge.svg?branch=master)](https://coveralls.io/github/go-http-utils/compressible?branch=master)

MIME compressible checking for Go

## Installation

```
go get -u github.com/go-http-utils/compressible
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/go-http-utils/compressible

## Usage

```go
import (
  _ "github.com/GitbookIO/mimedb/autoload"
  "github.com/go-http-utils/compressible"
)
```

```go
fmt.Println(compressible.Test("text/html"))
// -> true

fmt.Println(compressible.Test("image/jpeg"))
// -> false
```
