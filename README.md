# Hstsman

[![GoDoc](https://godoc.org/github.com/air-gases/hstsman?status.svg)](https://godoc.org/github.com/air-gases/hstsman)

A useful gas that used to manage the Strict-Transport-Security header for the
web applications built using [Air](https://github.com/aofei/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/hstsman
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.11.

## Usage

Create a file named `main.go`

```go
package main

import (
	"github.com/air-gases/hstsman"
	"github.com/aofei/air"
)

func main() {
	a := air.Default
	a.DebugMode = true
	a.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("Go and see the response header.")
	}, hstsman.Gas(hstsman.GasConfig{
		MaxAge:            31536000,
		IncludeSubDomains: true,
	}))
	a.Serve()
}
```

and run it

```bash
$ go run main.go
```

then visit `http://localhost:8080`.

## Community

If you want to discuss Hstsman, or ask questions about it, simply post questions
or ideas [here](https://github.com/air-gases/hstsman/issues).

## Contributing

If you want to help build Hstsman, simply follow
[this](https://github.com/air-gases/hstsman/wiki/Contributing) to send pull
requests [here](https://github.com/air-gases/hstsman/pulls).

## License

This project is licensed under the Unlicense.

License can be found [here](LICENSE).
